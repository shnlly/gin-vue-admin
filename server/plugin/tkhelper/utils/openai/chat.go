package openai

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/openai/openai-go/v2"
	"go.uber.org/zap"
)

// Chat 发送聊天请求（非流式）
func (c *Client) Chat(ctx context.Context, request ChatRequest) (*ChatResponse, error) {
	c.logRequestInfo(request.Messages, false)

	openaiMessages := c.convertMessages(request.Messages)

	params := openai.ChatCompletionNewParams{
		Model:    c.aiModel.ModelName,
		Messages: openaiMessages,
	}

	response, err := c.client.Chat.Completions.New(ctx, params)
	if err != nil {
		c.logError(err, "普通请求")
		return nil, fmt.Errorf("调用AI接口失败: %v", err)
	}

	if len(response.Choices) == 0 {
		return nil, fmt.Errorf("AI未返回任何回复")
	}

	assistantContent := response.Choices[0].Message.Content
	tokenCount := int(response.Usage.TotalTokens)

	return &ChatResponse{
		Content:    assistantContent,
		TokenCount: tokenCount,
		Role:       "assistant",
	}, nil
}

// ChatStream 发送聊天请求（流式）
func (c *Client) ChatStream(ctx context.Context, request ChatRequest, handler StreamHandler) error {
	c.logRequestInfo(request.Messages, true)

	openaiMessages := c.convertMessages(request.Messages)

	params := openai.ChatCompletionNewParams{
		Model:    c.aiModel.ModelName,
		Messages: openaiMessages,
	}

	stream := c.client.Chat.Completions.NewStreaming(ctx, params)

	var fullContent strings.Builder
	var tokenCount int

	for stream.Next() {
		chunk := stream.Current()
		if len(chunk.Choices) > 0 && chunk.Choices[0].Delta.Content != "" {
			content := chunk.Choices[0].Delta.Content
			fullContent.WriteString(content)
			tokenCount++

			// 调用处理器处理消息片段
			if err := handler.OnMessage(content); err != nil {
				return fmt.Errorf("消息处理失败: %v", err)
			}
		}
	}

	if err := stream.Err(); err != nil {
		c.logError(err, "流式请求")
		return handler.OnError(fmt.Errorf("流式响应失败: %v", err))
	}

	// 调用完成处理器
	response := ChatResponse{
		Content:    fullContent.String(),
		TokenCount: tokenCount,
		Role:       "assistant",
	}

	return handler.OnComplete(response)
}

// SSEStreamHandler SSE流式处理器
type SSEStreamHandler struct {
	Writer io.Writer
}

// OnMessage 处理消息片段
func (h *SSEStreamHandler) OnMessage(content string) error {
	data := map[string]interface{}{
		"type":    "message",
		"content": content,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(h.Writer, "data: %s\n\n", jsonData)
	if err != nil {
		return err
	}

	// 尝试刷新缓冲区
	if flusher, ok := h.Writer.(interface{ Flush() }); ok {
		flusher.Flush()
	}

	return nil
}

// OnComplete 处理完成事件
func (h *SSEStreamHandler) OnComplete(response ChatResponse) error {
	// 发送结束信号
	endData := map[string]interface{}{
		"type":        "end",
		"content":     response.Content,
		"token_count": response.TokenCount,
	}
	jsonData, err := json.Marshal(endData)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(h.Writer, "data: %s\n\n", jsonData)
	if err != nil {
		return err
	}

	if flusher, ok := h.Writer.(interface{ Flush() }); ok {
		flusher.Flush()
	}

	global.GVA_LOG.Info("流式响应完成",
		zap.Int("token_count", response.TokenCount),
		zap.Int("content_length", len(response.Content)),
	)

	return nil
}

// OnError 处理错误事件
func (h *SSEStreamHandler) OnError(err error) error {
	global.GVA_LOG.Error("流式响应处理错误", zap.Error(err))

	errorData := map[string]interface{}{
		"type":  "error",
		"error": err.Error(),
	}
	jsonData, _ := json.Marshal(errorData)
	fmt.Fprintf(h.Writer, "data: %s\n\n", jsonData)

	if flusher, ok := h.Writer.(interface{ Flush() }); ok {
		flusher.Flush()
	}

	return err
}
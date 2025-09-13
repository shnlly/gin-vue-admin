package example

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
	exampleReq "github.com/flipped-aurora/gin-vue-admin/server/model/example/request"
	"github.com/openai/openai-go/v2"
	"github.com/openai/openai-go/v2/option"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AiChatMessageService struct{}

var AiChatMessageServiceApp = new(AiChatMessageService)

// min 辅助函数
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// CreateAiChatMessage 创建AI对话消息记录
func (aiChatMessageService *AiChatMessageService) CreateAiChatMessage(aiChatMessage *example.AiChatMessage) (err error) {
	err = global.GVA_DB.Create(aiChatMessage).Error
	return err
}

// DeleteAiChatMessage 删除AI对话消息记录
func (aiChatMessageService *AiChatMessageService) DeleteAiChatMessage(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&example.AiChatMessage{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&example.AiChatMessage{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteAiChatMessageByIds 批量删除AI对话消息记录
func (aiChatMessageService *AiChatMessageService) DeleteAiChatMessageByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&example.AiChatMessage{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&example.AiChatMessage{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateAiChatMessage 更新AI对话消息记录
func (aiChatMessageService *AiChatMessageService) UpdateAiChatMessage(aiChatMessage example.AiChatMessage) (err error) {
	err = global.GVA_DB.Model(&example.AiChatMessage{}).Where("id = ?", aiChatMessage.ID).Updates(&aiChatMessage).Error
	return err
}

// GetAiChatMessage 根据ID获取AI对话消息记录
func (aiChatMessageService *AiChatMessageService) GetAiChatMessage(ID string) (aiChatMessage example.AiChatMessage, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&aiChatMessage).Error
	return
}

// GetAiChatMessageInfoList 分页获取AI对话消息记录
func (aiChatMessageService *AiChatMessageService) GetAiChatMessageInfoList(info exampleReq.AiChatMessageSearch) (list []example.AiChatMessage, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&example.AiChatMessage{})
	var aiChatMessages []example.AiChatMessage

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", *info.StartCreatedAt, *info.EndCreatedAt)
	}
	if info.SessionId != 0 {
		db = db.Where("session_id = ?", info.SessionId)
	}
	if info.Role != "" {
		db = db.Where("role = ?", info.Role)
	}
	if info.Content != "" {
		db = db.Where("content LIKE ?", "%"+info.Content+"%")
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	db = db.Order("message_index ASC")
	err = db.Find(&aiChatMessages).Error
	return aiChatMessages, total, err
}

// GetSessionMessages 获取会话的所有消息
func (aiChatMessageService *AiChatMessageService) GetSessionMessages(sessionID uint) (list []example.AiChatMessage, err error) {
	err = global.GVA_DB.Where("session_id = ?", sessionID).
		Order("message_index ASC").
		Find(&list).Error
	return
}

// SendChatMessage 发送聊天消息并获取AI回复
func (aiChatMessageService *AiChatMessageService) SendChatMessage(req exampleReq.ChatRequest, userID uint, responseWriter io.Writer) error {
	// 获取AI模型信息
	var aiModel example.AiModel
	if err := global.GVA_DB.Where("id = ? AND enabled = ?", req.ModelId, true).First(&aiModel).Error; err != nil {
		return fmt.Errorf("AI模型不存在或未启用")
	}

	// 获取会话信息
	var session example.AiChatSession
	if err := global.GVA_DB.Where("id = ? AND user_id = ?", req.SessionId, userID).First(&session).Error; err != nil {
		return fmt.Errorf("会话不存在或无权限访问")
	}

	// 获取会话历史消息
	var historyMessages []example.AiChatMessage
	if err := global.GVA_DB.Where("session_id = ?", req.SessionId).
		Order("message_index ASC").
		Find(&historyMessages).Error; err != nil {
		return fmt.Errorf("获取历史消息失败: %v", err)
	}

	// 保存用户消息
	nextIndex := len(historyMessages)
	userMessage := example.AiChatMessage{
		SessionId:    req.SessionId,
		Role:         "user",
		Content:      req.Message,
		MessageIndex: nextIndex,
	}
	if err := global.GVA_DB.Create(&userMessage).Error; err != nil {
		return fmt.Errorf("保存用户消息失败: %v", err)
	}

	// 构建OpenAI消息格式
	var openaiMessages []openai.ChatCompletionMessageParamUnion
	for _, msg := range historyMessages {
		switch msg.Role {
		case "user":
			openaiMessages = append(openaiMessages, openai.UserMessage(msg.Content))
		case "assistant":
			openaiMessages = append(openaiMessages, openai.AssistantMessage(msg.Content))
		case "system":
			openaiMessages = append(openaiMessages, openai.SystemMessage(msg.Content))
		}
	}
	// 添加当前用户消息
	openaiMessages = append(openaiMessages, openai.UserMessage(req.Message))

	// 打印调试信息
	global.GVA_LOG.Info("AI模型调用信息",
		zap.String("model_name", aiModel.ModelName),
		zap.String("base_url", aiModel.BaseUrl),
		zap.String("api_key_prefix", aiModel.ApiKey[:min(len(aiModel.ApiKey), 10)]+"..."),
		zap.Bool("stream_mode", req.Stream),
		zap.Uint("session_id", req.SessionId),
	)

	// 创建OpenAI客户端配置选项
	var opts []option.RequestOption
	opts = append(opts, option.WithAPIKey(aiModel.ApiKey))
	if aiModel.BaseUrl != "" {
		baseURL := strings.TrimSuffix(aiModel.BaseUrl, "/")
		opts = append(opts, option.WithBaseURL(baseURL))
		global.GVA_LOG.Info("使用自定义Base URL", zap.String("base_url", baseURL))
	}
	client := openai.NewClient(opts...)

	// 调用OpenAI API
	if req.Stream {
		return aiChatMessageService.handleStreamResponse(client, aiModel, openaiMessages, req.SessionId, nextIndex+1, responseWriter)
	} else {
		return aiChatMessageService.handleNormalResponse(client, aiModel, openaiMessages, req.SessionId, nextIndex+1)
	}
}

// handleStreamResponse 处理流式响应
func (aiChatMessageService *AiChatMessageService) handleStreamResponse(client openai.Client, aiModel example.AiModel, messages []openai.ChatCompletionMessageParamUnion, sessionID uint, messageIndex int, writer io.Writer) error {
	ctx := context.Background()

	global.GVA_LOG.Info("开始流式请求",
		zap.String("model_name", aiModel.ModelName),
		zap.Int("message_count", len(messages)),
		zap.Uint("session_id", sessionID),
	)

	stream := client.Chat.Completions.NewStreaming(ctx, openai.ChatCompletionNewParams{
		Model:    aiModel.ModelName,
		Messages: messages,
	})

	var assistantMessage strings.Builder
	var tokenCount int

	// 发送流式数据
	for stream.Next() {
		chunk := stream.Current()
		if len(chunk.Choices) > 0 && chunk.Choices[0].Delta.Content != "" {
			content := chunk.Choices[0].Delta.Content
			assistantMessage.WriteString(content)
			tokenCount++

			// 发送SSE数据
			data := map[string]interface{}{
				"type":    "message",
				"content": content,
			}
			jsonData, _ := json.Marshal(data)
			fmt.Fprintf(writer, "data: %s\n\n", jsonData)
			if flusher, ok := writer.(interface{ Flush() }); ok {
				flusher.Flush()
			}
		}
	}

	if err := stream.Err(); err != nil {
		global.GVA_LOG.Error("流式响应错误详情",
			zap.Error(err),
			zap.String("model_name", aiModel.ModelName),
			zap.String("base_url", aiModel.BaseUrl),
			zap.String("api_key_length", fmt.Sprintf("%d", len(aiModel.ApiKey))),
		)
		return fmt.Errorf("流式响应失败: %v", err)
	}

	// 保存AI回复消息
	aiMessage := example.AiChatMessage{
		SessionId:    sessionID,
		Role:         "assistant",
		Content:      assistantMessage.String(),
		TokenCount:   tokenCount,
		MessageIndex: messageIndex,
	}
	if err := global.GVA_DB.Create(&aiMessage).Error; err != nil {
		global.GVA_LOG.Error("保存AI回复消息失败: " + err.Error())
	}

	// 更新会话信息
	if err := AiChatSessionServiceApp.UpdateSessionMessage(sessionID, aiMessage.Content); err != nil {
		global.GVA_LOG.Error("更新会话信息失败: " + err.Error())
	}

	// 发送结束信号
	endData := map[string]interface{}{
		"type": "end",
	}
	jsonData, _ := json.Marshal(endData)
	fmt.Fprintf(writer, "data: %s\n\n", jsonData)
	if flusher, ok := writer.(interface{ Flush() }); ok {
		flusher.Flush()
	}

	return nil
}

// handleNormalResponse 处理普通响应
func (aiChatMessageService *AiChatMessageService) handleNormalResponse(client openai.Client, aiModel example.AiModel, messages []openai.ChatCompletionMessageParamUnion, sessionID uint, messageIndex int) error {
	ctx := context.Background()

	global.GVA_LOG.Info("开始普通请求",
		zap.String("model_name", aiModel.ModelName),
		zap.Int("message_count", len(messages)),
		zap.Uint("session_id", sessionID),
	)

	response, err := client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Model:    aiModel.ModelName,
		Messages: messages,
	})
	if err != nil {
		global.GVA_LOG.Error("普通请求错误详情",
			zap.Error(err),
			zap.String("model_name", aiModel.ModelName),
			zap.String("base_url", aiModel.BaseUrl),
			zap.String("api_key_length", fmt.Sprintf("%d", len(aiModel.ApiKey))),
		)
		return fmt.Errorf("调用AI接口失败: %v", err)
	}

	if len(response.Choices) == 0 {
		return fmt.Errorf("AI未返回任何回复")
	}

	assistantContent := response.Choices[0].Message.Content
	tokenCount := int(response.Usage.TotalTokens)

	// 保存AI回复消息
	aiMessage := example.AiChatMessage{
		SessionId:    sessionID,
		Role:         "assistant",
		Content:      assistantContent,
		TokenCount:   tokenCount,
		MessageIndex: messageIndex,
	}
	if err := global.GVA_DB.Create(&aiMessage).Error; err != nil {
		return fmt.Errorf("保存AI回复消息失败: %v", err)
	}

	// 更新会话信息
	if err := AiChatSessionServiceApp.UpdateSessionMessage(sessionID, assistantContent); err != nil {
		global.GVA_LOG.Error("更新会话信息失败: " + err.Error())
	}

	return nil
}

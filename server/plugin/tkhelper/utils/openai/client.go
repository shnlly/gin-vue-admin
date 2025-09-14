package openai

import (
	"fmt"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/tkhelper/model"
	"github.com/openai/openai-go/v2"
	"github.com/openai/openai-go/v2/option"
	"go.uber.org/zap"
)

// NewClient 创建新的OpenAI客户端
func NewClient(aiModel *model.AiModel) *Client {
	var opts []option.RequestOption
	opts = append(opts, option.WithAPIKey(aiModel.ApiKey))

	if aiModel.BaseUrl != "" {
		baseURL := strings.TrimSuffix(aiModel.BaseUrl, "/")
		opts = append(opts, option.WithBaseURL(baseURL))
		global.GVA_LOG.Info("使用自定义Base URL", zap.String("base_url", baseURL))
	}

	client := openai.NewClient(opts...)

	return &Client{
		client:  client,
		aiModel: aiModel,
	}
}

// NewClientWithConfig 通过配置创建客户端
func NewClientWithConfig(config ClientConfig) *Client {
	var opts []option.RequestOption
	opts = append(opts, option.WithAPIKey(config.APIKey))

	if config.BaseURL != "" {
		baseURL := strings.TrimSuffix(config.BaseURL, "/")
		opts = append(opts, option.WithBaseURL(baseURL))
	}

	client := openai.NewClient(opts...)

	// 创建临时模型对象
	aiModel := &model.AiModel{
		Name:      "临时模型",
		BaseUrl:   config.BaseURL,
		ApiKey:    config.APIKey,
		ModelName: config.Model,
	}

	return &Client{
		client:  client,
		aiModel: aiModel,
	}
}

// min 辅助函数
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// convertMessages 转换消息格式
func (c *Client) convertMessages(messages []ChatMessage) []openai.ChatCompletionMessageParamUnion {
	var openaiMessages []openai.ChatCompletionMessageParamUnion
	for _, msg := range messages {
		switch msg.Role {
		case "user":
			openaiMessages = append(openaiMessages, openai.UserMessage(msg.Content))
		case "assistant":
			openaiMessages = append(openaiMessages, openai.AssistantMessage(msg.Content))
		case "system":
			openaiMessages = append(openaiMessages, openai.SystemMessage(msg.Content))
		}
	}
	return openaiMessages
}

// logRequestInfo 记录请求信息
func (c *Client) logRequestInfo(messages []ChatMessage, stream bool) {
	global.GVA_LOG.Info("AI模型调用信息",
		zap.String("model_name", c.aiModel.ModelName),
		zap.String("base_url", c.aiModel.BaseUrl),
		zap.String("api_key_prefix", c.aiModel.ApiKey[:min(len(c.aiModel.ApiKey), 10)]+"..."),
		zap.Bool("stream_mode", stream),
		zap.Int("message_count", len(messages)),
	)
}

// logError 记录错误信息
func (c *Client) logError(err error, requestType string) {
	global.GVA_LOG.Error(fmt.Sprintf("%s错误详情", requestType),
		zap.Error(err),
		zap.String("model_name", c.aiModel.ModelName),
		zap.String("base_url", c.aiModel.BaseUrl),
		zap.String("api_key_length", fmt.Sprintf("%d", len(c.aiModel.ApiKey))),
	)
}
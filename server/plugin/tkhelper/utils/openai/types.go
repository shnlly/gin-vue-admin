package openai

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/tkhelper/model"
	"github.com/openai/openai-go/v2"
)

// ChatMessage 统一的聊天消息结构
type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ChatRequest 聊天请求
type ChatRequest struct {
	Model        string        `json:"model"`
	Messages     []ChatMessage `json:"messages"`
	Stream       bool          `json:"stream"`
	Temperature  float64       `json:"temperature,omitempty"`
	MaxTokens    int           `json:"max_tokens,omitempty"`
	TopP         float64       `json:"top_p,omitempty"`
	PresencePenalty float64    `json:"presence_penalty,omitempty"`
	FrequencyPenalty float64   `json:"frequency_penalty,omitempty"`
}

// ChatResponse 聊天响应
type ChatResponse struct {
	Content    string `json:"content"`
	TokenCount int    `json:"token_count"`
	Role       string `json:"role"`
}

// StreamHandler 流式响应处理器接口
type StreamHandler interface {
	OnMessage(content string) error
	OnComplete(response ChatResponse) error
	OnError(err error) error
}

// Client OpenAI客户端包装器
type Client struct {
	client   openai.Client
	aiModel  *model.AiModel
}

// ClientConfig 客户端配置
type ClientConfig struct {
	APIKey  string
	BaseURL string
	Model   string
}
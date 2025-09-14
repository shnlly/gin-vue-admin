package openai

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/tkhelper/model"
)

// ConvertFromChatMessages 从数据库聊天消息转换为OpenAI消息格式
func ConvertFromChatMessages(messages []model.AiChatMessage) []ChatMessage {
	var chatMessages []ChatMessage
	for _, msg := range messages {
		chatMessages = append(chatMessages, ChatMessage{
			Role:    msg.Role,
			Content: msg.Content,
		})
	}
	return chatMessages
}

// ConvertToChatMessage 转换为数据库聊天消息格式
func ConvertToChatMessage(sessionID uint, role, content string, messageIndex int, tokenCount int) *model.AiChatMessage {
	return &model.AiChatMessage{
		SessionId:    sessionID,
		Role:         role,
		Content:      content,
		TokenCount:   tokenCount,
		MessageIndex: messageIndex,
	}
}

// BuildChatRequest 构建聊天请求
func BuildChatRequest(modelName string, messages []model.AiChatMessage, newMessage string, stream bool) ChatRequest {
	// 转换历史消息
	chatMessages := ConvertFromChatMessages(messages)

	// 添加新消息
	if newMessage != "" {
		chatMessages = append(chatMessages, ChatMessage{
			Role:    "user",
			Content: newMessage,
		})
	}

	return ChatRequest{
		Model:    modelName,
		Messages: chatMessages,
		Stream:   stream,
	}
}

// TruncateMessage 截取消息预览（用于会话最后消息）
func TruncateMessage(message string, maxLength int) string {
	if maxLength <= 0 {
		maxLength = 100
	}

	if len(message) <= maxLength {
		return message
	}

	// 按Unicode字符截取，避免中文乱码
	runes := []rune(message)
	if len(runes) > maxLength {
		return string(runes[:maxLength]) + "..."
	}

	return message
}
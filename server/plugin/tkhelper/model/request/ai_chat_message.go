package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/tkhelper/model"
)

type AiChatMessageSearch struct {
	model.AiChatMessage
	request.PageInfo
	StartCreatedAt *string `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *string `json:"endCreatedAt" form:"endCreatedAt"`
}

// ChatRequest 发送消息请求
type ChatRequest struct {
	SessionId uint   `json:"sessionId" binding:"required"` // 会话ID
	ModelId   uint   `json:"modelId" binding:"required"`   // 模型ID
	Message   string `json:"message" binding:"required"`   // 用户消息内容
	Stream    bool   `json:"stream"`                       // 是否流式响应
}
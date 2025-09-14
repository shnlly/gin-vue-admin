package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// AiChatMessage AI对话消息管理
type AiChatMessage struct {
	global.GVA_MODEL
	SessionId    uint   `json:"sessionId" form:"sessionId" gorm:"column:session_id;comment:所属对话会话ID;not null;index"`              // 会话ID
	Role         string `json:"role" form:"role" gorm:"column:role;comment:消息角色：user-用户,assistant-助手,system-系统;size:20;not null"` // 消息角色
	Content      string `json:"content" form:"content" gorm:"column:content;comment:消息的具体内容;type:text;not null"`                  // 消息内容
	TokenCount   int    `json:"tokenCount" form:"tokenCount" gorm:"column:token_count;comment:消息消耗的Token数量;default:0"`            // Token数量
	MessageIndex int    `json:"messageIndex" form:"messageIndex" gorm:"column:message_index;comment:消息在会话中的序号;default:0;index"`   // 消息序号
}

// TableName 设置表名，添加tkhelper前缀避免冲突
func (AiChatMessage) TableName() string {
	return "tkhelper_ai_chat_messages"
}
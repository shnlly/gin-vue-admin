package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// AiChatSession AI对话会话管理
type AiChatSession struct {
	global.GVA_MODEL
	Title        string `json:"title" form:"title" gorm:"column:title;comment:对话会话的标题;size:200;not null"`                              // 会话标题
	ModelId      uint   `json:"modelId" form:"modelId" gorm:"column:model_id;comment:使用的AI模型ID;not null"`                              // AI模型ID
	ModelName    string `json:"modelName" form:"modelName" gorm:"column:model_name;comment:AI模型名称快照;size:100"`                         // 模型名称
	MessageCount int    `json:"messageCount" form:"messageCount" gorm:"column:message_count;comment:会话中的消息总数;default:0"`               // 消息数量
	LastMessage  string `json:"lastMessage" form:"lastMessage" gorm:"column:last_message;comment:会话的最后一条消息内容;size:500"`                // 最后消息
	Status       string `json:"status" form:"status" gorm:"column:status;comment:会话状态：active-活跃,archived-归档;size:20;default:'active'"` // 会话状态
	UserId       uint   `json:"userId" form:"userId" gorm:"column:user_id;comment:会话所属用户ID;index"`                                     // 用户ID
	CreatedBy    uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy    uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy    uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 设置表名，添加tkhelper前缀避免冲突
func (AiChatSession) TableName() string {
	return "tkhelper_ai_chat_sessions"
}
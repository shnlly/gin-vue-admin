package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// AiModel AI大模型管理
type AiModel struct {
	global.GVA_MODEL
	Name      string `json:"name" form:"name" gorm:"column:name;comment:AI模型显示名称;size:100;not null"`                     // 模型名称
	BaseUrl   string `json:"baseUrl" form:"baseUrl" gorm:"column:base_url;comment:大模型API的基础URL地址;size:500;not null"`     // API基础URL
	ApiKey    string `json:"apiKey" form:"apiKey" gorm:"column:api_key;comment:访问大模型API的密钥;size:255;not null"`           // API密钥
	ModelName string `json:"modelName" form:"modelName" gorm:"column:model_name;comment:大模型的具体模型标识名称;size:100;not null"` // 模型标识
	Enabled   *bool  `json:"enabled" form:"enabled" gorm:"column:enabled;comment:模型是否启用状态;default:true"`                 // 是否启用
	IsDefault *bool  `json:"isDefault" form:"isDefault" gorm:"column:is_default;comment:是否为系统默认使用的模型;default:false"`     // 是否默认
	Sort      int    `json:"sort" form:"sort" gorm:"column:sort;comment:显示排序，数字越小越靠前;default:999"`                       // 排序
	Remark    string `json:"remark" form:"remark" gorm:"column:remark;comment:模型的备注信息;size:500"`                         // 备注
	CreatedBy uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 设置表名，添加tkhelper前缀避免冲突
func (AiModel) TableName() string {
	return "tkhelper_ai_models"
}
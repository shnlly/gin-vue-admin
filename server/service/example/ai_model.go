package example

import (
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
	exampleReq "github.com/flipped-aurora/gin-vue-admin/server/model/example/request"
	"gorm.io/gorm"
)

type AiModelService struct{}

var AiModelServiceApp = new(AiModelService)

// CreateAiModel 创建AI模型记录
func (aiModelService *AiModelService) CreateAiModel(aiModel *example.AiModel) (err error) {
	err = global.GVA_DB.Create(aiModel).Error
	return err
}

// DeleteAiModel 删除AI模型记录
func (aiModelService *AiModelService) DeleteAiModel(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&example.AiModel{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&example.AiModel{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteAiModelByIds 批量删除AI模型记录
func (aiModelService *AiModelService) DeleteAiModelByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&example.AiModel{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&example.AiModel{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateAiModel 更新AI模型记录
func (aiModelService *AiModelService) UpdateAiModel(aiModel example.AiModel) (err error) {
	err = global.GVA_DB.Model(&example.AiModel{}).Where("id = ?", aiModel.ID).Updates(&aiModel).Error
	return err
}

// GetAiModel 根据ID获取AI模型记录
func (aiModelService *AiModelService) GetAiModel(ID string) (aiModel example.AiModel, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&aiModel).Error
	return
}

// GetAiModelInfoList 分页获取AI模型记录
func (aiModelService *AiModelService) GetAiModelInfoList(info exampleReq.AiModelSearch) (list []example.AiModel, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&example.AiModel{})
	var aiModels []example.AiModel

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", *info.StartCreatedAt, *info.EndCreatedAt)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.ModelName != "" {
		db = db.Where("model_name LIKE ?", "%"+info.ModelName+"%")
	}
	if info.Enabled != nil {
		db = db.Where("enabled = ?", *info.Enabled)
	}
	if info.IsDefault != nil {
		db = db.Where("is_default = ?", *info.IsDefault)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	db = db.Order("sort ASC, created_at DESC")
	err = db.Find(&aiModels).Error
	return aiModels, total, err
}

// GetAiModelPublic 不鉴权的AI模型接口
func (aiModelService *AiModelService) GetAiModelPublic() {
	// 此方法为空实现，供公共访问
}

// SetDefaultAiModel 设置默认AI模型
func (aiModelService *AiModelService) SetDefaultAiModel(ID uint) (err error) {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// 先将所有模型设为非默认
		if err := tx.Model(&example.AiModel{}).Where("is_default = ?", true).Update("is_default", false).Error; err != nil {
			return err
		}
		// 设置指定模型为默认
		if err := tx.Model(&example.AiModel{}).Where("id = ?", ID).Update("is_default", true).Error; err != nil {
			return err
		}
		return nil
	})
}

// GetDefaultAiModel 获取默认AI模型
func (aiModelService *AiModelService) GetDefaultAiModel() (aiModel example.AiModel, err error) {
	err = global.GVA_DB.Where("is_default = ? AND enabled = ?", true, true).First(&aiModel).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 如果没有默认模型，返回第一个启用的模型
		err = global.GVA_DB.Where("enabled = ?", true).Order("sort ASC, created_at ASC").First(&aiModel).Error
	}
	return
}

// GetEnabledAiModels 获取所有启用的AI模型
func (aiModelService *AiModelService) GetEnabledAiModels() (list []example.AiModel, err error) {
	err = global.GVA_DB.Where("enabled = ?", true).Order("sort ASC, created_at DESC").Find(&list).Error
	return
}

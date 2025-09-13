package example

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
	exampleReq "github.com/flipped-aurora/gin-vue-admin/server/model/example/request"
	"gorm.io/gorm"
)

type AiChatSessionService struct{}

var AiChatSessionServiceApp = new(AiChatSessionService)

// CreateAiChatSession 创建AI对话会话记录
func (aiChatSessionService *AiChatSessionService) CreateAiChatSession(aiChatSession *example.AiChatSession) (err error) {
	err = global.GVA_DB.Create(aiChatSession).Error
	return err
}

// DeleteAiChatSession 删除AI对话会话记录
func (aiChatSessionService *AiChatSessionService) DeleteAiChatSession(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&example.AiChatSession{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&example.AiChatSession{}, "id = ?", ID).Error; err != nil {
			return err
		}
		// 同时删除相关的消息记录
		if err = tx.Where("session_id = ?", ID).Delete(&example.AiChatMessage{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteAiChatSessionByIds 批量删除AI对话会话记录
func (aiChatSessionService *AiChatSessionService) DeleteAiChatSessionByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&example.AiChatSession{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&example.AiChatSession{}).Error; err != nil {
			return err
		}
		// 同时删除相关的消息记录
		if err := tx.Where("session_id in ?", IDs).Delete(&example.AiChatMessage{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateAiChatSession 更新AI对话会话记录
func (aiChatSessionService *AiChatSessionService) UpdateAiChatSession(aiChatSession example.AiChatSession) (err error) {
	err = global.GVA_DB.Model(&example.AiChatSession{}).Where("id = ?", aiChatSession.ID).Updates(&aiChatSession).Error
	return err
}

// GetAiChatSession 根据ID获取AI对话会话记录
func (aiChatSessionService *AiChatSessionService) GetAiChatSession(ID string) (aiChatSession example.AiChatSession, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&aiChatSession).Error
	return
}

// GetAiChatSessionInfoList 分页获取AI对话会话记录
func (aiChatSessionService *AiChatSessionService) GetAiChatSessionInfoList(info exampleReq.AiChatSessionSearch) (list []example.AiChatSession, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&example.AiChatSession{})
	var aiChatSessions []example.AiChatSession

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", *info.StartCreatedAt, *info.EndCreatedAt)
	}
	if info.Title != "" {
		db = db.Where("title LIKE ?", "%"+info.Title+"%")
	}
	if info.Status != "" {
		db = db.Where("status = ?", info.Status)
	}
	if info.ModelId != 0 {
		db = db.Where("model_id = ?", info.ModelId)
	}
	if info.UserId != 0 {
		db = db.Where("user_id = ?", info.UserId)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	db = db.Order("updated_at DESC")
	err = db.Find(&aiChatSessions).Error
	return aiChatSessions, total, err
}

// GetUserChatSessions 获取用户的对话会话列表
func (aiChatSessionService *AiChatSessionService) GetUserChatSessions(userID uint) (list []example.AiChatSession, err error) {
	err = global.GVA_DB.Where("user_id = ? AND status = ?", userID, "active").
		Order("updated_at DESC").
		Find(&list).Error
	return
}

// UpdateSessionMessage 更新会话的最后消息和消息数量
func (aiChatSessionService *AiChatSessionService) UpdateSessionMessage(sessionID uint, lastMessage string) (err error) {
	err = global.GVA_DB.Model(&example.AiChatSession{}).
		Where("id = ?", sessionID).
		Updates(map[string]interface{}{
			"last_message":  lastMessage,
			"message_count": gorm.Expr("message_count + 1"),
		}).Error
	return err
}

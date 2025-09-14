package initialize

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/tkhelper/model"
	"go.uber.org/zap"
)

func Gorm(ctx context.Context) {
	err := global.GVA_DB.AutoMigrate(
		&model.VideoContentAnalysis{},
		&model.AiModel{},
		&model.AiChatSession{},
		&model.AiChatMessage{},
	)
	if err != nil {
		global.GVA_LOG.Error("tkhelper插件数据库迁移失败", zap.Error(err))
	} else {
		global.GVA_LOG.Info("tkhelper插件数据库迁移成功")
	}
}
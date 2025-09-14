package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/tkhelper/router"
	"github.com/gin-gonic/gin"
)

func Router(engine *gin.Engine) {
	private := engine.Group(global.GVA_CONFIG.System.RouterPrefix).Group("")
	private.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())

	// 创建tkhelper插件路由组，所有API都以/tkhelper开头
	tkhelperGroup := private.Group("tkhelper")

	// 注册各功能模块路由
	router.RouterGroupApp.InitVideoContentAnalysisRouter(tkhelperGroup)
	router.RouterGroupApp.InitAiModelRouter(tkhelperGroup)
	router.RouterGroupApp.InitAiChatSessionRouter(tkhelperGroup)
	router.RouterGroupApp.InitAiChatMessageRouter(tkhelperGroup)
}
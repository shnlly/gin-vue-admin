package example

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AiChatSessionRouter struct{}

// InitAiChatSessionRouter 初始化 AI对话会话 路由信息
func (s *AiChatSessionRouter) InitAiChatSessionRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	aiChatSessionRouter := Router.Group("aiChatSession").Use(middleware.OperationRecord())
	aiChatSessionRouterWithoutRecord := Router.Group("aiChatSession")

	var aiChatSessionApi = v1.ApiGroupApp.ExampleApiGroup.AiChatSessionApi
	{
		aiChatSessionRouter.POST("createAiChatSession", aiChatSessionApi.CreateAiChatSession)             // 新建AI对话会话
		aiChatSessionRouter.DELETE("deleteAiChatSession", aiChatSessionApi.DeleteAiChatSession)           // 删除AI对话会话
		aiChatSessionRouter.DELETE("deleteAiChatSessionByIds", aiChatSessionApi.DeleteAiChatSessionByIds) // 批量删除AI对话会话
		aiChatSessionRouter.PUT("updateAiChatSession", aiChatSessionApi.UpdateAiChatSession)              // 更新AI对话会话
	}
	{
		aiChatSessionRouterWithoutRecord.GET("findAiChatSession", aiChatSessionApi.FindAiChatSession)       // 根据ID获取AI对话会话
		aiChatSessionRouterWithoutRecord.GET("getAiChatSessionList", aiChatSessionApi.GetAiChatSessionList) // 获取AI对话会话列表
		aiChatSessionRouterWithoutRecord.GET("getUserSessions", aiChatSessionApi.GetUserChatSessions)       // 获取用户会话列表
	}
}

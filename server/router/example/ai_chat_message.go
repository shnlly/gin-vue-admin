package example

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AiChatMessageRouter struct{}

// InitAiChatMessageRouter 初始化 AI对话消息 路由信息
func (s *AiChatMessageRouter) InitAiChatMessageRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	aiChatMessageRouter := Router.Group("aiChatMessage").Use(middleware.OperationRecord())
	aiChatMessageRouterWithoutRecord := Router.Group("aiChatMessage")

	var aiChatMessageApi = v1.ApiGroupApp.ExampleApiGroup.AiChatMessageApi
	{
		aiChatMessageRouter.POST("createAiChatMessage", aiChatMessageApi.CreateAiChatMessage)             // 新建AI对话消息
		aiChatMessageRouter.DELETE("deleteAiChatMessage", aiChatMessageApi.DeleteAiChatMessage)           // 删除AI对话消息
		aiChatMessageRouter.DELETE("deleteAiChatMessageByIds", aiChatMessageApi.DeleteAiChatMessageByIds) // 批量删除AI对话消息
		aiChatMessageRouter.PUT("updateAiChatMessage", aiChatMessageApi.UpdateAiChatMessage)              // 更新AI对话消息
		aiChatMessageRouter.POST("sendMessage", aiChatMessageApi.SendMessage)                             // 发送消息
	}
	{
		aiChatMessageRouterWithoutRecord.GET("findAiChatMessage", aiChatMessageApi.FindAiChatMessage)       // 根据ID获取AI对话消息
		aiChatMessageRouterWithoutRecord.GET("getAiChatMessageList", aiChatMessageApi.GetAiChatMessageList) // 获取AI对话消息列表
		aiChatMessageRouterWithoutRecord.GET("getSessionMessages", aiChatMessageApi.GetSessionMessages)     // 获取会话消息列表
	}
}

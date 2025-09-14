package router

import (
	"github.com/gin-gonic/gin"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/tkhelper/api"
)

type AiChatSessionRouter struct{}

func (r *AiChatSessionRouter) InitAiChatSessionRouter(Router *gin.RouterGroup) {
	chatSessionRouter := Router.Group("chat-session")
	{
		chatSessionRouter.POST("create", api.ApiGroupApp.AiChatSessionApi.CreateAiChatSession)    // 创建对话会话
		chatSessionRouter.DELETE("delete", api.ApiGroupApp.AiChatSessionApi.DeleteAiChatSession) // 删除对话会话
		chatSessionRouter.DELETE("deleteByIds", api.ApiGroupApp.AiChatSessionApi.DeleteAiChatSessionByIds) // 批量删除对话会话
		chatSessionRouter.PUT("update", api.ApiGroupApp.AiChatSessionApi.UpdateAiChatSession)    // 更新对话会话
		chatSessionRouter.GET("find", api.ApiGroupApp.AiChatSessionApi.FindAiChatSession)        // 根据ID获取对话会话
		chatSessionRouter.GET("getAiChatSessionList", api.ApiGroupApp.AiChatSessionApi.GetAiChatSessionList) // 获取对话会话列表
		chatSessionRouter.GET("getUserSessions", api.ApiGroupApp.AiChatSessionApi.GetUserChatSessions) // 获取用户会话列表
	}
}
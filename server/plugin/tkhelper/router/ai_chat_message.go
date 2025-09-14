package router

import (
	"github.com/gin-gonic/gin"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/tkhelper/api"
)

type AiChatMessageRouter struct{}

func (r *AiChatMessageRouter) InitAiChatMessageRouter(Router *gin.RouterGroup) {
	chatMessageRouter := Router.Group("chat-message")
	{
		chatMessageRouter.POST("create", api.ApiGroupApp.AiChatMessageApi.CreateAiChatMessage)    // 创建对话消息
		chatMessageRouter.DELETE("delete", api.ApiGroupApp.AiChatMessageApi.DeleteAiChatMessage) // 删除对话消息
		chatMessageRouter.DELETE("deleteByIds", api.ApiGroupApp.AiChatMessageApi.DeleteAiChatMessageByIds) // 批量删除对话消息
		chatMessageRouter.PUT("update", api.ApiGroupApp.AiChatMessageApi.UpdateAiChatMessage)    // 更新对话消息
		chatMessageRouter.GET("find", api.ApiGroupApp.AiChatMessageApi.FindAiChatMessage)        // 根据ID获取对话消息
		chatMessageRouter.GET("getAiChatMessageList", api.ApiGroupApp.AiChatMessageApi.GetAiChatMessageList) // 获取对话消息列表
		chatMessageRouter.GET("getSessionMessages", api.ApiGroupApp.AiChatMessageApi.GetSessionMessages) // 获取会话消息
		chatMessageRouter.POST("sendMessage", api.ApiGroupApp.AiChatMessageApi.SendChatMessage) // 发送聊天消息
	}
}
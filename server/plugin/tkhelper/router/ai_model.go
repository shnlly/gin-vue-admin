package router

import (
	"github.com/gin-gonic/gin"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/tkhelper/api"
)

type AiModelRouter struct{}

func (r *AiModelRouter) InitAiModelRouter(Router *gin.RouterGroup) {
	aiModelRouter := Router.Group("ai-model")
	{
		aiModelRouter.POST("create", api.ApiGroupApp.AiModelApi.CreateAiModel)    // 创建AI模型
		aiModelRouter.DELETE("delete", api.ApiGroupApp.AiModelApi.DeleteAiModel) // 删除AI模型
		aiModelRouter.DELETE("deleteByIds", api.ApiGroupApp.AiModelApi.DeleteAiModelByIds) // 批量删除AI模型
		aiModelRouter.PUT("update", api.ApiGroupApp.AiModelApi.UpdateAiModel)    // 更新AI模型
		aiModelRouter.GET("find", api.ApiGroupApp.AiModelApi.FindAiModel)        // 根据ID获取AI模型
		aiModelRouter.GET("getAiModelList", api.ApiGroupApp.AiModelApi.GetAiModelList) // 获取AI模型列表
		aiModelRouter.PUT("setDefault", api.ApiGroupApp.AiModelApi.SetDefaultAiModel) // 设置默认模型
		aiModelRouter.GET("getDefault", api.ApiGroupApp.AiModelApi.GetDefaultAiModel) // 获取默认模型
		aiModelRouter.GET("getEnabled", api.ApiGroupApp.AiModelApi.GetEnabledAiModels) // 获取启用的模型列表
	}
}
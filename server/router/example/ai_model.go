package example

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AiModelRouter struct{}

// InitAiModelRouter 初始化 AI模型 路由信息
func (s *AiModelRouter) InitAiModelRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	aiModelRouter := Router.Group("aiModel").Use(middleware.OperationRecord())
	aiModelRouterWithoutRecord := Router.Group("aiModel")
	aiModelRouterWithoutAuth := PublicRouter.Group("aiModel")

	var aiModelApi = v1.ApiGroupApp.ExampleApiGroup.AiModelApi
	{
		aiModelRouter.POST("createAiModel", aiModelApi.CreateAiModel)             // 新建AI模型
		aiModelRouter.DELETE("deleteAiModel", aiModelApi.DeleteAiModel)           // 删除AI模型
		aiModelRouter.DELETE("deleteAiModelByIds", aiModelApi.DeleteAiModelByIds) // 批量删除AI模型
		aiModelRouter.PUT("updateAiModel", aiModelApi.UpdateAiModel)              // 更新AI模型
		aiModelRouter.POST("setDefault", aiModelApi.SetDefaultAiModel)            // 设置默认AI模型
	}
	{
		aiModelRouterWithoutRecord.GET("findAiModel", aiModelApi.FindAiModel)       // 根据ID获取AI模型
		aiModelRouterWithoutRecord.GET("getAiModelList", aiModelApi.GetAiModelList) // 获取AI模型列表
		aiModelRouterWithoutRecord.GET("getDefault", aiModelApi.GetDefaultAiModel)  // 获取默认AI模型
		aiModelRouterWithoutRecord.GET("getEnabled", aiModelApi.GetEnabledAiModels) // 获取启用的AI模型列表
	}
	{
		aiModelRouterWithoutAuth.GET("getAiModelPublic", aiModelApi.GetAiModelPublic) // AI模型开放接口
	}
}

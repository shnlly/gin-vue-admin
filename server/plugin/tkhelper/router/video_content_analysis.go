package router

import (
	"github.com/gin-gonic/gin"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/tkhelper/api"
)

type VideoContentAnalysisRouter struct{}

func (r *VideoContentAnalysisRouter) InitVideoContentAnalysisRouter(Router *gin.RouterGroup) {
	videoRouter := Router.Group("video-analysis")
	{
		videoRouter.POST("create", api.ApiGroupApp.VideoContentAnalysisApi.CreateVideoContentAnalysis)    // 创建视频分析记录
		videoRouter.DELETE("delete", api.ApiGroupApp.VideoContentAnalysisApi.DeleteVideoContentAnalysis) // 删除视频分析记录
		videoRouter.DELETE("deleteByIds", api.ApiGroupApp.VideoContentAnalysisApi.DeleteVideoContentAnalysisByIds) // 批量删除
		videoRouter.PUT("update", api.ApiGroupApp.VideoContentAnalysisApi.UpdateVideoContentAnalysis)    // 更新视频分析记录
		videoRouter.GET("find", api.ApiGroupApp.VideoContentAnalysisApi.FindVideoContentAnalysis)        // 根据ID获取视频分析记录
		videoRouter.GET("getVideoContentAnalysisList", api.ApiGroupApp.VideoContentAnalysisApi.GetVideoContentAnalysisList) // 获取视频分析记录列表

		// 新增的功能路由
		videoRouter.GET("getVideoContentAnalysisStatistics", api.ApiGroupApp.VideoContentAnalysisApi.GetVideoContentAnalysisStatistics) // 获取统计信息
		videoRouter.POST("importVideoContentAnalysisFromExcel", api.ApiGroupApp.VideoContentAnalysisApi.ImportVideoContentAnalysisFromExcel) // 从Excel导入
		videoRouter.PUT("batchUpdateProcessStatus", api.ApiGroupApp.VideoContentAnalysisApi.BatchUpdateProcessStatus) // 批量更新处理状态
		videoRouter.GET("getVideoContentAnalysisPublic", api.ApiGroupApp.VideoContentAnalysisApi.GetVideoContentAnalysisPublic) // 不鉴权接口
	}
}
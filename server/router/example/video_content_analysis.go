package example

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type VideoContentAnalysisRouter struct{}

func (v *VideoContentAnalysisRouter) InitVideoContentAnalysisRouter(Router *gin.RouterGroup) {
	videoContentRouter := Router.Group("videoContent").Use(middleware.OperationRecord())
	videoContentRouterWithoutRecord := Router.Group("videoContent")
	videoContentRouterWithoutAuth := Router.Group("videoContent")
	{
		videoContentRouter.POST("createVideoContentAnalysis", videoContentAnalysisApi.CreateVideoContentAnalysis)                   // 创建视频内容分析
		videoContentRouter.DELETE("deleteVideoContentAnalysis", videoContentAnalysisApi.DeleteVideoContentAnalysis)                 // 删除视频内容分析
		videoContentRouter.DELETE("deleteVideoContentAnalysisByIds", videoContentAnalysisApi.DeleteVideoContentAnalysisByIds)       // 批量删除
		videoContentRouter.PUT("updateVideoContentAnalysis", videoContentAnalysisApi.UpdateVideoContentAnalysis)                    // 更新视频内容分析
		videoContentRouter.POST("importVideoContentAnalysisFromExcel", videoContentAnalysisApi.ImportVideoContentAnalysisFromExcel) // Excel导入
		videoContentRouter.PUT("batchUpdateProcessStatus", videoContentAnalysisApi.BatchUpdateProcessStatus)                        // 批量更新处理状态
	}
	{
		videoContentRouterWithoutRecord.GET("findVideoContentAnalysis", videoContentAnalysisApi.FindVideoContentAnalysis)                   // 根据ID查询视频内容分析
		videoContentRouterWithoutRecord.GET("getVideoContentAnalysisList", videoContentAnalysisApi.GetVideoContentAnalysisList)             // 分页获取视频内容分析列表
		videoContentRouterWithoutRecord.GET("getVideoContentAnalysisStatistics", videoContentAnalysisApi.GetVideoContentAnalysisStatistics) // 获取统计信息
	}
	{
		videoContentRouterWithoutAuth.GET("getVideoContentAnalysisPublic", videoContentAnalysisApi.GetVideoContentAnalysisPublic) // 不鉴权接口
	}
}

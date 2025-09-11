package example

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
	exampleReq "github.com/flipped-aurora/gin-vue-admin/server/model/example/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type VideoContentAnalysisApi struct{}

var videoContentService = service.ServiceGroupApp.ExampleServiceGroup.VideoContentAnalysisService

// CreateVideoContentAnalysis 创建视频内容分析
// @Tags     VideoContentAnalysis
// @Summary  创建视频内容分析
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     example.VideoContentAnalysis           true "视频链接, 达人名称等"
// @Success  200  {object} response.Response{msg=string} "创建成功"
// @Router   /videoContent/createVideoContentAnalysis [post]
func (vca *VideoContentAnalysisApi) CreateVideoContentAnalysis(c *gin.Context) {
	var videoContent example.VideoContentAnalysis
	err := c.ShouldBindJSON(&videoContent)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = videoContentService.CreateVideoContentAnalysis(&videoContent)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteVideoContentAnalysis 删除视频内容分析
// @Tags     VideoContentAnalysis
// @Summary  删除视频内容分析
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     request.IdsReq                 true "ID"
// @Success  200  {object} response.Response{msg=string} "删除成功"
// @Router   /videoContent/deleteVideoContentAnalysis [delete]
func (vca *VideoContentAnalysisApi) DeleteVideoContentAnalysis(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	err := videoContentService.DeleteVideoContentAnalysis(ID, userID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteVideoContentAnalysisByIds 批量删除视频内容分析
// @Tags     VideoContentAnalysis
// @Summary  批量删除视频内容分析
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Success  200 {object} response.Response{msg=string} "批量删除成功"
// @Router   /videoContent/deleteVideoContentAnalysisByIds [delete]
func (vca *VideoContentAnalysisApi) DeleteVideoContentAnalysisByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	err := videoContentService.DeleteVideoContentAnalysisByIds(IDs, userID)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateVideoContentAnalysis 更新视频内容分析
// @Tags     VideoContentAnalysis
// @Summary  更新视频内容分析
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     example.VideoContentAnalysis           true "视频内容分析信息"
// @Success  200  {object} response.Response{msg=string}  "更新成功"
// @Router   /videoContent/updateVideoContentAnalysis [put]
func (vca *VideoContentAnalysisApi) UpdateVideoContentAnalysis(c *gin.Context) {
	var videoContent example.VideoContentAnalysis
	err := c.ShouldBindJSON(&videoContent)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = videoContentService.UpdateVideoContentAnalysis(videoContent)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindVideoContentAnalysis 用id查询视频内容分析
// @Tags     VideoContentAnalysis
// @Summary  用id查询视频内容分析
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data query    request.IdsReq                                              true "ID"
// @Success  200  {object} response.Response{data=object{videoContent=example.VideoContentAnalysis},msg=string} "查询成功"
// @Router   /videoContent/findVideoContentAnalysis [get]
func (vca *VideoContentAnalysisApi) FindVideoContentAnalysis(c *gin.Context) {
	ID := c.Query("ID")
	videoContent, err := videoContentService.GetVideoContentAnalysis(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithData(gin.H{"videoContent": videoContent}, c)
}

// GetVideoContentAnalysisList 分页获取视频内容分析列表
// @Tags     VideoContentAnalysis
// @Summary  分页获取视频内容分析列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data query    exampleReq.VideoContentAnalysisSearch                       true "页码, 每页大小, 搜索条件"
// @Success  200  {object} response.Response{data=response.PageResult,msg=string}    "获取成功"
// @Router   /videoContent/getVideoContentAnalysisList [get]
func (vca *VideoContentAnalysisApi) GetVideoContentAnalysisList(c *gin.Context) {
	var pageInfo exampleReq.VideoContentAnalysisSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := videoContentService.GetVideoContentAnalysisInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetVideoContentAnalysisPublic 不鉴权的视频内容分析接口
// @Tags     VideoContentAnalysis
// @Summary  不鉴权的视频内容分析接口
// @accept   application/json
// @Produce  application/json
// @Success  200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router   /videoContent/getVideoContentAnalysisPublic [get]
func (vca *VideoContentAnalysisApi) GetVideoContentAnalysisPublic(c *gin.Context) {
	videoContentService.GetVideoContentAnalysisPublic()
	response.OkWithDetailed(gin.H{
		"info": "不鉴权的视频内容分析接口信息",
	}, "获取成功", c)
}

// GetVideoContentAnalysisStatistics 获取视频内容分析统计信息
// @Tags     VideoContentAnalysis
// @Summary  获取视频内容分析统计信息
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Success  200 {object} response.Response{data=exampleRes.VideoContentAnalysisStatistics,msg=string} "获取成功"
// @Router   /videoContent/getVideoContentAnalysisStatistics [get]
func (vca *VideoContentAnalysisApi) GetVideoContentAnalysisStatistics(c *gin.Context) {
	statistics, err := videoContentService.GetVideoContentAnalysisStatistics()
	if err != nil {
		global.GVA_LOG.Error("获取统计信息失败!", zap.Error(err))
		response.FailWithMessage("获取统计信息失败", c)
		return
	}
	response.OkWithDetailed(statistics, "获取成功", c)
}

// ImportVideoContentAnalysisFromExcel 从Excel导入视频内容分析数据
// @Tags     VideoContentAnalysis
// @Summary  从Excel导入视频内容分析数据
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     exampleReq.VideoContentAnalysisImport true "文件路径"
// @Success  200  {object} response.Response{msg=string}  "导入成功"
// @Router   /videoContent/importVideoContentAnalysisFromExcel [post]
func (vca *VideoContentAnalysisApi) ImportVideoContentAnalysisFromExcel(c *gin.Context) {
	var importReq exampleReq.VideoContentAnalysisImport
	err := c.ShouldBindJSON(&importReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	successCount, failCount, err := videoContentService.ImportVideoContentAnalysisFromExcel(importReq.FilePath)
	if err != nil {
		global.GVA_LOG.Error("导入失败!", zap.Error(err))
		response.FailWithMessage("导入失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(gin.H{
		"successCount": successCount,
		"failCount":    failCount,
	}, "导入完成", c)
}

// BatchUpdateProcessStatus 批量更新处理状态
// @Tags     VideoContentAnalysis
// @Summary  批量更新处理状态
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     object{ids=[]uint,statusType=string,status=bool} true "ID列表, 状态类型, 状态值"
// @Success  200  {object} response.Response{msg=string}  "更新成功"
// @Router   /videoContent/batchUpdateProcessStatus [put]
func (vca *VideoContentAnalysisApi) BatchUpdateProcessStatus(c *gin.Context) {
	var req struct {
		IDs        []uint `json:"ids" binding:"required"`
		StatusType string `json:"statusType" binding:"required"`
		Status     bool   `json:"status"`
	}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = videoContentService.BatchUpdateProcessStatus(req.IDs, req.StatusType, req.Status)
	if err != nil {
		global.GVA_LOG.Error("批量更新失败!", zap.Error(err))
		response.FailWithMessage("批量更新失败: "+err.Error(), c)
		return
	}

	response.OkWithMessage("批量更新成功", c)
}

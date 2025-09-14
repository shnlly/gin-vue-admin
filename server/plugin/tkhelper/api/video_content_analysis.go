package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/gin-gonic/gin"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/tkhelper/model"
	pluginReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/tkhelper/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/tkhelper/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"go.uber.org/zap"
)

type VideoContentAnalysisApi struct{}

var videoContentAnalysisService = service.ServiceGroup{}.VideoContentAnalysisService

// CreateVideoContentAnalysis 创建视频内容分析记录
func (v *VideoContentAnalysisApi) CreateVideoContentAnalysis(c *gin.Context) {
	var videoContent model.VideoContentAnalysis
	err := c.ShouldBindJSON(&videoContent)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = videoContentAnalysisService.CreateVideoContentAnalysis(&videoContent)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteVideoContentAnalysis 删除视频内容分析记录
func (v *VideoContentAnalysisApi) DeleteVideoContentAnalysis(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	err := videoContentAnalysisService.DeleteVideoContentAnalysis(ID, userID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteVideoContentAnalysisByIds 批量删除视频内容分析记录
func (v *VideoContentAnalysisApi) DeleteVideoContentAnalysisByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	err := videoContentAnalysisService.DeleteVideoContentAnalysisByIds(IDs, userID)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateVideoContentAnalysis 更新视频内容分析记录
func (v *VideoContentAnalysisApi) UpdateVideoContentAnalysis(c *gin.Context) {
	var videoContent model.VideoContentAnalysis
	err := c.ShouldBindJSON(&videoContent)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = videoContentAnalysisService.UpdateVideoContentAnalysis(videoContent)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindVideoContentAnalysis 根据ID获取视频内容分析记录
func (v *VideoContentAnalysisApi) FindVideoContentAnalysis(c *gin.Context) {
	ID := c.Query("ID")
	videoContent, err := videoContentAnalysisService.GetVideoContentAnalysis(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithData(gin.H{"videoContent": videoContent}, c)
}

// GetVideoContentAnalysisList 分页获取视频内容分析记录列表
func (v *VideoContentAnalysisApi) GetVideoContentAnalysisList(c *gin.Context) {
	var pageInfo pluginReq.VideoContentAnalysisSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := videoContentAnalysisService.GetVideoContentAnalysisInfoList(pageInfo)
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
func (v *VideoContentAnalysisApi) GetVideoContentAnalysisPublic(c *gin.Context) {
	videoContentAnalysisService.GetVideoContentAnalysisPublic()
	response.OkWithDetailed(gin.H{
		"info": "不鉴权的视频内容分析接口信息",
	}, "获取成功", c)
}

// GetVideoContentAnalysisStatistics 获取视频内容分析统计信息
func (v *VideoContentAnalysisApi) GetVideoContentAnalysisStatistics(c *gin.Context) {
	statistics, err := videoContentAnalysisService.GetVideoContentAnalysisStatistics()
	if err != nil {
		global.GVA_LOG.Error("获取统计信息失败!", zap.Error(err))
		response.FailWithMessage("获取统计信息失败", c)
		return
	}
	response.OkWithDetailed(statistics, "获取成功", c)
}

// ImportVideoContentAnalysisFromExcel 从Excel导入视频内容分析数据
func (v *VideoContentAnalysisApi) ImportVideoContentAnalysisFromExcel(c *gin.Context) {
	var importReq pluginReq.VideoContentAnalysisImport
	err := c.ShouldBindJSON(&importReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	successCount, failCount, err := videoContentAnalysisService.ImportVideoContentAnalysisFromExcel(importReq.FilePath)
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
func (v *VideoContentAnalysisApi) BatchUpdateProcessStatus(c *gin.Context) {
	var req pluginReq.BatchUpdateProcessStatusRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = videoContentAnalysisService.BatchUpdateProcessStatus(req.IDs, req.StatusType, req.Status)
	if err != nil {
		global.GVA_LOG.Error("批量更新失败!", zap.Error(err))
		response.FailWithMessage("批量更新失败: "+err.Error(), c)
		return
	}

	response.OkWithMessage("批量更新成功", c)
}
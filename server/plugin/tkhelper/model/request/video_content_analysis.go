package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/tkhelper/model"
)

// VideoContentAnalysisSearch 视频内容分析搜索条件
type VideoContentAnalysisSearch struct {
	model.VideoContentAnalysis
	StartCreatedAt *string `json:"startCreatedAt" form:"startCreatedAt"` // 创建时间开始
	EndCreatedAt   *string `json:"endCreatedAt" form:"endCreatedAt"`     // 创建时间结束
	CreatorName    string  `json:"creatorName" form:"creatorName"`       // 达人名称
	VideoTitle     string  `json:"videoTitle" form:"videoTitle"`         // 视频标题
	StartPlayCount string  `json:"startPlayCount" form:"startPlayCount"` // 播放量开始
	EndPlayCount   string  `json:"endPlayCount" form:"endPlayCount"`     // 播放量结束
	IsDownloaded   *bool   `json:"isDownloaded" form:"isDownloaded"`     // 是否已下载
	request.PageInfo
}

// VideoContentAnalysisImport 批量导入请求
type VideoContentAnalysisImport struct {
	FilePath string `json:"filePath" form:"filePath" binding:"required"` // 文件路径
}

// BatchUpdateProcessStatusRequest 批量更新处理状态请求
type BatchUpdateProcessStatusRequest struct {
	IDs        []uint `json:"ids" binding:"required"`
	StatusType string `json:"statusType" binding:"required"`
	Status     bool   `json:"status"`
}
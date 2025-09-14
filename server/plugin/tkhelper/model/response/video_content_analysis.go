package response

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/tkhelper/model"

// VideoContentAnalysisResponse 视频内容分析响应
type VideoContentAnalysisResponse struct {
	VideoContent model.VideoContentAnalysis `json:"videoContent"`
}

// VideoContentAnalysisStatistics 视频内容分析统计响应
type VideoContentAnalysisStatistics struct {
	TotalCount       int64   `json:"totalCount"`       // 总视频数
	DownloadedCount  int64   `json:"downloadedCount"`  // 已下载数量
	AnalyzedCount    int64   `json:"analyzedCount"`    // 已分析数量
	TotalPlayCount   string  `json:"totalPlayCount"`   // 总播放量
	TotalLikeCount   string  `json:"totalLikeCount"`   // 总点赞数
	TotalSalesAmount string  `json:"totalSalesAmount"` // 总销售额
	TopCreators      []struct {
		CreatorName string `json:"creatorName"`
		VideoCount  int64  `json:"videoCount"`
		TotalPlay   string `json:"totalPlay"`
	} `json:"topCreators"` // 头部达人统计
}
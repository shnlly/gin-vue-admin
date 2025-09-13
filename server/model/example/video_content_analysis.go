package example

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// VideoContentAnalysis 视频内容聚合分析
type VideoContentAnalysis struct {
	global.GVA_MODEL
	VideoDescription string `json:"videoDescription" form:"videoDescription" gorm:"column:video_description;comment:视频描述内容;type:text"`       // 视频描述
	VideoLink        string `json:"videoLink" form:"videoLink" gorm:"column:video_link;comment:视频链接地址;size:500;not null"`                    // 视频链接
	CreatorName      string `json:"creatorName" form:"creatorName" gorm:"column:creator_name;comment:视频创作者达人名称;size:100;not null"`           // 达人名称
	UniqueId         string `json:"uniqueId" form:"uniqueId" gorm:"column:unique_id;comment:视频创作者唯一标识;size:100"`                             // 唯一标识
	FansCount        string `json:"fansCount" form:"fansCount" gorm:"column:fans_count;comment:达人粉丝数量;size:20"`                              // 粉丝数
	PlayCount        string `json:"playCount" form:"playCount" gorm:"column:play_count;comment:视频播放量;size:20"`                               // 播放量
	LikeCount        string `json:"likeCount" form:"likeCount" gorm:"column:like_count;comment:视频点赞数;size:20"`                               // 点赞数
	CommentCount     int    `json:"commentCount" form:"commentCount" gorm:"column:comment_count;comment:视频评论数;default:0"`                    // 评论数
	ShareCount       int    `json:"shareCount" form:"shareCount" gorm:"column:share_count;comment:视频转发数;default:0"`                          // 转发数
	SalesVolume      int    `json:"salesVolume" form:"salesVolume" gorm:"column:sales_volume;comment:商品销量件数;default:0"`                      // 销量(件)
	SalesAmount      string `json:"salesAmount" form:"salesAmount" gorm:"column:sales_amount;comment:商品销售金额;size:20"`                        // 销售额
	PublishTime      string `json:"publishTime" form:"publishTime" gorm:"column:publish_time;comment:视频发布时间;size:50"`                        // 发布时间
	VideoDuration    string `json:"videoDuration" form:"videoDuration" gorm:"column:video_duration;comment:视频播放时长;size:20"`                  // 视频时长
	VideoTitle       string `json:"videoTitle" form:"videoTitle" gorm:"column:video_title;comment:视频标题;size:500"`                            // 视频标题
	VideoScript      string `json:"videoScript" form:"videoScript" gorm:"column:video_script;comment:原始视频文案内容;type:text"`                    // 视频文案
	FixedScript      string `json:"fixedScript" form:"fixedScript" gorm:"column:fixed_script;comment:修复后的视频文案内容;type:text"`                  // 修复后文案
	VideoAnalysis    string `json:"videoAnalysis" form:"videoAnalysis" gorm:"column:video_analysis;comment:视频画面分析内容;type:text"`              // 视频画面分析
	IsDownloaded     *bool  `json:"isDownloaded" form:"isDownloaded" gorm:"column:is_downloaded;comment:是否已下载视频;default:false"`              // 是否下载
	IsAudioExtracted *bool  `json:"isAudioExtracted" form:"isAudioExtracted" gorm:"column:is_audio_extracted;comment:是否已提取音频;default:false"` // 是否提取音频
	IsTextConverted  *bool  `json:"isTextConverted" form:"isTextConverted" gorm:"column:is_text_converted;comment:是否已转为文本;default:false"`    // 是否转文本
	IsScriptFixed    *bool  `json:"isScriptFixed" form:"isScriptFixed" gorm:"column:is_script_fixed;comment:是否已修复文案;default:false"`          // 是否修复文案
	IsVideoAnalyzed  *bool  `json:"isVideoAnalyzed" form:"isVideoAnalyzed" gorm:"column:is_video_analyzed;comment:是否已分析视频画面;default:false"`  // 是否分析视频画面
	FixNote          string `json:"fixNote" form:"fixNote" gorm:"column:fix_note;comment:修复说明备注;size:500"`                                   // 修复说明
}

// TableName 设置表名
func (VideoContentAnalysis) TableName() string {
	return "video_content_analysis"
}

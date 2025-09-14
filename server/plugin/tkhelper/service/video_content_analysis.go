package service

import (
	"errors"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/tkhelper/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/tkhelper/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/tkhelper/model/response"
	"gorm.io/gorm"
)

type VideoContentAnalysisService struct{}

var excelImportService = new(ExcelImportService)

// CreateVideoContentAnalysis 创建视频内容分析记录
func (vca *VideoContentAnalysisService) CreateVideoContentAnalysis(videoContent *model.VideoContentAnalysis) (err error) {
	err = global.GVA_DB.Create(videoContent).Error
	return err
}

// DeleteVideoContentAnalysis 删除视频内容分析记录
func (vca *VideoContentAnalysisService) DeleteVideoContentAnalysis(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.VideoContentAnalysis{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&model.VideoContentAnalysis{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteVideoContentAnalysisByIds 批量删除视频内容分析记录
func (vca *VideoContentAnalysisService) DeleteVideoContentAnalysisByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.VideoContentAnalysis{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&model.VideoContentAnalysis{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateVideoContentAnalysis 更新视频内容分析记录
func (vca *VideoContentAnalysisService) UpdateVideoContentAnalysis(videoContent model.VideoContentAnalysis) (err error) {
	err = global.GVA_DB.Model(&model.VideoContentAnalysis{}).Where("id = ?", videoContent.ID).Updates(&videoContent).Error
	return err
}

// GetVideoContentAnalysis 根据ID获取视频内容分析记录
func (vca *VideoContentAnalysisService) GetVideoContentAnalysis(ID string) (videoContent model.VideoContentAnalysis, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&videoContent).Error
	return
}

// GetVideoContentAnalysisInfoList 分页获取视频内容分析记录
func (vca *VideoContentAnalysisService) GetVideoContentAnalysisInfoList(info request.VideoContentAnalysisSearch) (list []model.VideoContentAnalysis, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.VideoContentAnalysis{})
	var videoContents []model.VideoContentAnalysis

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", *info.StartCreatedAt, *info.EndCreatedAt)
	}
	if info.CreatorName != "" {
		db = db.Where("creator_name LIKE ?", "%"+info.CreatorName+"%")
	}
	if info.VideoTitle != "" {
		db = db.Where("video_title LIKE ?", "%"+info.VideoTitle+"%")
	}
	if info.StartPlayCount != "" && info.EndPlayCount != "" {
		db = db.Where("CAST(REPLACE(REPLACE(play_count, 'K', '*1000'), 'M', '*1000000') AS DECIMAL) BETWEEN ? AND ?", info.StartPlayCount, info.EndPlayCount)
	}
	if info.IsDownloaded != nil {
		db = db.Where("is_downloaded = ?", *info.IsDownloaded)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&videoContents).Error
	return videoContents, total, err
}

// GetVideoContentAnalysisPublic 不鉴权的视频内容分析接口
func (vca *VideoContentAnalysisService) GetVideoContentAnalysisPublic() {
	// 此方法为空实现，供公共访问
}

// GetVideoContentAnalysisStatistics 获取视频内容分析统计信息
func (vca *VideoContentAnalysisService) GetVideoContentAnalysisStatistics() (statistics response.VideoContentAnalysisStatistics, err error) {
	db := global.GVA_DB.Model(&model.VideoContentAnalysis{})

	// 总数统计
	err = db.Count(&statistics.TotalCount).Error
	if err != nil {
		return statistics, err
	}

	// 已下载数量
	err = db.Where("is_downloaded = ?", true).Count(&statistics.DownloadedCount).Error
	if err != nil {
		return statistics, err
	}

	// 已分析数量
	err = db.Where("is_video_analyzed = ?", true).Count(&statistics.AnalyzedCount).Error
	if err != nil {
		return statistics, err
	}

	// 头部达人统计（前10名）
	type CreatorStat struct {
		CreatorName string `json:"creator_name"`
		VideoCount  int64  `json:"video_count"`
		TotalPlay   string `json:"total_play"`
	}

	var topCreators []CreatorStat
	err = db.Select("creator_name, COUNT(*) as video_count, SUM(CAST(REPLACE(REPLACE(play_count, 'K', '*1000'), 'M', '*1000000') AS DECIMAL)) as total_play").
		Where("creator_name != ''").
		Group("creator_name").
		Order("video_count DESC").
		Limit(10).
		Scan(&topCreators).Error

	if err != nil {
		return statistics, err
	}

	for _, creator := range topCreators {
		statistics.TopCreators = append(statistics.TopCreators, struct {
			CreatorName string `json:"creatorName"`
			VideoCount  int64  `json:"videoCount"`
			TotalPlay   string `json:"totalPlay"`
		}{
			CreatorName: creator.CreatorName,
			VideoCount:  creator.VideoCount,
			TotalPlay:   creator.TotalPlay,
		})
	}

	return statistics, nil
}

// ImportVideoContentAnalysisFromExcel 从Excel导入视频内容分析数据
func (vca *VideoContentAnalysisService) ImportVideoContentAnalysisFromExcel(filePath string) (successCount int, failCount int, err error) {
	// 先验证Excel文件格式
	if err := excelImportService.ValidateExcelFormat(filePath); err != nil {
		return 0, 0, fmt.Errorf("Excel文件格式验证失败: %v", err)
	}

	// 调用Excel导入服务
	return excelImportService.ImportVideoContentFromExcel(filePath)
}

// BatchUpdateProcessStatus 批量更新处理状态
func (vca *VideoContentAnalysisService) BatchUpdateProcessStatus(ids []uint, statusType string, status bool) (err error) {
	if len(ids) == 0 {
		return errors.New("未选择任何记录")
	}

	updateData := make(map[string]interface{})

	switch statusType {
	case "downloaded":
		updateData["is_downloaded"] = status
	case "audio_extracted":
		updateData["is_audio_extracted"] = status
	case "text_converted":
		updateData["is_text_converted"] = status
	case "script_fixed":
		updateData["is_script_fixed"] = status
	case "video_analyzed":
		updateData["is_video_analyzed"] = status
	default:
		return errors.New("不支持的状态类型")
	}

	err = global.GVA_DB.Model(&model.VideoContentAnalysis{}).Where("id IN ?", ids).Updates(updateData).Error
	return err
}
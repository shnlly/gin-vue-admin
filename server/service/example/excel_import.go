package example

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
	"github.com/xuri/excelize/v2"
)

type ExcelImportService struct{}

var ExcelImportServiceApp = new(ExcelImportService)

// ImportVideoContentFromExcel 从Excel文件导入视频内容分析数据
func (e *ExcelImportService) ImportVideoContentFromExcel(filePath string) (successCount int, failCount int, err error) {
	if filePath == "" {
		return 0, 0, errors.New("文件路径不能为空")
	}

	// 打开Excel文件
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return 0, 0, fmt.Errorf("打开Excel文件失败: %v", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			global.GVA_LOG.Error("关闭Excel文件失败: " + err.Error())
		}
	}()

	// 获取工作表名称
	sheets := f.GetSheetList()
	if len(sheets) == 0 {
		return 0, 0, errors.New("Excel文件中没有工作表")
	}

	// 使用第一个工作表
	sheetName := sheets[0]
	rows, err := f.GetRows(sheetName)
	if err != nil {
		return 0, 0, fmt.Errorf("读取工作表数据失败: %v", err)
	}

	if len(rows) <= 1 {
		return 0, 0, errors.New("Excel文件中没有数据行")
	}

	// 第一行是标题行，从第二行开始处理数据
	headerRow := rows[0]

	// 创建字段映射表 (Excel列名 -> 数组索引)
	fieldMap := e.createFieldMapping(headerRow)

	successCount = 0
	failCount = 0

	// 处理数据行
	for i := 1; i < len(rows); i++ {
		row := rows[i]
		global.GVA_LOG.Info(fmt.Sprintf("正在处理第%d行数据，共%d行", i+1, len(rows)))

		videoContent, err := e.parseRowToVideoContent(row, fieldMap)
		if err != nil {
			global.GVA_LOG.Error(fmt.Sprintf("解析第%d行数据失败: %v", i+1, err))
			failCount++
			continue
		}

		// 检查是否已存在相同的视频链接
		var existingRecord example.VideoContentAnalysis
		if err := global.GVA_DB.Where("video_link = ?", videoContent.VideoLink).First(&existingRecord).Error; err == nil {
			global.GVA_LOG.Warn(fmt.Sprintf("第%d行数据已存在相同的视频链接，跳过: %s", i+1, videoContent.VideoLink))
			failCount++
			continue
		}

		// 插入数据库
		if err := global.GVA_DB.Create(&videoContent).Error; err != nil {
			global.GVA_LOG.Error(fmt.Sprintf("插入第%d行数据失败: %v", i+1, err))
			global.GVA_LOG.Error(fmt.Sprintf("失败的数据: VideoLink=%s, CreatorName=%s", videoContent.VideoLink, videoContent.CreatorName))
			failCount++
			continue
		}

		global.GVA_LOG.Info(fmt.Sprintf("第%d行数据插入成功", i+1))
		successCount++
	}

	return successCount, failCount, nil
}

// createFieldMapping 创建Excel字段映射
func (e *ExcelImportService) createFieldMapping(headerRow []string) map[string]int {
	fieldMap := make(map[string]int)

	// 定义Excel列名与数据库字段的映射关系
	fieldMapping := map[string]string{
		"视频描述":      "video_description",
		"视频链接":      "video_link",
		"达人名称":      "creator_name",
		"Unique Id": "unique_id",
		"粉丝数":       "fans_count",
		"播放量":       "play_count",
		"点赞数":       "like_count",
		"评论数":       "comment_count",
		"转发数":       "share_count",
		"销量（件）":     "sales_volume",
		"销售额":       "sales_amount",
		"发布时间":      "publish_time",
		"视频时长":      "video_duration",
		"视频标题":      "video_title",
		"视频文案":      "video_script",
		"修复后文案":     "fixed_script",
		"视频画面分析":    "video_analysis",
		"是否下载":      "is_downloaded",
		"是否提取音频":    "is_audio_extracted",
		"是否转文本":     "is_text_converted",
		"是否修复文案":    "is_script_fixed",
		"是否分析视频画面":  "is_video_analyzed",
		"修复说明":      "fix_note",
	}

	// 建立索引映射
	for i, header := range headerRow {
		if dbField, exists := fieldMapping[header]; exists {
			fieldMap[dbField] = i
		}
	}

	return fieldMap
}

// parseRowToVideoContent 将Excel行数据解析为VideoContentAnalysis对象
func (e *ExcelImportService) parseRowToVideoContent(row []string, fieldMap map[string]int) (example.VideoContentAnalysis, error) {
	var videoContent example.VideoContentAnalysis

	// 获取字段值的辅助函数
	getFieldValue := func(fieldName string) string {
		if index, exists := fieldMap[fieldName]; exists && index < len(row) {
			return strings.TrimSpace(row[index])
		}
		return ""
	}

	// 布尔值解析辅助函数
	parseBool := func(value string) *bool {
		value = strings.TrimSpace(value)
		if value == "" {
			result := false
			return &result
		}

		// 处理中文和英文的是/否
		lowerValue := strings.ToLower(value)
		switch lowerValue {
		case "是", "yes", "true", "1", "已下载", "已提取", "已转换", "已修复", "已分析":
			result := true
			return &result
		case "否", "no", "false", "0", "未下载", "未提取", "未转换", "未修复", "未分析":
			result := false
			return &result
		default:
			result := false
			return &result
		}
	}

	// 整数解析辅助函数
	parseInt := func(value string) int {
		if value == "" {
			return 0
		}
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
		return 0
	}

	// 必填字段验证
	creatorName := getFieldValue("creator_name")
	videoLink := getFieldValue("video_link")

	if creatorName == "" {
		return videoContent, errors.New("达人名称不能为空")
	}
	if videoLink == "" {
		return videoContent, errors.New("视频链接不能为空")
	}

	// 填充数据
	videoContent.VideoDescription = getFieldValue("video_description")
	videoContent.VideoLink = videoLink
	videoContent.CreatorName = creatorName
	videoContent.UniqueId = getFieldValue("unique_id")
	videoContent.FansCount = getFieldValue("fans_count")
	videoContent.PlayCount = getFieldValue("play_count")
	videoContent.LikeCount = getFieldValue("like_count")
	videoContent.CommentCount = parseInt(getFieldValue("comment_count"))
	videoContent.ShareCount = parseInt(getFieldValue("share_count"))
	videoContent.SalesVolume = parseInt(getFieldValue("sales_volume"))
	videoContent.SalesAmount = getFieldValue("sales_amount")
	videoContent.PublishTime = getFieldValue("publish_time")
	videoContent.VideoDuration = getFieldValue("video_duration")
	videoContent.VideoTitle = getFieldValue("video_title")
	videoContent.VideoScript = getFieldValue("video_script")
	videoContent.FixedScript = getFieldValue("fixed_script")
	videoContent.VideoAnalysis = getFieldValue("video_analysis")
	videoContent.IsDownloaded = parseBool(getFieldValue("is_downloaded"))
	videoContent.IsAudioExtracted = parseBool(getFieldValue("is_audio_extracted"))
	videoContent.IsTextConverted = parseBool(getFieldValue("is_text_converted"))
	videoContent.IsScriptFixed = parseBool(getFieldValue("is_script_fixed"))
	videoContent.IsVideoAnalyzed = parseBool(getFieldValue("is_video_analyzed"))
	videoContent.FixNote = getFieldValue("fix_note")

	return videoContent, nil
}

// ValidateExcelFormat 验证Excel文件格式
func (e *ExcelImportService) ValidateExcelFormat(filePath string) error {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return fmt.Errorf("无法打开Excel文件: %v", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			global.GVA_LOG.Error("关闭Excel文件失败: " + err.Error())
		}
	}()

	sheets := f.GetSheetList()
	if len(sheets) == 0 {
		return errors.New("Excel文件中没有工作表")
	}

	rows, err := f.GetRows(sheets[0])
	if err != nil {
		return fmt.Errorf("读取工作表失败: %v", err)
	}

	if len(rows) <= 1 {
		return errors.New("Excel文件中没有数据行")
	}

	// 验证必要的列是否存在
	headerRow := rows[0]
	requiredColumns := []string{"达人名称", "视频链接"}

	for _, required := range requiredColumns {
		found := false
		for _, header := range headerRow {
			if strings.TrimSpace(header) == required {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("缺少必要的列: %s", required)
		}
	}

	return nil
}

package initialize

import (
	"context"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

func Api(ctx context.Context) {
	entities := []model.SysApi{
		// AI模型管理 API
		{
			Path:        "/tkhelper/ai-model/create",
			Description: "新建AI模型",
			ApiGroup:    "AI模型管理",
			Method:      "POST",
		},
		{
			Path:        "/tkhelper/ai-model/delete",
			Description: "删除AI模型",
			ApiGroup:    "AI模型管理",
			Method:      "DELETE",
		},
		{
			Path:        "/tkhelper/ai-model/deleteByIds",
			Description: "批量删除AI模型",
			ApiGroup:    "AI模型管理",
			Method:      "DELETE",
		},
		{
			Path:        "/tkhelper/ai-model/update",
			Description: "更新AI模型",
			ApiGroup:    "AI模型管理",
			Method:      "PUT",
		},
		{
			Path:        "/tkhelper/ai-model/find",
			Description: "根据ID获取AI模型",
			ApiGroup:    "AI模型管理",
			Method:      "GET",
		},
		{
			Path:        "/tkhelper/ai-model/getAiModelList",
			Description: "获取AI模型列表",
			ApiGroup:    "AI模型管理",
			Method:      "GET",
		},
		{
			Path:        "/tkhelper/ai-model/setDefault",
			Description: "设置默认AI模型",
			ApiGroup:    "AI模型管理",
			Method:      "PUT",
		},
		{
			Path:        "/tkhelper/ai-model/getDefault",
			Description: "获取默认AI模型",
			ApiGroup:    "AI模型管理",
			Method:      "GET",
		},
		{
			Path:        "/tkhelper/ai-model/getEnabled",
			Description: "获取启用的AI模型列表",
			ApiGroup:    "AI模型管理",
			Method:      "GET",
		},

		// AI对话会话管理 API
		{
			Path:        "/tkhelper/ai-chat-session/create",
			Description: "新建AI对话会话",
			ApiGroup:    "AI对话管理",
			Method:      "POST",
		},
		{
			Path:        "/tkhelper/ai-chat-session/delete",
			Description: "删除AI对话会话",
			ApiGroup:    "AI对话管理",
			Method:      "DELETE",
		},
		{
			Path:        "/tkhelper/ai-chat-session/deleteByIds",
			Description: "批量删除AI对话会话",
			ApiGroup:    "AI对话管理",
			Method:      "DELETE",
		},
		{
			Path:        "/tkhelper/ai-chat-session/update",
			Description: "更新AI对话会话",
			ApiGroup:    "AI对话管理",
			Method:      "PUT",
		},
		{
			Path:        "/tkhelper/ai-chat-session/find",
			Description: "根据ID获取AI对话会话",
			ApiGroup:    "AI对话管理",
			Method:      "GET",
		},
		{
			Path:        "/tkhelper/ai-chat-session/getAiChatSessionList",
			Description: "获取AI对话会话列表",
			ApiGroup:    "AI对话管理",
			Method:      "GET",
		},
		{
			Path:        "/tkhelper/ai-chat-session/getUserSessions",
			Description: "获取用户会话列表",
			ApiGroup:    "AI对话管理",
			Method:      "GET",
		},

		// AI对话消息管理 API
		{
			Path:        "/tkhelper/ai-chat-message/create",
			Description: "新建AI对话消息",
			ApiGroup:    "AI对话管理",
			Method:      "POST",
		},
		{
			Path:        "/tkhelper/ai-chat-message/delete",
			Description: "删除AI对话消息",
			ApiGroup:    "AI对话管理",
			Method:      "DELETE",
		},
		{
			Path:        "/tkhelper/ai-chat-message/deleteByIds",
			Description: "批量删除AI对话消息",
			ApiGroup:    "AI对话管理",
			Method:      "DELETE",
		},
		{
			Path:        "/tkhelper/ai-chat-message/update",
			Description: "更新AI对话消息",
			ApiGroup:    "AI对话管理",
			Method:      "PUT",
		},
		{
			Path:        "/tkhelper/ai-chat-message/find",
			Description: "根据ID获取AI对话消息",
			ApiGroup:    "AI对话管理",
			Method:      "GET",
		},
		{
			Path:        "/tkhelper/ai-chat-message/getAiChatMessageList",
			Description: "获取AI对话消息列表",
			ApiGroup:    "AI对话管理",
			Method:      "GET",
		},
		{
			Path:        "/tkhelper/ai-chat-message/getSessionMessages",
			Description: "获取会话消息",
			ApiGroup:    "AI对话管理",
			Method:      "GET",
		},
		{
			Path:        "/tkhelper/ai-chat-message/sendMessage",
			Description: "发送聊天消息",
			ApiGroup:    "AI对话管理",
			Method:      "POST",
		},

		// 视频内容分析 API
		{
			Path:        "/tkhelper/video-content-analysis/create",
			Description: "新建视频内容分析",
			ApiGroup:    "视频内容分析",
			Method:      "POST",
		},
		{
			Path:        "/tkhelper/video-content-analysis/delete",
			Description: "删除视频内容分析",
			ApiGroup:    "视频内容分析",
			Method:      "DELETE",
		},
		{
			Path:        "/tkhelper/video-content-analysis/deleteByIds",
			Description: "批量删除视频内容分析",
			ApiGroup:    "视频内容分析",
			Method:      "DELETE",
		},
		{
			Path:        "/tkhelper/video-content-analysis/update",
			Description: "更新视频内容分析",
			ApiGroup:    "视频内容分析",
			Method:      "PUT",
		},
		{
			Path:        "/tkhelper/video-content-analysis/find",
			Description: "根据ID获取视频内容分析",
			ApiGroup:    "视频内容分析",
			Method:      "GET",
		},
		{
			Path:        "/tkhelper/video-content-analysis/getVideoContentAnalysisList",
			Description: "获取视频内容分析列表",
			ApiGroup:    "视频内容分析",
			Method:      "GET",
		},
		{
			Path:        "/tkhelper/video-analysis/getVideoContentAnalysisStatistics",
			Description: "获取视频内容分析统计信息",
			ApiGroup:    "视频内容分析",
			Method:      "GET",
		},
		{
			Path:        "/tkhelper/video-analysis/importVideoContentAnalysisFromExcel",
			Description: "导入视频内容分析",
			ApiGroup:    "视频内容分析",
			Method:      "POST",
		},
		{
			Path:        "/tkhelper/video-analysis/batchUpdateProcessStatus",
			Description: "批量更新处理状态",
			ApiGroup:    "视频内容分析",
			Method:      "PUT",
		},
		{
			Path:        "/tkhelper/video-analysis/getVideoContentAnalysisPublic",
			Description: "获取视频内容分析公开接口",
			ApiGroup:    "视频内容分析",
			Method:      "GET",
		},
	}
	utils.RegisterApis(entities...)
}
package initialize

import (
	"context"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

func Menu(ctx context.Context) {
	entities := []model.SysBaseMenu{
		{
			ParentId:  9, // 插件系统父菜单ID
			Path:      "videoContentAnalysis",
			Name:      "videoContentAnalysis",
			Hidden:    false,
			Component: "plugin/tkhelper/view/videoContentAnalysis.vue",
			Sort:      101,
			Meta: model.Meta{
				Title:     "视频内容分析",
				Icon:      "video-camera",
				KeepAlive: true,
			},
			MenuBtn: []model.SysBaseMenuBtn{
				{Name: "add", Desc: "新增"},
				{Name: "edit", Desc: "编辑"},
				{Name: "delete", Desc: "删除"},
				{Name: "view", Desc: "查看"},
				{Name: "import", Desc: "导入"},
				{Name: "statistics", Desc: "统计"},
			},
		},
		{
			ParentId:  9, // 插件系统父菜单ID
			Path:      "aiModel",
			Name:      "aiModel",
			Hidden:    false,
			Component: "plugin/tkhelper/view/aiModel.vue",
			Sort:      102,
			Meta: model.Meta{
				Title:     "AI模型管理",
				Icon:      "cpu",
				KeepAlive: false,
			},
			MenuBtn: []model.SysBaseMenuBtn{
				{Name: "add", Desc: "新增"},
				{Name: "edit", Desc: "编辑"},
				{Name: "delete", Desc: "删除"},
				{Name: "view", Desc: "查看"},
				{Name: "setDefault", Desc: "设为默认"},
			},
		},
		{
			ParentId:  9, // 插件系统父菜单ID
			Path:      "aiChat",
			Name:      "aiChat",
			Hidden:    false,
			Component: "plugin/tkhelper/view/aiChat.vue",
			Sort:      103,
			Meta: model.Meta{
				Title:     "AI对话",
				Icon:      "chat-dot-round",
				KeepAlive: true,
			},
			MenuBtn: []model.SysBaseMenuBtn{
				{Name: "chat", Desc: "发起对话"},
				{Name: "history", Desc: "历史会话"},
				{Name: "model", Desc: "选择模型"},
			},
		},
	}
	utils.RegisterMenus(entities...)
}

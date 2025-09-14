package initialize

import (
	"context"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

func Menu(ctx context.Context) {
	entities := []model.SysBaseMenu{
		// 主菜单：TK助手
		{
			ParentId:  0, // 根级菜单
			Path:      "tkhelper",
			Name:      "tkhelper",
			Hidden:    false,
			Component: "view/routerHolder.vue",
			Sort:      10,
			Meta: model.Meta{
				Title: "TK助手",
				Icon:  "tools",
			},
		},
		// 子菜单1：视频内容分析
		{
			ParentId:  0, // 这里会在插件注册时自动设置为上面主菜单的ID
			Path:      "videoContentAnalysis",
			Name:      "videoContentAnalysis",
			Hidden:    false,
			Component: "plugin/tkhelper/view/videoContentAnalysis.vue",
			Sort:      1,
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
		// 子菜单2：AI模型管理
		{
			ParentId:  0, // 这里会在插件注册时自动设置为主菜单的ID
			Path:      "aiModel",
			Name:      "aiModel",
			Hidden:    false,
			Component: "plugin/tkhelper/view/aiModel.vue",
			Sort:      2,
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
		// 子菜单3：AI对话
		{
			ParentId:  0, // 这里会在插件注册时自动设置为主菜单的ID
			Path:      "aiChat",
			Name:      "aiChat",
			Hidden:    false,
			Component: "plugin/tkhelper/view/aiChat.vue",
			Sort:      3,
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

	// utils.RegisterMenus 会自动处理父子关系：
	// - 第一个菜单作为父菜单
	// - 后续菜单自动设置为第一个菜单的子菜单
	utils.RegisterMenus(entities...)
}

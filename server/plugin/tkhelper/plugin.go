package tkhelper

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/tkhelper/initialize"
	interfaces "github.com/flipped-aurora/gin-vue-admin/server/utils/plugin/v2"
	"github.com/gin-gonic/gin"
)

var _ interfaces.Plugin = (*plugin)(nil)

var Plugin = new(plugin)

type plugin struct{}

func (p *plugin) Register(group *gin.Engine) {
	ctx := context.Background()
	// 安装插件时候自动注册的api数据请到下方法.Api方法中实现
	initialize.Api(ctx)
	// 安装插件时候自动注册的菜单数据请到下方法.Menu方法中实现
	initialize.Menu(ctx)
	// 初始化数据库
	initialize.Gorm(ctx)
	// 注册路由
	initialize.Router(group)
}
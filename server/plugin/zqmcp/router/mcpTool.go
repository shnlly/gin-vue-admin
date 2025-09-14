package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var McpTool = new(mcpTool)

type mcpTool struct {}

// Init 初始化 MCP工具管理 路由信息
func (r *mcpTool) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
	    group := private.Group("mcpTool").Use(middleware.OperationRecord())
		group.POST("createMcpTool", apiMcpTool.CreateMcpTool)   // 新建MCP工具管理
		group.DELETE("deleteMcpTool", apiMcpTool.DeleteMcpTool) // 删除MCP工具管理
		group.DELETE("deleteMcpToolByIds", apiMcpTool.DeleteMcpToolByIds) // 批量删除MCP工具管理
		group.PUT("updateMcpTool", apiMcpTool.UpdateMcpTool)    // 更新MCP工具管理
	}
	{
	    group := private.Group("mcpTool")
		group.GET("findMcpTool", apiMcpTool.FindMcpTool)        // 根据ID获取MCP工具管理
		group.GET("getMcpToolList", apiMcpTool.GetMcpToolList)  // 获取MCP工具管理列表
	}
	{
	    group := public.Group("mcpTool")
	    group.GET("getMcpToolDataSource", apiMcpTool.GetMcpToolDataSource)  // 获取MCP工具管理数据源
	    group.GET("getMcpToolPublic", apiMcpTool.GetMcpToolPublic)  // MCP工具管理开放接口
	}
}

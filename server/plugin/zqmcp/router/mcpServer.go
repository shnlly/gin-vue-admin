package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var McpServer = new(mcpServer)

type mcpServer struct {}

// Init 初始化 MCP服务器管理 路由信息
func (r *mcpServer) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
	    group := private.Group("mcpServer").Use(middleware.OperationRecord())
		group.POST("createMcpServer", apiMcpServer.CreateMcpServer)   // 新建MCP服务器管理
		group.DELETE("deleteMcpServer", apiMcpServer.DeleteMcpServer) // 删除MCP服务器管理
		group.DELETE("deleteMcpServerByIds", apiMcpServer.DeleteMcpServerByIds) // 批量删除MCP服务器管理
		group.PUT("updateMcpServer", apiMcpServer.UpdateMcpServer)    // 更新MCP服务器管理
	}
	{
	    group := private.Group("mcpServer")
		group.GET("findMcpServer", apiMcpServer.FindMcpServer)        // 根据ID获取MCP服务器管理
		group.GET("getMcpServerList", apiMcpServer.GetMcpServerList)  // 获取MCP服务器管理列表
		group.POST("start/:id", apiMcpServer.StartMcpServer)          // 启动MCP服务器
		group.POST("stop/:id", apiMcpServer.StopMcpServer)            // 停止MCP服务器
		group.GET("status/:id", apiMcpServer.GetMcpServerStatus)      // 获取MCP服务器状态
		group.GET("allocatePort", apiMcpServer.AllocatePort)          // 分配可用端口
	}
	{
	    group := public.Group("mcpServer")
	    group.GET("getMcpServerPublic", apiMcpServer.GetMcpServerPublic)  // MCP服务器管理开放接口
	}
}

package router

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/zqmcp/api"

var (
	Router       = new(router)
	apiMcpServer = api.Api.McpServer
	apiMcpTool   = api.Api.McpTool
)

type router struct {
	McpServer mcpServer
	McpTool   mcpTool
}

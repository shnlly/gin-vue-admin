package api

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/zqmcp/service"

var (
	Api              = new(api)
	serviceMcpServer = service.Service.McpServer
	serviceMcpTool   = service.Service.McpTool
)

type api struct {
	McpServer mcpServer
	McpTool   mcpTool
}

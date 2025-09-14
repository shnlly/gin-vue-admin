package service

var Service = new(service)

type service struct {
	McpServer mcpServer
	McpTool   mcpTool
}

package mcp

import (
	"context"
	"fmt"
	"sync"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/zqmcp/manager"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// ZqMcpTool 定义了ZQMCP工具必须实现的接口
type ZqMcpTool interface {
	// Handle 返回工具调用信息
	Handle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error)
	// New 返回工具注册信息
	New() mcp.Tool
	// ServerID 返回该工具属于的服务器ID
	ServerID() uint
}

// 工具注册表，按服务器ID分组
var toolRegisters = make(map[uint]map[string]ZqMcpTool)
var registerMutex sync.RWMutex

// RegisterTool 供工具在init时调用，将自己注册到指定服务器的工具注册表中
func RegisterTool(serverID uint, tool ZqMcpTool) {
	registerMutex.Lock()
	defer registerMutex.Unlock()

	mcpTool := tool.New()

	if toolRegisters[serverID] == nil {
		toolRegisters[serverID] = make(map[string]ZqMcpTool)
	}

	toolRegisters[serverID][mcpTool.Name] = tool

	// 保存到数据库
	mgr := manager.GetManager()
	err := mgr.RegisterTool(serverID, mcpTool.Name, mcpTool.Description, "")
	if err != nil {
		global.GVA_LOG.Error(fmt.Sprintf("注册工具到数据库失败: %s", err.Error()))
	}
}

// RegisterToolsToServer 将指定服务器的所有注册工具注册到MCP服务中
func RegisterToolsToServer(serverID uint, mcpServer *server.MCPServer) error {
	registerMutex.RLock()
	defer registerMutex.RUnlock()

	tools, exists := toolRegisters[serverID]
	if !exists {
		return fmt.Errorf("服务器 ID %d 没有注册的工具", serverID)
	}

	for _, tool := range tools {
		mcpServer.AddTool(tool.New(), tool.Handle)
	}

	return nil
}

// GetRegisteredTools 获取指定服务器的已注册工具列表
func GetRegisteredTools(serverID uint) map[string]ZqMcpTool {
	registerMutex.RLock()
	defer registerMutex.RUnlock()

	tools := toolRegisters[serverID]
	if tools == nil {
		return make(map[string]ZqMcpTool)
	}

	result := make(map[string]ZqMcpTool)
	for name, tool := range tools {
		result[name] = tool
	}
	return result
}

// GetAllRegisteredTools 获取所有服务器的已注册工具
func GetAllRegisteredTools() map[uint]map[string]ZqMcpTool {
	registerMutex.RLock()
	defer registerMutex.RUnlock()

	result := make(map[uint]map[string]ZqMcpTool)
	for serverID, tools := range toolRegisters {
		result[serverID] = make(map[string]ZqMcpTool)
		for name, tool := range tools {
			result[serverID][name] = tool
		}
	}
	return result
}

// ZqMcpRun 创建一个支持端口参数的MCP服务器
func ZqMcpRun(serverName, serverVersion string, port int) *server.SSEServer {
	s := server.NewMCPServer(
		serverName,
		serverVersion,
	)

	// 这里可以根据需要注册特定的工具到这个服务器
	// RegisterToolsToServer 会在服务器启动时被调用

	return server.NewSSEServer(s,
		server.WithSSEEndpoint("/sse"),
		server.WithMessageEndpoint("/message"),
		server.WithBaseURL(""))
}

// UnregisterTool 从服务器注册表中移除工具
func UnregisterTool(serverID uint, toolName string) {
	registerMutex.Lock()
	defer registerMutex.Unlock()

	if tools, exists := toolRegisters[serverID]; exists {
		delete(tools, toolName)
		if len(tools) == 0 {
			delete(toolRegisters, serverID)
		}
	}
}

// ListServerTools 列出指定服务器的工具
func ListServerTools(serverID uint) []string {
	registerMutex.RLock()
	defer registerMutex.RUnlock()

	tools, exists := toolRegisters[serverID]
	if !exists {
		return []string{}
	}

	var toolNames []string
	for name := range tools {
		toolNames = append(toolNames, name)
	}
	return toolNames
}
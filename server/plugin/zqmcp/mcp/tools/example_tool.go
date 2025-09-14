package tools

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/plugin/zqmcp/mcp"
	"github.com/mark3labs/mcp-go/mcp"
)

// ExampleTool 示例工具
type ExampleTool struct {
	serverID uint
}

// NewExampleTool 创建示例工具实例
func NewExampleTool(serverID uint) *ExampleTool {
	return &ExampleTool{serverID: serverID}
}

// Handle 处理工具调用
func (t *ExampleTool) Handle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	var args map[string]interface{}
	if request.Params.Arguments != nil {
		if err := json.Unmarshal(*request.Params.Arguments, &args); err != nil {
			return nil, fmt.Errorf("解析参数失败: %v", err)
		}
	}

	message, ok := args["message"].(string)
	if !ok {
		message = "Hello from ZQMCP!"
	}

	result := map[string]interface{}{
		"serverID": t.serverID,
		"message":  message,
		"status":   "success",
	}

	resultBytes, _ := json.Marshal(result)

	return &mcp.CallToolResult{
		Content: []interface{}{
			map[string]interface{}{
				"type": "text",
				"text": string(resultBytes),
			},
		},
		IsError: false,
	}, nil
}

// New 返回工具注册信息
func (t *ExampleTool) New() mcp.Tool {
	return mcp.Tool{
		Name:        fmt.Sprintf("example_tool_server_%d", t.serverID),
		Description: fmt.Sprintf("服务器 %d 的示例工具", t.serverID),
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"message": map[string]interface{}{
					"type":        "string",
					"description": "要处理的消息",
				},
			},
		},
	}
}

// ServerID 返回该工具属于的服务器ID
func (t *ExampleTool) ServerID() uint {
	return t.serverID
}

// 自动注册工具到服务器1（示例）
func init() {
	// 这里可以根据需要注册到不同的服务器
	// mcp.RegisterTool(1, NewExampleTool(1))
}
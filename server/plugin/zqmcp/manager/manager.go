package manager

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"sync"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/zqmcp/model"
	"github.com/mark3labs/mcp-go/server"
	"gorm.io/gorm"
)

// McpServerManager MCP服务器管理器
type McpServerManager struct {
	servers    map[uint]*ServerInstance // 服务器实例映射
	mu         sync.RWMutex             // 读写锁
	portRange  PortRange                // 端口范围
	db         *gorm.DB                 // 数据库连接
}

// ServerInstance MCP服务器实例
type ServerInstance struct {
	ID         uint                  `json:"id"`
	Config     *model.McpServer      `json:"config"`
	McpServer  *server.MCPServer     `json:"-"`
	SSEServer  *server.SSEServer     `json:"-"`
	Status     string                `json:"status"`
	StartTime  time.Time             `json:"startTime"`
	Tools      map[string]*ToolInfo  `json:"tools"`
	Connections int                  `json:"connections"`
}

// ToolInfo 工具信息
type ToolInfo struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Schema      string    `json:"schema"`
	CallCount   int64     `json:"callCount"`
	LastCall    time.Time `json:"lastCall"`
}

// PortRange 端口范围配置
type PortRange struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

var (
	defaultManager *McpServerManager
	once           sync.Once
)

// GetManager 获取服务器管理器单例
func GetManager() *McpServerManager {
	once.Do(func() {
		defaultManager = &McpServerManager{
			servers: make(map[uint]*ServerInstance),
			portRange: PortRange{
				Min: 30000,
				Max: 40000,
			},
			db: global.GVA_DB,
		}
		// 启动时恢复所有服务器状态
		defaultManager.restoreServers()
	})
	return defaultManager
}

// AllocatePort 自动分配端口
func (m *McpServerManager) AllocatePort() (int, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	// 获取已使用的端口
	usedPorts := make(map[int]bool)
	for _, instance := range m.servers {
		if instance.Config.ServerPort != nil {
			usedPorts[int(*instance.Config.ServerPort)] = true
		}
	}

	// 从数据库获取已使用的端口
	var servers []model.McpServer
	if err := m.db.Find(&servers).Error; err != nil {
		return 0, err
	}
	for _, srv := range servers {
		if srv.ServerPort != nil {
			usedPorts[int(*srv.ServerPort)] = true
		}
	}

	// 随机分配端口
	rand.Seed(time.Now().UnixNano())
	maxAttempts := 1000
	for i := 0; i < maxAttempts; i++ {
		port := m.portRange.Min + rand.Intn(m.portRange.Max-m.portRange.Min)
		if !usedPorts[port] && m.isPortAvailable(port) {
			return port, nil
		}
	}

	return 0, fmt.Errorf("无法分配可用端口")
}

// isPortAvailable 检查端口是否可用
func (m *McpServerManager) isPortAvailable(port int) bool {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return false
	}
	ln.Close()
	return true
}

// CreateServer 创建MCP服务器
func (m *McpServerManager) CreateServer(config *model.McpServer) (*ServerInstance, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	// 如果没有指定端口，自动分配
	if config.ServerPort == nil || *config.ServerPort == 0 {
		port, err := m.AllocatePort()
		if err != nil {
			return nil, err
		}
		config.ServerPort = new(int64)
		*config.ServerPort = int64(port)
		config.AutoAllocated = new(bool)
		*config.AutoAllocated = true
	}

	// 检查端口是否已被使用
	for _, instance := range m.servers {
		if instance.Config.ServerPort != nil && config.ServerPort != nil &&
			*instance.Config.ServerPort == *config.ServerPort {
			return nil, fmt.Errorf("端口 %d 已被使用", *config.ServerPort)
		}
	}

	// 创建MCP服务器实例
	mcpServer := server.NewMCPServer(
		*config.ServerName,
		*config.ServerVersion,
	)

	// 创建服务器实例
	instance := &ServerInstance{
		ID:          config.ID,
		Config:      config,
		McpServer:   mcpServer,
		Status:      "stopped",
		Tools:       make(map[string]*ToolInfo),
		Connections: 0,
	}

	m.servers[config.ID] = instance
	return instance, nil
}

// StartServer 启动MCP服务器
func (m *McpServerManager) StartServer(serverID uint) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	instance, exists := m.servers[serverID]
	if !exists {
		return fmt.Errorf("服务器 ID %d 不存在", serverID)
	}

	if instance.Status == "running" {
		return fmt.Errorf("服务器已经在运行")
	}

	// 更新状态为启动中
	instance.Status = "starting"
	m.updateServerStatus(serverID, "starting")

	// 创建SSE服务器
	sseServer := server.NewSSEServer(instance.McpServer,
		server.WithSSEEndpoint("/sse"),
		server.WithMessageEndpoint("/message"),
		server.WithBaseURL(""))

	instance.SSEServer = sseServer
	instance.StartTime = time.Now()
	instance.Status = "running"

	// 更新数据库状态
	m.updateServerStatus(serverID, "running")
	m.updateServerStartTime(serverID, instance.StartTime)

	return nil
}

// StopServer 停止MCP服务器
func (m *McpServerManager) StopServer(serverID uint) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	instance, exists := m.servers[serverID]
	if !exists {
		return fmt.Errorf("服务器 ID %d 不存在", serverID)
	}

	if instance.Status == "stopped" {
		return fmt.Errorf("服务器已经停止")
	}

	// 停止服务器
	instance.Status = "stopped"
	instance.SSEServer = nil
	instance.Connections = 0

	// 更新数据库状态
	m.updateServerStatus(serverID, "stopped")
	m.updateConnectionCount(serverID, 0)

	return nil
}

// GetServer 获取服务器实例
func (m *McpServerManager) GetServer(serverID uint) (*ServerInstance, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	instance, exists := m.servers[serverID]
	if !exists {
		return nil, fmt.Errorf("服务器 ID %d 不存在", serverID)
	}
	return instance, nil
}

// ListServers 列出所有服务器
func (m *McpServerManager) ListServers() map[uint]*ServerInstance {
	m.mu.RLock()
	defer m.mu.RUnlock()

	result := make(map[uint]*ServerInstance)
	for id, instance := range m.servers {
		result[id] = instance
	}
	return result
}

// RegisterTool 为指定服务器注册工具
func (m *McpServerManager) RegisterTool(serverID uint, toolName, description, schema string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	instance, exists := m.servers[serverID]
	if !exists {
		return fmt.Errorf("服务器 ID %d 不存在", serverID)
	}

	// 添加工具信息
	instance.Tools[toolName] = &ToolInfo{
		Name:        toolName,
		Description: description,
		Schema:      schema,
		CallCount:   0,
		LastCall:    time.Time{},
	}

	// 保存到数据库
	toolModel := &model.McpTool{
		ServerId:        new(int64),
		ToolName:        &toolName,
		ToolDescription: &description,
		ToolCategory:    new(string),
		IsEnabled:       new(bool),
		CallCount:       new(int64),
	}
	*toolModel.ServerId = int64(serverID)
	*toolModel.ToolCategory = "custom"
	*toolModel.IsEnabled = true
	*toolModel.CallCount = 0

	if err := m.db.Create(toolModel).Error; err != nil {
		return err
	}

	return nil
}

// restoreServers 恢复服务器状态
func (m *McpServerManager) restoreServers() {
	var servers []model.McpServer
	if err := m.db.Find(&servers).Error; err != nil {
		global.GVA_LOG.Error("恢复MCP服务器状态失败: " + err.Error())
		return
	}

	for _, srv := range servers {
		instance, err := m.CreateServer(&srv)
		if err != nil {
			global.GVA_LOG.Error(fmt.Sprintf("恢复服务器 %d 失败: %s", srv.ID, err.Error()))
			continue
		}

		// 如果之前是运行状态，尝试重新启动
		if *srv.ServerStatus == "running" {
			if err := m.StartServer(srv.ID); err != nil {
				global.GVA_LOG.Error(fmt.Sprintf("重启服务器 %d 失败: %s", srv.ID, err.Error()))
				// 更新状态为错误
				m.updateServerStatus(srv.ID, "error")
			}
		}
	}
}

// updateServerStatus 更新服务器状态
func (m *McpServerManager) updateServerStatus(serverID uint, status string) {
	if err := m.db.Model(&model.McpServer{}).Where("id = ?", serverID).
		Update("server_status", status).Error; err != nil {
		global.GVA_LOG.Error(fmt.Sprintf("更新服务器 %d 状态失败: %s", serverID, err.Error()))
	}
}

// updateServerStartTime 更新服务器启动时间
func (m *McpServerManager) updateServerStartTime(serverID uint, startTime time.Time) {
	if err := m.db.Model(&model.McpServer{}).Where("id = ?", serverID).
		Update("last_start_time", startTime).Error; err != nil {
		global.GVA_LOG.Error(fmt.Sprintf("更新服务器 %d 启动时间失败: %s", serverID, err.Error()))
	}
}

// updateConnectionCount 更新连接数
func (m *McpServerManager) updateConnectionCount(serverID uint, count int) {
	if err := m.db.Model(&model.McpServer{}).Where("id = ?", serverID).
		Update("connection_count", count).Error; err != nil {
		global.GVA_LOG.Error(fmt.Sprintf("更新服务器 %d 连接数失败: %s", serverID, err.Error()))
	}
}
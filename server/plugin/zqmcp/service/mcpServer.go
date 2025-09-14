
package service

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/zqmcp/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/zqmcp/model/request"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/zqmcp/manager"
    "gorm.io/gorm"
)

var McpServer = new(mcpServer)

type mcpServer struct {}
// CreateMcpServer 创建MCP服务器管理记录
// Author [yourname](https://github.com/yourname)
func (s *mcpServer) CreateMcpServer(ctx context.Context, mcpServer *model.McpServer) (err error) {
	// 如果端口为0，自动分配端口
	if mcpServer.ServerPort == nil || *mcpServer.ServerPort == 0 {
		mgr := manager.GetManager()
		port, err := mgr.AllocatePort()
		if err != nil {
			return err
		}
		mcpServer.ServerPort = new(int64)
		*mcpServer.ServerPort = int64(port)
		mcpServer.AutoAllocated = new(bool)
		*mcpServer.AutoAllocated = true
	}

	// 创建数据库记录
	err = global.GVA_DB.Create(mcpServer).Error
	if err != nil {
		return err
	}

	// 创建管理器中的服务器实例
	mgr := manager.GetManager()
	_, err = mgr.CreateServer(mcpServer)

	return err
}

// DeleteMcpServer 删除MCP服务器管理记录
// Author [yourname](https://github.com/yourname)
func (s *mcpServer) DeleteMcpServer(ctx context.Context, ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&model.McpServer{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&model.McpServer{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteMcpServerByIds 批量删除MCP服务器管理记录
// Author [yourname](https://github.com/yourname)
func (s *mcpServer) DeleteMcpServerByIds(ctx context.Context, IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&model.McpServer{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&model.McpServer{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateMcpServer 更新MCP服务器管理记录
// Author [yourname](https://github.com/yourname)
func (s *mcpServer) UpdateMcpServer(ctx context.Context, mcpServer model.McpServer) (err error) {
	err = global.GVA_DB.Model(&model.McpServer{}).Where("id = ?",mcpServer.ID).Updates(&mcpServer).Error
	return err
}

// GetMcpServer 根据ID获取MCP服务器管理记录
// Author [yourname](https://github.com/yourname)
func (s *mcpServer) GetMcpServer(ctx context.Context, ID string) (mcpServer model.McpServer, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&mcpServer).Error
	return
}
// GetMcpServerInfoList 分页获取MCP服务器管理记录
// Author [yourname](https://github.com/yourname)
func (s *mcpServer) GetMcpServerInfoList(ctx context.Context, info request.McpServerSearch) (list []model.McpServer, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.McpServer{})
    var mcpServers []model.McpServer
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
  
    if info.ServerName != nil && *info.ServerName != "" {
        db = db.Where("server_name LIKE ?", "%"+ *info.ServerName+"%")
    }
    if info.ServerPort != nil {
        db = db.Where("server_port = ?", *info.ServerPort)
    }
    if info.ServerVersion != nil && *info.ServerVersion != "" {
        db = db.Where("server_version = ?", *info.ServerVersion)
    }
    if info.ServerStatus != nil && *info.ServerStatus != "" {
        db = db.Where("server_status = ?", *info.ServerStatus)
    }
    if info.AutoAllocated != nil {
        db = db.Where("auto_allocated = ?", *info.AutoAllocated)
    }
    if info.Description != nil && *info.Description != "" {
        db = db.Where("description LIKE ?", "%"+ *info.Description+"%")
    }
			if len(info.LastStartTimeRange) == 2 {
				db = db.Where("last_start_time BETWEEN ? AND ? ", info.LastStartTimeRange[0], info.LastStartTimeRange[1])
			}
    if info.ConnectionCount != nil {
        db = db.Where("connection_count >= ?", *info.ConnectionCount)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
        orderMap["id"] = true
        orderMap["created_at"] = true
        orderMap["last_start_time"] = true
        orderMap["connection_count"] = true
       if orderMap[info.Sort] {
          OrderStr = info.Sort
          if info.Order == "descending" {
             OrderStr = OrderStr + " desc"
          }
          db = db.Order(OrderStr)
       }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	err = db.Find(&mcpServers).Error
	return  mcpServers, total, err
}

func (s *mcpServer)GetMcpServerPublic(ctx context.Context) {

}

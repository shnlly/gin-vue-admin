
package service

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/zqmcp/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/zqmcp/model/request"
    "gorm.io/gorm"
)

var McpTool = new(mcpTool)

type mcpTool struct {}
// CreateMcpTool 创建MCP工具管理记录
// Author [yourname](https://github.com/yourname)
func (s *mcpTool) CreateMcpTool(ctx context.Context, mcpTool *model.McpTool) (err error) {
	err = global.GVA_DB.Create(mcpTool).Error
	return err
}

// DeleteMcpTool 删除MCP工具管理记录
// Author [yourname](https://github.com/yourname)
func (s *mcpTool) DeleteMcpTool(ctx context.Context, ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&model.McpTool{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&model.McpTool{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteMcpToolByIds 批量删除MCP工具管理记录
// Author [yourname](https://github.com/yourname)
func (s *mcpTool) DeleteMcpToolByIds(ctx context.Context, IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&model.McpTool{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&model.McpTool{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateMcpTool 更新MCP工具管理记录
// Author [yourname](https://github.com/yourname)
func (s *mcpTool) UpdateMcpTool(ctx context.Context, mcpTool model.McpTool) (err error) {
	err = global.GVA_DB.Model(&model.McpTool{}).Where("id = ?",mcpTool.ID).Updates(&mcpTool).Error
	return err
}

// GetMcpTool 根据ID获取MCP工具管理记录
// Author [yourname](https://github.com/yourname)
func (s *mcpTool) GetMcpTool(ctx context.Context, ID string) (mcpTool model.McpTool, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&mcpTool).Error
	return
}
// GetMcpToolInfoList 分页获取MCP工具管理记录
// Author [yourname](https://github.com/yourname)
func (s *mcpTool) GetMcpToolInfoList(ctx context.Context, info request.McpToolSearch) (list []model.McpTool, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.McpTool{})
    var mcpTools []model.McpTool
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
  
    if info.ServerId != nil {
        db = db.Where("server_id = ?", *info.ServerId)
    }
    if info.ToolName != nil && *info.ToolName != "" {
        db = db.Where("tool_name LIKE ?", "%"+ *info.ToolName+"%")
    }
    if info.ToolDescription != nil && *info.ToolDescription != "" {
        db = db.Where("tool_description LIKE ?", "%"+ *info.ToolDescription+"%")
    }
    if info.ToolCategory != nil && *info.ToolCategory != "" {
        db = db.Where("tool_category = ?", *info.ToolCategory)
    }
    if info.IsEnabled != nil {
        db = db.Where("is_enabled = ?", *info.IsEnabled)
    }
    if info.CallCount != nil {
        db = db.Where("call_count >= ?", *info.CallCount)
    }
			if len(info.LastCallTimeRange) == 2 {
				db = db.Where("last_call_time BETWEEN ? AND ? ", info.LastCallTimeRange[0], info.LastCallTimeRange[1])
			}
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
        orderMap["id"] = true
        orderMap["created_at"] = true
        orderMap["call_count"] = true
        orderMap["last_call_time"] = true
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
	err = db.Find(&mcpTools).Error
	return  mcpTools, total, err
}
func (s *mcpTool)GetMcpToolDataSource(ctx context.Context) (res map[string][]map[string]any, err error) {
	res = make(map[string][]map[string]any)
	
	   serverId := make([]map[string]any, 0)
	   global.GVA_DB.Table("zqmcp_servers").Where("deleted_at IS NULL").Select("server_name as label,id as value").Scan(&serverId)
	   res["serverId"] = serverId
	return
}

func (s *mcpTool)GetMcpToolPublic(ctx context.Context) {

}

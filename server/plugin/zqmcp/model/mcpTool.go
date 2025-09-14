
package model
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
	"gorm.io/datatypes"
)

// McpTool MCP工具管理 结构体
type McpTool struct {
    global.GVA_MODEL
  ServerId  *int64 `json:"serverId" form:"serverId" gorm:"INDEX;comment:关联的MCP服务器ID;column:server_id;" binding:"required"`  //所属MCP服务器ID
  ToolName  *string `json:"toolName" form:"toolName" gorm:"INDEX;comment:MCP工具名称;column:tool_name;size:100;" binding:"required"`  //工具名称
  ToolDescription  *string `json:"toolDescription" form:"toolDescription" gorm:"comment:MCP工具功能描述;column:tool_description;size:500;"`  //工具描述
  ToolSchema  datatypes.JSON `json:"toolSchema" form:"toolSchema" gorm:"comment:工具参数JSON Schema;column:tool_schema;" swaggertype:"object"`  //工具参数模式
  ToolCategory  *string `json:"toolCategory" form:"toolCategory" gorm:"INDEX;default:custom;comment:工具分类：system-系统工具,database-数据库工具,file-文件工具,network-网络工具,custom-自定义工具;column:tool_category;size:50;" binding:"required"`  //工具分类
  IsEnabled  *bool `json:"isEnabled" form:"isEnabled" gorm:"default:true;comment:工具是否启用状态;column:is_enabled;"`  //是否启用
  CallCount  *int64 `json:"callCount" form:"callCount" gorm:"default:0;comment:工具总调用次数统计;column:call_count;"`  //调用次数
  LastCallTime  *time.Time `json:"lastCallTime" form:"lastCallTime" gorm:"comment:工具最后一次调用时间;column:last_call_time;"`  //最后调用时间
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName MCP工具管理 McpTool自定义表名 zqmcp_tools
func (McpTool) TableName() string {
    return "zqmcp_tools"
}








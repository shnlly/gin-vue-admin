
package model
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
	"gorm.io/datatypes"
)

// McpServer MCP服务器管理 结构体
type McpServer struct {
    global.GVA_MODEL
  ServerName  *string `json:"serverName" form:"serverName" gorm:"INDEX;comment:MCP服务器显示名称;column:server_name;size:100;" binding:"required"`  //MCP服务器名称
  ServerPort  *int64 `json:"serverPort" form:"serverPort" gorm:"UNIQUE;default:0;comment:MCP服务器监听端口;column:server_port;" binding:"required"`  //服务器端口号
  ServerVersion  *string `json:"serverVersion" form:"serverVersion" gorm:"default:v1.0.0;comment:MCP服务器版本号;column:server_version;size:50;"`  //服务器版本
  ServerStatus  *string `json:"serverStatus" form:"serverStatus" gorm:"INDEX;default:stopped;comment:服务器运行状态：stopped-已停止,starting-启动中,running-运行中,error-异常;column:server_status;size:20;" binding:"required"`  //服务器状态
  AutoAllocated  *bool `json:"autoAllocated" form:"autoAllocated" gorm:"default:false;comment:端口是否由系统自动分配;column:auto_allocated;"`  //是否自动分配端口
  ServerConfig  datatypes.JSON `json:"serverConfig" form:"serverConfig" gorm:"comment:MCP服务器配置信息JSON;column:server_config;" swaggertype:"object"`  //服务器配置
  Description  *string `json:"description" form:"description" gorm:"comment:MCP服务器描述信息;column:description;size:500;"`  //服务器描述
  LastStartTime  *time.Time `json:"lastStartTime" form:"lastStartTime" gorm:"comment:服务器最后启动时间;column:last_start_time;"`  //最后启动时间
  ConnectionCount  *int64 `json:"connectionCount" form:"connectionCount" gorm:"default:0;comment:当前活跃连接数量;column:connection_count;"`  //当前连接数
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName MCP服务器管理 McpServer自定义表名 zqmcp_servers
func (McpServer) TableName() string {
    return "zqmcp_servers"
}








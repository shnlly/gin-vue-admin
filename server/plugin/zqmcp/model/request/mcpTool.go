
package request
import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)
type McpToolSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
       ServerId  *int `json:"serverId" form:"serverId"` 
       ToolName  *string `json:"toolName" form:"toolName"` 
       ToolDescription  *string `json:"toolDescription" form:"toolDescription"` 
       ToolCategory  *string `json:"toolCategory" form:"toolCategory"` 
       IsEnabled  *bool `json:"isEnabled" form:"isEnabled"` 
       CallCount  *int `json:"callCount" form:"callCount"` 
       LastCallTimeRange  []time.Time  `json:"lastCallTimeRange" form:"lastCallTimeRange[]"`
    request.PageInfo
    Sort  string `json:"sort" form:"sort"`
    Order string `json:"order" form:"order"`
}

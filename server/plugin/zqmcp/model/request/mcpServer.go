
package request
import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)
type McpServerSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
       ServerName  *string `json:"serverName" form:"serverName"` 
       ServerPort  *int `json:"serverPort" form:"serverPort"` 
       ServerVersion  *string `json:"serverVersion" form:"serverVersion"` 
       ServerStatus  *string `json:"serverStatus" form:"serverStatus"` 
       AutoAllocated  *bool `json:"autoAllocated" form:"autoAllocated"` 
       Description  *string `json:"description" form:"description"` 
       LastStartTimeRange  []time.Time  `json:"lastStartTimeRange" form:"lastStartTimeRange[]"`
       ConnectionCount  *int `json:"connectionCount" form:"connectionCount"` 
    request.PageInfo
    Sort  string `json:"sort" form:"sort"`
    Order string `json:"order" form:"order"`
}

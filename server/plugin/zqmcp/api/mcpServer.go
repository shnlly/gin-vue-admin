package api

import (

	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/zqmcp/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/zqmcp/model/request"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/zqmcp/manager"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
    "strconv"
)

var McpServer = new(mcpServer)

type mcpServer struct {}

// CreateMcpServer 创建MCP服务器管理
// @Tags McpServer
// @Summary 创建MCP服务器管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.McpServer true "创建MCP服务器管理"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /mcpServer/createMcpServer [post]
func (a *mcpServer) CreateMcpServer(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.McpServer
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    info.CreatedBy = utils.GetUserID(c)
	err = serviceMcpServer.CreateMcpServer(ctx,&info)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteMcpServer 删除MCP服务器管理
// @Tags McpServer
// @Summary 删除MCP服务器管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.McpServer true "删除MCP服务器管理"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /mcpServer/deleteMcpServer [delete]
func (a *mcpServer) DeleteMcpServer(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
    userID := utils.GetUserID(c)
	err := serviceMcpServer.DeleteMcpServer(ctx,ID,userID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("删除成功", c)
}

// DeleteMcpServerByIds 批量删除MCP服务器管理
// @Tags McpServer
// @Summary 批量删除MCP服务器管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /mcpServer/deleteMcpServerByIds [delete]
func (a *mcpServer) DeleteMcpServerByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	err := serviceMcpServer.DeleteMcpServerByIds(ctx,IDs,userID)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("批量删除成功", c)
}

// UpdateMcpServer 更新MCP服务器管理
// @Tags McpServer
// @Summary 更新MCP服务器管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.McpServer true "更新MCP服务器管理"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /mcpServer/updateMcpServer [put]
func (a *mcpServer) UpdateMcpServer(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.McpServer
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    info.UpdatedBy = utils.GetUserID(c)
	err = serviceMcpServer.UpdateMcpServer(ctx,info)
    if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("更新成功", c)
}

// FindMcpServer 用id查询MCP服务器管理
// @Tags McpServer
// @Summary 用id查询MCP服务器管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询MCP服务器管理"
// @Success 200 {object} response.Response{data=model.McpServer,msg=string} "查询成功"
// @Router /mcpServer/findMcpServer [get]
func (a *mcpServer) FindMcpServer(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	remcpServer, err := serviceMcpServer.GetMcpServer(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
    response.OkWithData(remcpServer, c)
}
// GetMcpServerList 分页获取MCP服务器管理列表
// @Tags McpServer
// @Summary 分页获取MCP服务器管理列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.McpServerSearch true "分页获取MCP服务器管理列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /mcpServer/getMcpServerList [get]
func (a *mcpServer) GetMcpServerList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo request.McpServerSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := serviceMcpServer.GetMcpServerInfoList(ctx,pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}
// GetMcpServerPublic 不需要鉴权的MCP服务器管理接口
// @Tags McpServer
// @Summary 不需要鉴权的MCP服务器管理接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /mcpServer/getMcpServerPublic [get]
func (a *mcpServer) GetMcpServerPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    serviceMcpServer.GetMcpServerPublic(ctx)
    response.OkWithDetailed(gin.H{"info": "不需要鉴权的MCP服务器管理接口信息"}, "获取成功", c)
}

// StartMcpServer 启动MCP服务器
// @Tags McpServer
// @Summary 启动MCP服务器
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id path int true "服务器ID"
// @Success 200 {object} response.Response{msg=string} "启动成功"
// @Router /mcpServer/start/{id} [post]
func (a *mcpServer) StartMcpServer(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.ParseUint(idStr, 10, 32)
    if err != nil {
        response.FailWithMessage("无效的服务器ID", c)
        return
    }

    mgr := manager.GetManager()
    err = mgr.StartServer(uint(id))
    if err != nil {
        global.GVA_LOG.Error("启动MCP服务器失败!", zap.Error(err))
        response.FailWithMessage("启动失败: "+err.Error(), c)
        return
    }

    response.OkWithMessage("启动成功", c)
}

// StopMcpServer 停止MCP服务器
// @Tags McpServer
// @Summary 停止MCP服务器
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id path int true "服务器ID"
// @Success 200 {object} response.Response{msg=string} "停止成功"
// @Router /mcpServer/stop/{id} [post]
func (a *mcpServer) StopMcpServer(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.ParseUint(idStr, 10, 32)
    if err != nil {
        response.FailWithMessage("无效的服务器ID", c)
        return
    }

    mgr := manager.GetManager()
    err = mgr.StopServer(uint(id))
    if err != nil {
        global.GVA_LOG.Error("停止MCP服务器失败!", zap.Error(err))
        response.FailWithMessage("停止失败: "+err.Error(), c)
        return
    }

    response.OkWithMessage("停止成功", c)
}

// GetMcpServerStatus 获取MCP服务器状态
// @Tags McpServer
// @Summary 获取MCP服务器状态
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id path int true "服务器ID"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /mcpServer/status/{id} [get]
func (a *mcpServer) GetMcpServerStatus(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.ParseUint(idStr, 10, 32)
    if err != nil {
        response.FailWithMessage("无效的服务器ID", c)
        return
    }

    mgr := manager.GetManager()
    instance, err := mgr.GetServer(uint(id))
    if err != nil {
        response.FailWithMessage("获取服务器状态失败: "+err.Error(), c)
        return
    }

    response.OkWithData(gin.H{
        "id":          instance.ID,
        "status":      instance.Status,
        "startTime":   instance.StartTime,
        "tools":       instance.Tools,
        "connections": instance.Connections,
    }, c)
}

// AllocatePort 分配可用端口
// @Tags McpServer
// @Summary 分配可用端口
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "分配成功"
// @Router /mcpServer/allocatePort [get]
func (a *mcpServer) AllocatePort(c *gin.Context) {
    mgr := manager.GetManager()
    port, err := mgr.AllocatePort()
    if err != nil {
        response.FailWithMessage("分配端口失败: "+err.Error(), c)
        return
    }

    response.OkWithData(gin.H{"port": port}, c)
}

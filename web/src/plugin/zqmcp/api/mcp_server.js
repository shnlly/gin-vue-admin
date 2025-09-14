import service from '@/utils/request'
// @Tags McpServer
// @Summary 创建MCP服务器管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.McpServer true "创建MCP服务器管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /mcpServer/createMcpServer [post]
export const createMcpServer = (data) => {
  return service({
    url: '/mcpServer/createMcpServer',
    method: 'post',
    data
  })
}

// @Tags McpServer
// @Summary 删除MCP服务器管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.McpServer true "删除MCP服务器管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mcpServer/deleteMcpServer [delete]
export const deleteMcpServer = (params) => {
  return service({
    url: '/mcpServer/deleteMcpServer',
    method: 'delete',
    params
  })
}

// @Tags McpServer
// @Summary 批量删除MCP服务器管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MCP服务器管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mcpServer/deleteMcpServer [delete]
export const deleteMcpServerByIds = (params) => {
  return service({
    url: '/mcpServer/deleteMcpServerByIds',
    method: 'delete',
    params
  })
}

// @Tags McpServer
// @Summary 更新MCP服务器管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.McpServer true "更新MCP服务器管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /mcpServer/updateMcpServer [put]
export const updateMcpServer = (data) => {
  return service({
    url: '/mcpServer/updateMcpServer',
    method: 'put',
    data
  })
}

// @Tags McpServer
// @Summary 用id查询MCP服务器管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.McpServer true "用id查询MCP服务器管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /mcpServer/findMcpServer [get]
export const findMcpServer = (params) => {
  return service({
    url: '/mcpServer/findMcpServer',
    method: 'get',
    params
  })
}

// @Tags McpServer
// @Summary 分页获取MCP服务器管理列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取MCP服务器管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mcpServer/getMcpServerList [get]
export const getMcpServerList = (params) => {
  return service({
    url: '/mcpServer/getMcpServerList',
    method: 'get',
    params
  })
}
// @Tags McpServer
// @Summary 不需要鉴权的MCP服务器管理接口
// @Accept application/json
// @Produce application/json
// @Param data query request.McpServerSearch true "分页获取MCP服务器管理列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /mcpServer/getMcpServerPublic [get]
export const getMcpServerPublic = () => {
  return service({
    url: '/mcpServer/getMcpServerPublic',
    method: 'get',
  })
}

// @Tags McpServer
// @Summary 启动MCP服务器
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id path int true "服务器ID"
// @Success 200 {object} response.Response{msg=string} "启动成功"
// @Router /mcpServer/start/{id} [post]
export const startMcpServer = (id) => {
  return service({
    url: `/mcpServer/start/${id}`,
    method: 'post'
  })
}

// @Tags McpServer
// @Summary 停止MCP服务器
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id path int true "服务器ID"
// @Success 200 {object} response.Response{msg=string} "停止成功"
// @Router /mcpServer/stop/{id} [post]
export const stopMcpServer = (id) => {
  return service({
    url: `/mcpServer/stop/${id}`,
    method: 'post'
  })
}

// @Tags McpServer
// @Summary 获取MCP服务器状态
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id path int true "服务器ID"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /mcpServer/status/{id} [get]
export const getMcpServerStatus = (id) => {
  return service({
    url: `/mcpServer/status/${id}`,
    method: 'get'
  })
}

// @Tags McpServer
// @Summary 分配可用端口
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "分配成功"
// @Router /mcpServer/allocatePort [get]
export const allocatePort = () => {
  return service({
    url: '/mcpServer/allocatePort',
    method: 'get'
  })
}

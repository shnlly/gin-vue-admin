import service from '@/utils/request'
// @Tags McpTool
// @Summary 创建MCP工具管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.McpTool true "创建MCP工具管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /mcpTool/createMcpTool [post]
export const createMcpTool = (data) => {
  return service({
    url: '/mcpTool/createMcpTool',
    method: 'post',
    data
  })
}

// @Tags McpTool
// @Summary 删除MCP工具管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.McpTool true "删除MCP工具管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mcpTool/deleteMcpTool [delete]
export const deleteMcpTool = (params) => {
  return service({
    url: '/mcpTool/deleteMcpTool',
    method: 'delete',
    params
  })
}

// @Tags McpTool
// @Summary 批量删除MCP工具管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MCP工具管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mcpTool/deleteMcpTool [delete]
export const deleteMcpToolByIds = (params) => {
  return service({
    url: '/mcpTool/deleteMcpToolByIds',
    method: 'delete',
    params
  })
}

// @Tags McpTool
// @Summary 更新MCP工具管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.McpTool true "更新MCP工具管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /mcpTool/updateMcpTool [put]
export const updateMcpTool = (data) => {
  return service({
    url: '/mcpTool/updateMcpTool',
    method: 'put',
    data
  })
}

// @Tags McpTool
// @Summary 用id查询MCP工具管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.McpTool true "用id查询MCP工具管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /mcpTool/findMcpTool [get]
export const findMcpTool = (params) => {
  return service({
    url: '/mcpTool/findMcpTool',
    method: 'get',
    params
  })
}

// @Tags McpTool
// @Summary 分页获取MCP工具管理列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取MCP工具管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mcpTool/getMcpToolList [get]
export const getMcpToolList = (params) => {
  return service({
    url: '/mcpTool/getMcpToolList',
    method: 'get',
    params
  })
}
// @Tags McpTool
// @Summary 获取数据源
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /mcpTool/findMcpToolDataSource [get]
export const getMcpToolDataSource = () => {
  return service({
    url: '/mcpTool/getMcpToolDataSource',
    method: 'get',
  })
}
// @Tags McpTool
// @Summary 不需要鉴权的MCP工具管理接口
// @Accept application/json
// @Produce application/json
// @Param data query request.McpToolSearch true "分页获取MCP工具管理列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /mcpTool/getMcpToolPublic [get]
export const getMcpToolPublic = () => {
  return service({
    url: '/mcpTool/getMcpToolPublic',
    method: 'get',
  })
}

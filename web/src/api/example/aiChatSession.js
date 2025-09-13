import service from '@/utils/request'

// @Tags AiChatSession
// @Summary 创建AI对话会话
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AiChatSession true "创建AI对话会话"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /aiChatSession/createAiChatSession [post]
export const createAiChatSession = (data) => {
  return service({
    url: '/aiChatSession/createAiChatSession',
    method: 'post',
    data
  })
}

// @Tags AiChatSession
// @Summary 删除AI对话会话
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AiChatSession true "删除AI对话会话"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /aiChatSession/deleteAiChatSession [delete]
export const deleteAiChatSession = (params) => {
  return service({
    url: '/aiChatSession/deleteAiChatSession',
    method: 'delete',
    params
  })
}

// @Tags AiChatSession
// @Summary 批量删除AI对话会话
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /aiChatSession/deleteAiChatSessionByIds [delete]
export const deleteAiChatSessionByIds = (params) => {
  return service({
    url: '/aiChatSession/deleteAiChatSessionByIds',
    method: 'delete',
    params
  })
}

// @Tags AiChatSession
// @Summary 更新AI对话会话
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AiChatSession true "更新AI对话会话"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /aiChatSession/updateAiChatSession [put]
export const updateAiChatSession = (data) => {
  return service({
    url: '/aiChatSession/updateAiChatSession',
    method: 'put',
    data
  })
}

// @Tags AiChatSession
// @Summary 用id查询AI对话会话
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.AiChatSession true "用id查询AI对话会话"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /aiChatSession/findAiChatSession [get]
export const findAiChatSession = (params) => {
  return service({
    url: '/aiChatSession/findAiChatSession',
    method: 'get',
    params
  })
}

// @Tags AiChatSession
// @Summary 分页获取AI对话会话列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取AI对话会话列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /aiChatSession/getAiChatSessionList [get]
export const getAiChatSessionList = (params) => {
  return service({
    url: '/aiChatSession/getAiChatSessionList',
    method: 'get',
    params
  })
}

// @Tags AiChatSession
// @Summary 获取用户的对话会话列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /aiChatSession/getUserSessions [get]
export const getUserChatSessions = () => {
  return service({
    url: '/aiChatSession/getUserSessions',
    method: 'get'
  })
}
import service from '@/utils/request'

// @Tags AiModel
// @Summary 创建AI模型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AiModel true "创建AI模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /aiModel/createAiModel [post]
export const createAiModel = (data) => {
  return service({
    url: '/aiModel/createAiModel',
    method: 'post',
    data
  })
}

// @Tags AiModel
// @Summary 删除AI模型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AiModel true "删除AI模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /aiModel/deleteAiModel [delete]
export const deleteAiModel = (params) => {
  return service({
    url: '/aiModel/deleteAiModel',
    method: 'delete',
    params
  })
}

// @Tags AiModel
// @Summary 批量删除AI模型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /aiModel/deleteAiModelByIds [delete]
export const deleteAiModelByIds = (params) => {
  return service({
    url: '/aiModel/deleteAiModelByIds',
    method: 'delete',
    params
  })
}

// @Tags AiModel
// @Summary 更新AI模型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AiModel true "更新AI模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /aiModel/updateAiModel [put]
export const updateAiModel = (data) => {
  return service({
    url: '/aiModel/updateAiModel',
    method: 'put',
    data
  })
}

// @Tags AiModel
// @Summary 用id查询AI模型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.AiModel true "用id查询AI模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /aiModel/findAiModel [get]
export const findAiModel = (params) => {
  return service({
    url: '/aiModel/findAiModel',
    method: 'get',
    params
  })
}

// @Tags AiModel
// @Summary 分页获取AI模型列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取AI模型列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /aiModel/getAiModelList [get]
export const getAiModelList = (params) => {
  return service({
    url: '/aiModel/getAiModelList',
    method: 'get',
    params
  })
}

// @Tags AiModel
// @Summary 不鉴权的AI模型接口
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /aiModel/getAiModelPublic [get]
export const getAiModelPublic = () => {
  return service({
    url: '/aiModel/getAiModelPublic',
    method: 'get'
  })
}

// @Tags AiModel
// @Summary 设置默认AI模型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"设置成功"}"
// @Router /aiModel/setDefault [post]
export const setDefaultAiModel = (params) => {
  return service({
    url: '/aiModel/setDefault',
    method: 'post',
    params
  })
}

// @Tags AiModel
// @Summary 获取默认AI模型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /aiModel/getDefault [get]
export const getDefaultAiModel = () => {
  return service({
    url: '/aiModel/getDefault',
    method: 'get'
  })
}

// @Tags AiModel
// @Summary 获取启用的AI模型列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /aiModel/getEnabled [get]
export const getEnabledAiModels = () => {
  return service({
    url: '/aiModel/getEnabled',
    method: 'get'
  })
}
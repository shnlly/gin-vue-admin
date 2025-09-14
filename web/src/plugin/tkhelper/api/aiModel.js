import service from '@/utils/request'

// @Tags AiModel
// @Summary 创建AI模型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AiModel true "创建AI模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /tkhelper/ai-model/create [post]
export const createAiModel = (data) => {
  return service({
    url: '/tkhelper/ai-model/create',
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
// @Router /tkhelper/ai-model/delete [delete]
export const deleteAiModel = (params) => {
  return service({
    url: '/tkhelper/ai-model/delete',
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
// @Router /tkhelper/ai-model/deleteByIds [delete]
export const deleteAiModelByIds = (params) => {
  return service({
    url: '/tkhelper/ai-model/deleteByIds',
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
// @Router /tkhelper/ai-model/update [put]
export const updateAiModel = (data) => {
  return service({
    url: '/tkhelper/ai-model/update',
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
// @Router /tkhelper/ai-model/find [get]
export const findAiModel = (params) => {
  return service({
    url: '/tkhelper/ai-model/find',
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
// @Router /tkhelper/ai-model/getAiModelList [get]
export const getAiModelList = (params) => {
  return service({
    url: '/tkhelper/ai-model/getAiModelList',
    method: 'get',
    params
  })
}

// @Tags AiModel
// @Summary 设置默认AI模型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"设置成功"}"
// @Router /tkhelper/ai-model/setDefault [put]
export const setDefaultAiModel = (params) => {
  return service({
    url: '/tkhelper/ai-model/setDefault',
    method: 'put',
    params
  })
}

// @Tags AiModel
// @Summary 获取默认AI模型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tkhelper/ai-model/getDefault [get]
export const getDefaultAiModel = () => {
  return service({
    url: '/tkhelper/ai-model/getDefault',
    method: 'get'
  })
}

// @Tags AiModel
// @Summary 获取启用的AI模型列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tkhelper/ai-model/getEnabled [get]
export const getEnabledAiModels = () => {
  return service({
    url: '/tkhelper/ai-model/getEnabled',
    method: 'get'
  })
}
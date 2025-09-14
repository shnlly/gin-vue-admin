import service from '@/utils/request'

// @Tags VideoContentAnalysis
// @Summary 创建视频内容分析
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.VideoContentAnalysis true "创建视频内容分析"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /tkhelper/video-analysis/create [post]
export const createVideoContentAnalysis = (data) => {
  return service({
    url: '/tkhelper/video-analysis/create',
    method: 'post',
    data: data
  })
}

// @Tags VideoContentAnalysis
// @Summary 删除视频内容分析
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param ID query string true "视频内容分析ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tkhelper/video-analysis/delete [delete]
export const deleteVideoContentAnalysis = (ID) => {
  return service({
    url: '/tkhelper/video-analysis/delete',
    method: 'delete',
    params: { ID }
  })
}

// @Tags VideoContentAnalysis
// @Summary 批量删除视频内容分析
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param IDs query []string true "视频内容分析ID数组"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /tkhelper/video-analysis/deleteByIds [delete]
export const deleteVideoContentAnalysisByIds = (IDs) => {
  return service({
    url: '/tkhelper/video-analysis/deleteByIds',
    method: 'delete',
    params: { 'IDs[]': IDs }
  })
}

// @Tags VideoContentAnalysis
// @Summary 更新视频内容分析
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.VideoContentAnalysis true "更新视频内容分析"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tkhelper/video-analysis/update [put]
export const updateVideoContentAnalysis = (data) => {
  return service({
    url: '/tkhelper/video-analysis/update',
    method: 'put',
    data: data
  })
}

// @Tags VideoContentAnalysis
// @Summary 根据ID查询视频内容分析
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param ID query string true "视频内容分析ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tkhelper/video-analysis/find [get]
export const findVideoContentAnalysis = (ID) => {
  return service({
    url: '/tkhelper/video-analysis/find',
    method: 'get',
    params: { ID }
  })
}

// @Tags VideoContentAnalysis
// @Summary 分页获取视频内容分析列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param params query request.PageInfo true "分页参数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tkhelper/video-analysis/getVideoContentAnalysisList [get]
export const getVideoContentAnalysisList = (params) => {
  return service({
    url: '/tkhelper/video-analysis/getVideoContentAnalysisList',
    method: 'get',
    params
  })
}

// @Tags VideoContentAnalysis
// @Summary 获取视频内容分析统计信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tkhelper/video-analysis/getVideoContentAnalysisStatistics [get]
export const getVideoContentAnalysisStatistics = () => {
  return service({
    url: '/tkhelper/video-analysis/getVideoContentAnalysisStatistics',
    method: 'get'
  })
}

// @Tags VideoContentAnalysis
// @Summary 从Excel导入视频内容分析数据
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce application/json
// @Param file formData file true "Excel文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"导入成功"}"
// @Router /tkhelper/video-analysis/importVideoContentAnalysisFromExcel [post]
export const importVideoContentAnalysisFromExcel = (data) => {
  return service({
    url: '/tkhelper/video-analysis/importVideoContentAnalysisFromExcel',
    method: 'post',
    data: data
  })
}

// @Tags VideoContentAnalysis
// @Summary 批量更新处理状态
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.BatchUpdateProcessStatusRequest true "批量更新请求"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tkhelper/video-analysis/batchUpdateProcessStatus [put]
export const batchUpdateProcessStatus = (data) => {
  return service({
    url: '/tkhelper/video-analysis/batchUpdateProcessStatus',
    method: 'put',
    data: data
  })
}

// @Tags VideoContentAnalysis
// @Summary 不鉴权的视频内容分析接口
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tkhelper/video-analysis/getVideoContentAnalysisPublic [get]
export const getVideoContentAnalysisPublic = () => {
  return service({
    url: '/tkhelper/video-analysis/getVideoContentAnalysisPublic',
    method: 'get'
  })
}
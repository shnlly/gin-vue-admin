import service from '@/utils/request'

// 创建视频内容分析
export const createVideoContentAnalysis = (data) => {
  return service({
    url: '/tkhelper/video-analysis/create',
    method: 'post',
    data: data
  })
}

// 删除视频内容分析
export const deleteVideoContentAnalysis = (ID) => {
  return service({
    url: '/tkhelper/video-analysis/delete',
    method: 'delete',
    params: { ID }
  })
}

// 批量删除视频内容分析
export const deleteVideoContentAnalysisByIds = (IDs) => {
  return service({
    url: '/tkhelper/video-analysis/deleteByIds',
    method: 'delete',
    params: { 'IDs[]': IDs }
  })
}

// 更新视频内容分析
export const updateVideoContentAnalysis = (data) => {
  return service({
    url: '/tkhelper/video-analysis/update',
    method: 'put',
    data: data
  })
}

// 根据ID查询视频内容分析
export const findVideoContentAnalysis = (ID) => {
  return service({
    url: '/tkhelper/video-analysis/find',
    method: 'get',
    params: { ID }
  })
}

// 分页获取视频内容分析列表
export const getVideoContentAnalysisList = (params) => {
  return service({
    url: '/tkhelper/video-analysis/getVideoContentAnalysisList',
    method: 'get',
    params
  })
}

// 获取视频内容分析统计信息
export const getVideoContentAnalysisStatistics = () => {
  return service({
    url: '/tkhelper/video-analysis/getVideoContentAnalysisStatistics',
    method: 'get'
  })
}

// 从Excel导入视频内容分析数据
export const importVideoContentAnalysisFromExcel = (data) => {
  return service({
    url: '/tkhelper/video-analysis/importVideoContentAnalysisFromExcel',
    method: 'post',
    data: data
  })
}

// 批量更新处理状态
export const batchUpdateProcessStatus = (data) => {
  return service({
    url: '/tkhelper/video-analysis/batchUpdateProcessStatus',
    method: 'put',
    data: data
  })
}

// 不鉴权的视频内容分析接口
export const getVideoContentAnalysisPublic = () => {
  return service({
    url: '/tkhelper/video-analysis/getVideoContentAnalysisPublic',
    method: 'get'
  })
}
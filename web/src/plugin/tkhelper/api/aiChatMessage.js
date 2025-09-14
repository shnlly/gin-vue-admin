import service from '@/utils/request'

// @Tags AiChatMessage
// @Summary 创建AI对话消息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AiChatMessage true "创建AI对话消息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /tkhelper/chat-message/create [post]
export const createAiChatMessage = (data) => {
  return service({
    url: '/tkhelper/chat-message/create',
    method: 'post',
    data
  })
}

// @Tags AiChatMessage
// @Summary 删除AI对话消息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AiChatMessage true "删除AI对话消息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tkhelper/chat-message/delete [delete]
export const deleteAiChatMessage = (params) => {
  return service({
    url: '/tkhelper/chat-message/delete',
    method: 'delete',
    params
  })
}

// @Tags AiChatMessage
// @Summary 批量删除AI对话消息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /tkhelper/chat-message/deleteByIds [delete]
export const deleteAiChatMessageByIds = (params) => {
  return service({
    url: '/tkhelper/chat-message/deleteByIds',
    method: 'delete',
    params
  })
}

// @Tags AiChatMessage
// @Summary 更新AI对话消息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AiChatMessage true "更新AI对话消息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tkhelper/chat-message/update [put]
export const updateAiChatMessage = (data) => {
  return service({
    url: '/tkhelper/chat-message/update',
    method: 'put',
    data
  })
}

// @Tags AiChatMessage
// @Summary 用id查询AI对话消息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.AiChatMessage true "用id查询AI对话消息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tkhelper/chat-message/find [get]
export const findAiChatMessage = (params) => {
  return service({
    url: '/tkhelper/chat-message/find',
    method: 'get',
    params
  })
}

// @Tags AiChatMessage
// @Summary 分页获取AI对话消息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取AI对话消息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tkhelper/chat-message/getAiChatMessageList [get]
export const getAiChatMessageList = (params) => {
  return service({
    url: '/tkhelper/chat-message/getAiChatMessageList',
    method: 'get',
    params
  })
}

// @Tags AiChatMessage
// @Summary 获取会话消息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param sessionId query uint true "会话ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tkhelper/chat-message/getSessionMessages [get]
export const getSessionMessages = (params) => {
  return service({
    url: '/tkhelper/chat-message/getSessionMessages',
    method: 'get',
    params
  })
}

// @Tags AiChatMessage
// @Summary 发送消息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.ChatRequest true "发送消息请求"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /tkhelper/chat-message/sendMessage [post]
export const sendMessage = (data) => {
  return service({
    url: '/tkhelper/chat-message/sendMessage',
    method: 'post',
    data
  })
}

// 发送流式消息
export const sendStreamMessage = (data, onMessage, onError, onEnd) => {
  const baseURL = service.defaults.baseURL || '/api'
  const token = localStorage.getItem('token')

  // 使用fetch来发送POST请求并处理流式响应
  fetch(`${baseURL}/tkhelper/chat-message/sendMessage`, {
    method: 'POST',
    headers: {
      'Authorization': `Bearer ${token}`,
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({ ...data, stream: true })
  })
  .then(response => {
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`)
    }

    const reader = response.body.getReader()
    const decoder = new TextDecoder()

    function readStream() {
      return reader.read().then(({ done, value }) => {
        if (done) {
          onEnd && onEnd()
          return
        }

        const chunk = decoder.decode(value, { stream: true })
        const lines = chunk.split('\n')

        for (const line of lines) {
          if (line.trim() && line.startsWith('data: ')) {
            try {
              const jsonData = JSON.parse(line.substring(6))
              if (jsonData.type === 'message') {
                onMessage && onMessage(jsonData.content)
              } else if (jsonData.type === 'end') {
                onEnd && onEnd()
                return
              }
            } catch (error) {
              console.error('解析SSE数据失败:', error, line)
            }
          }
        }

        return readStream()
      })
    }

    return readStream()
  })
  .catch(error => {
    console.error('流式请求失败:', error)
    onError && onError(error)
  })

  // 返回一个包含close方法的对象以保持接口兼容
  return {
    close: () => {
      // 这里实际上无法中止fetch请求，但保持接口兼容
      console.log('流式连接已关闭')
    }
  }
}
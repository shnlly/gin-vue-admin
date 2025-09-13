<template>
  <div class="ai-chat-container">
    <!-- 头部工具栏 -->
    <div class="chat-header">
      <div class="flex items-center justify-between">
        <h2 class="text-xl font-semibold">AI 对话</h2>
        <div class="flex items-center space-x-4">
          <!-- 模型选择 -->
          <el-select 
            v-model="selectedModelId" 
            placeholder="选择AI模型"
            @change="onModelChange"
            style="width: 200px"
          >
            <el-option
              v-for="model in availableModels"
              :key="model.ID"
              :label="model.name"
              :value="model.ID"
              :disabled="!model.enabled"
            >
              <span>{{ model.name }}</span>
              <span class="text-gray-400 text-xs ml-2">{{ model.modelName }}</span>
            </el-option>
          </el-select>
          
          <!-- 新建对话 -->
          <el-button type="primary" icon="plus" @click="createNewChat">
            新建对话
          </el-button>
        </div>
      </div>
    </div>

    <div class="chat-content">
      <!-- 侧边栏：会话列表 -->
      <div class="chat-sidebar">
        <div class="sessions-header">
          <h3 class="font-semibold mb-3">对话历史</h3>
        </div>
        
        <div class="sessions-list">
          <div 
            v-for="session in chatSessions" 
            :key="session.ID"
            :class="['session-item', { 'active': currentSessionId === session.ID }]"
            @click="selectSession(session.ID)"
          >
            <div class="session-title">{{ session.title }}</div>
            <div class="session-meta">
              <span class="text-xs text-gray-500">{{ formatDate(session.UpdatedAt) }}</span>
              <el-button 
                type="danger" 
                link 
                size="small" 
                icon="delete"
                @click.stop="deleteSession(session.ID)"
              />
            </div>
          </div>
          
          <div v-if="!chatSessions.length" class="empty-sessions">
            <span class="text-gray-400">暂无对话记录</span>
          </div>
        </div>
      </div>

      <!-- 主聊天区域 -->
      <div class="chat-main">
        <!-- 消息列表 -->
        <div class="messages-container" ref="messagesContainer">
          <div v-if="!currentMessages.length && currentSessionId" class="empty-messages">
            <div class="text-center text-gray-400 py-8">
              <el-icon size="48" class="mb-4"><chat-dot-round /></el-icon>
              <p>开始你的AI对话吧！</p>
            </div>
          </div>
          
          <div v-for="message in currentMessages" :key="message.ID" class="message-item">
            <div :class="['message', message.role]">
              <div class="message-avatar">
                <el-avatar v-if="message.role === 'user'" icon="user" />
                <el-avatar v-else>
                  <el-icon><robot /></el-icon>
                </el-avatar>
              </div>
              
              <div class="message-content">
                <div class="message-header">
                  <span class="role-name">{{ message.role === 'user' ? '用户' : 'AI助手' }}</span>
                  <span class="message-time">{{ formatDate(message.CreatedAt) }}</span>
                </div>
                
                <div class="message-text" v-html="formatMessageContent(message.content)"></div>
                
                <div v-if="message.tokenCount > 0" class="message-meta">
                  <span class="text-xs text-gray-400">Token: {{ message.tokenCount }}</span>
                </div>
              </div>
            </div>
          </div>
          
          <!-- 流式输出显示 -->
          <div v-if="isStreaming" class="message-item">
            <div class="message assistant">
              <div class="message-avatar">
                <el-avatar>
                  <el-icon><robot /></el-icon>
                </el-avatar>
              </div>
              
              <div class="message-content">
                <div class="message-header">
                  <span class="role-name">AI助手</span>
                  <span class="streaming-indicator">
                    <el-icon class="animate-spin"><loading /></el-icon>
                    正在回复...
                  </span>
                </div>
                
                <div class="message-text" v-html="formatMessageContent(streamingContent)"></div>
              </div>
            </div>
          </div>
        </div>

        <!-- 输入区域 -->
        <div class="input-container">
          <div class="input-wrapper">
            <el-input
              v-model="inputMessage"
              type="textarea"
              :rows="3"
              :disabled="!selectedModelId || isStreaming"
              placeholder="输入你的消息... (Ctrl+Enter 发送)"
              resize="none"
              @keydown="handleKeyDown"
            />
            
            <div class="input-actions">
              <div class="flex items-center space-x-2">
                <el-checkbox v-model="streamMode">流式输出</el-checkbox>
                <span class="text-xs text-gray-400" v-if="selectedModel">
                  当前模型：{{ selectedModel.name }}
                </span>
              </div>
              
              <el-button 
                type="primary" 
                :disabled="!inputMessage.trim() || !selectedModelId || isStreaming"
                :loading="isStreaming"
                @click="sendMessage"
              >
                <template v-if="isStreaming">
                  <el-icon class="animate-spin"><loading /></el-icon>
                  发送中...
                </template>
                <template v-else>
                  发送
                </template>
              </el-button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted, computed, nextTick, watch } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { 
  getUserChatSessions, 
  createAiChatSession, 
  deleteAiChatSession
} from '@/api/example/aiChatSession'
import { 
  getSessionMessages,
  sendMessage as sendChatMessage,
  sendStreamMessage
} from '@/api/example/aiChatMessage'
import { getAiModelList } from '@/api/example/aiModel'

// 响应式数据
const selectedModelId = ref('')
const availableModels = ref([])
const chatSessions = ref([])
const currentSessionId = ref('')
const currentMessages = ref([])
const inputMessage = ref('')
const streamMode = ref(true)
const isStreaming = ref(false)
const streamingContent = ref('')
const messagesContainer = ref()

// 计算属性
const selectedModel = computed(() => {
  return availableModels.value.find(model => model.ID === selectedModelId.value)
})

// 初始化
onMounted(async () => {
  await loadModels()
  await loadSessions()
})

// 监听会话变化
watch(currentSessionId, async (newSessionId) => {
  if (newSessionId) {
    await loadMessages(newSessionId)
  } else {
    currentMessages.value = []
  }
})

// 加载可用模型
const loadModels = async () => {
  try {
    const res = await getAiModelList({ page: 1, pageSize: 100 })
    if (res.code === 0) {
      availableModels.value = res.data.list.filter(model => model.enabled)
      // 选择默认模型
      const defaultModel = availableModels.value.find(model => model.isDefault) || availableModels.value[0]
      if (defaultModel) {
        selectedModelId.value = defaultModel.ID
      }
    }
  } catch (error) {
    console.error('加载模型失败:', error)
    ElMessage.error('加载AI模型失败')
  }
}

// 加载用户会话列表
const loadSessions = async () => {
  try {
    const res = await getUserChatSessions()
    if (res.code === 0) {
      chatSessions.value = res.data || []
      // 如果有会话，选择最新的一个
      if (chatSessions.value.length > 0) {
        currentSessionId.value = chatSessions.value[0].ID
      }
    }
  } catch (error) {
    console.error('加载会话失败:', error)
  }
}

// 加载会话消息
const loadMessages = async (sessionId) => {
  try {
    const res = await getSessionMessages({ sessionId })
    if (res.code === 0) {
      currentMessages.value = res.data || []
      await nextTick()
      scrollToBottom()
    }
  } catch (error) {
    console.error('加载消息失败:', error)
  }
}

// 创建新对话
const createNewChat = async () => {
  if (!selectedModelId.value) {
    ElMessage.warning('请先选择AI模型')
    return
  }
  
  try {
    const model = selectedModel.value
    const res = await createAiChatSession({
      title: `与${model.name}的对话`,
      modelId: selectedModelId.value,
      modelName: model.modelName
    })
    
    if (res.code === 0) {
      ElMessage.success('创建对话成功')
      await loadSessions()
      currentSessionId.value = res.data.ID
    }
  } catch (error) {
    console.error('创建会话失败:', error)
    ElMessage.error('创建对话失败')
  }
}

// 选择会话
const selectSession = (sessionId) => {
  currentSessionId.value = sessionId
}

// 删除会话
const deleteSession = async (sessionId) => {
  try {
    await ElMessageBox.confirm('确定要删除这个对话吗？', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
    
    const res = await deleteAiChatSession({ ID: sessionId })
    if (res.code === 0) {
      ElMessage.success('删除成功')
      await loadSessions()
      if (currentSessionId.value === sessionId) {
        currentSessionId.value = chatSessions.value[0]?.ID || ''
      }
    }
  } catch (error) {
    if (error !== 'cancel') {
      console.error('删除会话失败:', error)
      ElMessage.error('删除失败')
    }
  }
}

// 模型变化处理
const onModelChange = async () => {
  // 如果当前有会话但模型不匹配，提示创建新会话
  if (currentSessionId.value) {
    const currentSession = chatSessions.value.find(s => s.ID === currentSessionId.value)
    if (currentSession && currentSession.modelId !== selectedModelId.value) {
      ElMessage.info('模型已切换，建议创建新对话以使用新模型')
    }
  }
}

// 发送消息
const sendMessage = async () => {
  if (!inputMessage.value.trim()) return
  if (!currentSessionId.value) {
    ElMessage.warning('请先创建对话')
    return
  }
  
  const message = inputMessage.value.trim()
  inputMessage.value = ''
  
  // 添加用户消息到界面
  const userMessage = {
    ID: Date.now(),
    role: 'user',
    content: message,
    CreatedAt: new Date().toISOString(),
    sessionId: currentSessionId.value
  }
  currentMessages.value.push(userMessage)
  await nextTick()
  scrollToBottom()
  
  isStreaming.value = true
  streamingContent.value = ''
  
  try {
    if (streamMode.value) {
      // 流式输出
      const eventSource = sendStreamMessage(
        {
          sessionId: currentSessionId.value,
          message: message,
          modelId: selectedModelId.value,
          stream: true
        },
        // onMessage callback
        (content) => {
          streamingContent.value += content
          nextTick(() => scrollToBottom())
        },
        // onError callback
        (error) => {
          console.error('流式响应错误:', error)
          ElMessage.error('发送消息失败')
          isStreaming.value = false
        },
        // onEnd callback
        () => {
          isStreaming.value = false
          streamingContent.value = ''
          loadMessages(currentSessionId.value)
          loadSessions() // 更新会话列表以显示最新消息
        }
      )
    } else {
      // 普通发送
      const res = await sendChatMessage({
        sessionId: currentSessionId.value,
        message: message,
        modelId: selectedModelId.value,
        stream: false
      })
      
      if (res.code === 0) {
        await loadMessages(currentSessionId.value)
        await loadSessions()
      } else {
        ElMessage.error(res.msg || '发送消息失败')
      }
      isStreaming.value = false
    }
  } catch (error) {
    console.error('发送消息失败:', error)
    ElMessage.error('发送消息失败')
    isStreaming.value = false
  }
}

// 键盘事件处理
const handleKeyDown = (event) => {
  if ((event.ctrlKey || event.metaKey) && event.key === 'Enter') {
    event.preventDefault()
    sendMessage()
  }
}

// 滚动到底部
const scrollToBottom = () => {
  if (messagesContainer.value) {
    messagesContainer.value.scrollTop = messagesContainer.value.scrollHeight
  }
}

// 格式化日期
const formatDate = (dateStr) => {
  if (!dateStr) return ''
  const date = new Date(dateStr)
  return date.toLocaleString('zh-CN', {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit',
    hour: '2-digit',
    minute: '2-digit'
  })
}

// 格式化消息内容（支持简单的Markdown）
const formatMessageContent = (content) => {
  if (!content) return ''
  
  return content
    .replace(/\n/g, '<br>')
    .replace(/`([^`]+)`/g, '<code>$1</code>')
    .replace(/\*\*(.*?)\*\*/g, '<strong>$1</strong>')
    .replace(/\*(.*?)\*/g, '<em>$1</em>')
}
</script>

<style scoped>
.ai-chat-container {
  height: calc(100vh - 120px);
  display: flex;
  flex-direction: column;
  background: #f5f5f5;
}

.chat-header {
  background: white;
  padding: 16px 24px;
  border-bottom: 1px solid #e5e7eb;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
}

.chat-content {
  flex: 1;
  display: flex;
  min-height: 0;
}

.chat-sidebar {
  width: 280px;
  background: white;
  border-right: 1px solid #e5e7eb;
  display: flex;
  flex-direction: column;
}

.sessions-header {
  padding: 16px;
  border-bottom: 1px solid #e5e7eb;
}

.sessions-list {
  flex: 1;
  overflow-y: auto;
  padding: 8px;
}

.session-item {
  padding: 12px;
  margin-bottom: 8px;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.2s;
  border: 1px solid transparent;
}

.session-item:hover {
  background: #f3f4f6;
}

.session-item.active {
  background: #e3f2fd;
  border-color: #2196f3;
}

.session-title {
  font-weight: 500;
  margin-bottom: 4px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.session-meta {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.empty-sessions {
  text-align: center;
  padding: 32px 16px;
}

.chat-main {
  flex: 1;
  display: flex;
  flex-direction: column;
  background: white;
  min-width: 0;
}

.messages-container {
  flex: 1;
  overflow-y: auto;
  padding: 16px;
  min-height: 0;
}

.empty-messages {
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
}

.message-item {
  margin-bottom: 24px;
}

.message {
  display: flex;
  align-items: flex-start;
  gap: 12px;
}

.message.user {
  flex-direction: row-reverse;
}

.message.user .message-content {
  background: #2196f3;
  color: white;
  margin-left: auto;
}

.message.assistant .message-content {
  background: #f5f5f5;
  color: #333;
}

.message-avatar {
  flex-shrink: 0;
}

.message-content {
  max-width: 70%;
  padding: 12px 16px;
  border-radius: 12px;
  word-wrap: break-word;
}

.message-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 8px;
  font-size: 12px;
  opacity: 0.8;
}

.role-name {
  font-weight: 500;
}

.message-time {
  font-size: 11px;
}

.streaming-indicator {
  display: flex;
  align-items: center;
  gap: 4px;
  color: #2196f3;
}

.message-text {
  line-height: 1.6;
}

.message-text :deep(code) {
  background: rgba(0, 0, 0, 0.1);
  padding: 2px 4px;
  border-radius: 4px;
  font-family: 'Courier New', monospace;
}

.message-meta {
  margin-top: 8px;
  text-align: right;
}

.input-container {
  border-top: 1px solid #e5e7eb;
  padding: 16px;
  background: #fafafa;
}

.input-wrapper {
  background: white;
  border-radius: 12px;
  border: 1px solid #e5e7eb;
  overflow: hidden;
}

.input-wrapper:focus-within {
  border-color: #2196f3;
  box-shadow: 0 0 0 2px rgba(33, 150, 243, 0.1);
}

.input-wrapper :deep(.el-textarea__inner) {
  border: none;
  padding: 16px;
  box-shadow: none;
  resize: none;
  font-size: 14px;
}

.input-actions {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 16px;
  border-top: 1px solid #e5e7eb;
  background: #fafafa;
}

@keyframes spin {
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
}

.animate-spin {
  animation: spin 1s linear infinite;
}
</style>
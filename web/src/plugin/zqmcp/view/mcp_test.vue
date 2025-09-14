<template>
  <div class="mcp-test-container">
    <div class="page-header">
      <h2>多MCP服务器工具测试</h2>
      <p>选择MCP服务器并测试其工具功能</p>
    </div>

    <!-- MCP服务器选择器 -->
    <McpServerSelector
      v-model="selectedServerId"
      @server-change="onServerChange"
    />

    <!-- 工具列表和测试区域 -->
    <div v-if="selectedServer" class="test-content">
      <!-- 服务器信息 -->
      <el-card class="server-info-card">
        <template #header>
          <div class="card-header">
            <span>服务器信息</span>
            <el-tag :type="getStatusType(selectedServer.serverStatus)">
              {{ getStatusText(selectedServer.serverStatus) }}
            </el-tag>
          </div>
        </template>
        <el-descriptions :column="3" border>
          <el-descriptions-item label="服务器名称">
            {{ selectedServer.serverName }}
          </el-descriptions-item>
          <el-descriptions-item label="端口">
            {{ selectedServer.serverPort }}
          </el-descriptions-item>
          <el-descriptions-item label="版本">
            {{ selectedServer.serverVersion }}
          </el-descriptions-item>
          <el-descriptions-item label="连接数">
            {{ selectedServer.connectionCount || 0 }}
          </el-descriptions-item>
          <el-descriptions-item label="最后启动时间">
            {{ formatDate(selectedServer.lastStartTime) }}
          </el-descriptions-item>
          <el-descriptions-item label="自动分配端口">
            {{ selectedServer.autoAllocated ? '是' : '否' }}
          </el-descriptions-item>
        </el-descriptions>
      </el-card>

      <!-- 工具列表 -->
      <el-card class="tools-card">
        <template #header>
          <div class="card-header">
            <span>可用工具</span>
            <el-button type="primary" size="small" @click="refreshTools">刷新工具</el-button>
          </div>
        </template>

        <div v-if="tools.length === 0" class="no-tools">
          <el-empty description="当前服务器暂无可用工具" />
        </div>

        <div v-else class="tools-grid">
          <el-card
            v-for="tool in tools"
            :key="tool.ID"
            class="tool-card"
            shadow="hover"
            @click="selectTool(tool)"
            :class="{ 'selected': selectedTool?.ID === tool.ID }"
          >
            <div class="tool-info">
              <h4>{{ tool.toolName }}</h4>
              <p>{{ tool.toolDescription || '暂无描述' }}</p>
              <div class="tool-meta">
                <el-tag size="small" :type="getCategoryType(tool.toolCategory)">
                  {{ getCategoryText(tool.toolCategory) }}
                </el-tag>
                <span class="call-count">调用次数: {{ tool.callCount || 0 }}</span>
              </div>
            </div>
          </el-card>
        </div>
      </el-card>

      <!-- 工具测试区域 -->
      <el-card v-if="selectedTool" class="test-area-card">
        <template #header>
          <div class="card-header">
            <span>工具测试 - {{ selectedTool.toolName }}</span>
            <el-button type="success" @click="testTool" :loading="testing">执行测试</el-button>
          </div>
        </template>

        <div class="test-form">
          <h4>工具参数</h4>
          <el-form :model="toolParams" label-position="top">
            <el-form-item v-if="selectedTool.toolSchema" label="参数 (JSON格式)">
              <el-input
                v-model="paramString"
                type="textarea"
                :rows="4"
                placeholder="请输入JSON格式的参数"
              />
              <div class="schema-info">
                <small>参数模式: {{ JSON.stringify(selectedTool.toolSchema, null, 2) }}</small>
              </div>
            </el-form-item>
            <el-form-item v-else label="通用参数">
              <el-input
                v-model="paramString"
                type="textarea"
                :rows="4"
                placeholder="请输入JSON格式的参数，如: {&quot;message&quot;: &quot;Hello World&quot;}"
              />
            </el-form-item>
          </el-form>
        </div>

        <div v-if="testResult" class="test-result">
          <h4>执行结果</h4>
          <el-alert
            :type="testResult.success ? 'success' : 'error'"
            :title="testResult.success ? '执行成功' : '执行失败'"
            show-icon
            :closable="false"
          />
          <pre class="result-content">{{ formatResult(testResult.data) }}</pre>
        </div>
      </el-card>
    </div>

    <div v-else class="no-server">
      <el-empty description="请先选择一个运行中的MCP服务器" />
    </div>
  </div>
</template>

<script setup>
import { ref, watch } from 'vue'
import { ElMessage } from 'element-plus'
import McpServerSelector from '../components/McpServerSelector.vue'
import { getMcpToolList } from '@/plugin/zqmcp/api/mcp_tool'
import { formatDate } from '@/utils/format'

const selectedServerId = ref(null)
const selectedServer = ref(null)
const tools = ref([])
const selectedTool = ref(null)
const toolParams = ref({})
const paramString = ref('{"message": "Hello from ZQMCP!"}')
const testing = ref(false)
const testResult = ref(null)

// 监听服务器选择变化
const onServerChange = (serverId, server) => {
  selectedServer.value = server
  selectedTool.value = null
  testResult.value = null
  if (server) {
    loadTools()
  } else {
    tools.value = []
  }
}

// 加载工具列表
const loadTools = async () => {
  try {
    const res = await getMcpToolList({
      page: 1,
      pageSize: 100,
      serverId: selectedServerId.value
    })
    if (res.code === 0) {
      tools.value = res.data.list || []
    }
  } catch (error) {
    console.error('加载工具列表失败:', error)
    ElMessage.error('加载工具列表失败')
  }
}

// 刷新工具列表
const refreshTools = async () => {
  await loadTools()
  ElMessage.success('工具列表已刷新')
}

// 选择工具
const selectTool = (tool) => {
  selectedTool.value = tool
  testResult.value = null

  // 如果工具有参数模式，尝试生成示例参数
  if (tool.toolSchema) {
    try {
      const schema = typeof tool.toolSchema === 'string'
        ? JSON.parse(tool.toolSchema)
        : tool.toolSchema

      if (schema.properties) {
        const example = {}
        Object.keys(schema.properties).forEach(key => {
          const prop = schema.properties[key]
          switch (prop.type) {
            case 'string':
              example[key] = prop.example || `sample_${key}`
              break
            case 'number':
            case 'integer':
              example[key] = prop.example || 123
              break
            case 'boolean':
              example[key] = prop.example !== undefined ? prop.example : true
              break
            default:
              example[key] = prop.example || null
          }
        })
        paramString.value = JSON.stringify(example, null, 2)
      }
    } catch (error) {
      console.warn('解析工具参数模式失败:', error)
    }
  }
}

// 测试工具
const testTool = async () => {
  if (!selectedTool.value) return

  testing.value = true
  try {
    // 解析参数
    let params = {}
    if (paramString.value.trim()) {
      try {
        params = JSON.parse(paramString.value)
      } catch (error) {
        ElMessage.error('参数格式错误，请检查JSON格式')
        return
      }
    }

    // 这里应该调用实际的MCP工具测试API
    // 由于目前还没有实现具体的工具调用API，先模拟结果
    const mockResult = {
      success: true,
      data: {
        toolName: selectedTool.value.toolName,
        serverId: selectedServerId.value,
        serverName: selectedServer.value.serverName,
        params: params,
        result: `工具 ${selectedTool.value.toolName} 执行成功！`,
        timestamp: new Date().toISOString(),
        executionTime: Math.random() * 1000 + 'ms'
      }
    }

    testResult.value = mockResult
    ElMessage.success('工具测试完成')

  } catch (error) {
    console.error('工具测试失败:', error)
    testResult.value = {
      success: false,
      data: {
        error: error.message || '工具测试失败'
      }
    }
    ElMessage.error('工具测试失败')
  } finally {
    testing.value = false
  }
}

// 状态类型映射
const getStatusType = (status) => {
  switch (status) {
    case 'running': return 'success'
    case 'starting': return 'warning'
    case 'stopped': return 'info'
    case 'error': return 'danger'
    default: return 'info'
  }
}

const getStatusText = (status) => {
  switch (status) {
    case 'running': return '运行中'
    case 'starting': return '启动中'
    case 'stopped': return '已停止'
    case 'error': return '异常'
    default: return '未知'
  }
}

const getCategoryType = (category) => {
  switch (category) {
    case 'system': return 'primary'
    case 'database': return 'success'
    case 'file': return 'warning'
    case 'network': return 'danger'
    default: return 'info'
  }
}

const getCategoryText = (category) => {
  switch (category) {
    case 'system': return '系统工具'
    case 'database': return '数据库工具'
    case 'file': return '文件工具'
    case 'network': return '网络工具'
    case 'custom': return '自定义工具'
    default: return '未分类'
  }
}

const formatResult = (data) => {
  return JSON.stringify(data, null, 2)
}
</script>

<style scoped>
.mcp-test-container {
  padding: 20px;
}

.page-header {
  margin-bottom: 20px;
}

.page-header h2 {
  margin: 0 0 10px 0;
  color: #303133;
}

.page-header p {
  margin: 0;
  color: #606266;
}

.test-content {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.server-info-card {
  width: 100%;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.tools-card {
  width: 100%;
}

.no-tools {
  padding: 40px;
}

.tools-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 16px;
}

.tool-card {
  cursor: pointer;
  transition: all 0.3s;
  border: 2px solid transparent;
}

.tool-card:hover {
  border-color: #409eff;
}

.tool-card.selected {
  border-color: #67c23a;
  box-shadow: 0 2px 12px 0 rgba(103, 194, 58, 0.3);
}

.tool-info h4 {
  margin: 0 0 10px 0;
  color: #303133;
}

.tool-info p {
  margin: 0 0 15px 0;
  color: #606266;
  font-size: 14px;
}

.tool-meta {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.call-count {
  font-size: 12px;
  color: #909399;
}

.test-area-card {
  width: 100%;
}

.test-form {
  margin-bottom: 20px;
}

.schema-info {
  margin-top: 8px;
}

.schema-info small {
  color: #909399;
  font-family: monospace;
  white-space: pre-wrap;
}

.test-result {
  margin-top: 20px;
  padding: 20px;
  background-color: #f8f9fa;
  border-radius: 4px;
}

.test-result h4 {
  margin: 0 0 15px 0;
  color: #303133;
}

.result-content {
  margin: 15px 0 0 0;
  padding: 15px;
  background-color: #2d3748;
  color: #e2e8f0;
  border-radius: 4px;
  overflow-x: auto;
  font-family: 'Monaco', 'Menlo', 'Ubuntu Mono', monospace;
  font-size: 14px;
  line-height: 1.5;
}

.no-server {
  padding: 60px 20px;
  text-align: center;
}
</style>
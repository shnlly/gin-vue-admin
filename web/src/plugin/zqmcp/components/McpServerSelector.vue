<template>
  <div class="mcp-server-selector">
    <el-select
      v-model="selectedServer"
      placeholder="选择MCP服务器"
      @change="handleServerChange"
      style="width: 300px"
    >
      <el-option
        v-for="server in servers"
        :key="server.ID"
        :label="getServerLabel(server)"
        :value="server.ID"
        :disabled="server.serverStatus !== 'running'"
      >
        <div class="server-option">
          <div class="server-info">
            <span class="server-name">{{ server.serverName }}</span>
            <span class="server-port">:{{ server.serverPort }}</span>
          </div>
          <div class="server-status">
            <el-tag
              :type="getStatusType(server.serverStatus)"
              size="mini"
            >
              {{ getStatusText(server.serverStatus) }}
            </el-tag>
          </div>
        </div>
      </el-option>
    </el-select>

    <el-button
      type="primary"
      size="mini"
      @click="refreshServers"
      :loading="loading"
      style="margin-left: 10px"
    >
      刷新
    </el-button>

    <el-button
      type="success"
      size="mini"
      @click="startServer"
      :disabled="!selectedServer || getServerById(selectedServer)?.serverStatus === 'running'"
      style="margin-left: 10px"
    >
      启动
    </el-button>

    <el-button
      type="warning"
      size="mini"
      @click="stopServer"
      :disabled="!selectedServer || getServerById(selectedServer)?.serverStatus !== 'running'"
      style="margin-left: 10px"
    >
      停止
    </el-button>
  </div>
</template>

<script>
import { getMcpServerList, startMcpServer, stopMcpServer } from '@/plugin/zqmcp/api/mcp_server'

export default {
  name: 'McpServerSelector',
  props: {
    value: {
      type: [Number, String],
      default: null
    }
  },
  data() {
    return {
      selectedServer: null,
      servers: [],
      loading: false
    }
  },
  watch: {
    value: {
      immediate: true,
      handler(val) {
        this.selectedServer = val
      }
    }
  },
  mounted() {
    this.loadServers()
  },
  methods: {
    async loadServers() {
      this.loading = true
      try {
        const res = await getMcpServerList({
          page: 1,
          pageSize: 100
        })
        if (res.code === 0) {
          this.servers = res.data.list || []
          // 如果没有选中的服务器且有可用服务器，自动选择第一个运行中的
          if (!this.selectedServer && this.servers.length > 0) {
            const runningServer = this.servers.find(s => s.serverStatus === 'running')
            if (runningServer) {
              this.selectedServer = runningServer.ID
              this.handleServerChange(this.selectedServer)
            }
          }
        }
      } catch (error) {
        console.error('加载服务器列表失败:', error)
        this.$message.error('加载服务器列表失败')
      } finally {
        this.loading = false
      }
    },

    handleServerChange(serverId) {
      this.$emit('input', serverId)
      this.$emit('server-change', serverId, this.getServerById(serverId))
    },

    getServerById(id) {
      return this.servers.find(s => s.ID === id)
    },

    getServerLabel(server) {
      return `${server.serverName} (${server.serverPort})`
    },

    getStatusType(status) {
      switch (status) {
        case 'running': return 'success'
        case 'starting': return 'warning'
        case 'stopped': return 'info'
        case 'error': return 'danger'
        default: return 'info'
      }
    },

    getStatusText(status) {
      switch (status) {
        case 'running': return '运行中'
        case 'starting': return '启动中'
        case 'stopped': return '已停止'
        case 'error': return '异常'
        default: return '未知'
      }
    },

    async refreshServers() {
      await this.loadServers()
      this.$message.success('服务器列表已刷新')
    },

    async startServer() {
      if (!this.selectedServer) return

      try {
        const res = await startMcpServer(this.selectedServer)
        if (res.code === 0) {
          this.$message.success('服务器启动成功')
          await this.loadServers()
        } else {
          this.$message.error(res.msg || '服务器启动失败')
        }
      } catch (error) {
        console.error('启动服务器失败:', error)
        this.$message.error('启动服务器失败')
      }
    },

    async stopServer() {
      if (!this.selectedServer) return

      try {
        const res = await stopMcpServer(this.selectedServer)
        if (res.code === 0) {
          this.$message.success('服务器停止成功')
          await this.loadServers()
        } else {
          this.$message.error(res.msg || '服务器停止失败')
        }
      } catch (error) {
        console.error('停止服务器失败:', error)
        this.$message.error('停止服务器失败')
      }
    }
  }
}
</script>

<style scoped>
.mcp-server-selector {
  display: flex;
  align-items: center;
  margin-bottom: 20px;
}

.server-option {
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100%;
}

.server-info {
  display: flex;
  align-items: center;
}

.server-name {
  font-weight: bold;
  margin-right: 5px;
}

.server-port {
  color: #666;
  font-size: 12px;
}

.server-status {
  margin-left: 10px;
}
</style>
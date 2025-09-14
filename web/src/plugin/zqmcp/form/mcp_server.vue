
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="MCP服务器名称:" prop="serverName">
          <el-input v-model="formData.serverName" :clearable="true"  placeholder="请输入MCP服务器名称" />
       </el-form-item>
        <el-form-item label="服务器端口号:" prop="serverPort">
          <div style="display: flex; gap: 10px;">
            <el-input v-model.number="formData.serverPort" :clearable="true" placeholder="留空自动分配端口" />
            <el-button type="primary" @click="autoAllocatePort" :loading="allocatingPort">自动分配</el-button>
          </div>
       </el-form-item>
        <el-form-item label="服务器版本:" prop="serverVersion">
          <el-input v-model="formData.serverVersion" :clearable="true"  placeholder="请输入服务器版本" />
       </el-form-item>
        <el-form-item label="服务器状态:" prop="serverStatus">
           <el-select v-model="formData.serverStatus" placeholder="请选择服务器状态" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in mcp_server_statusOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="是否自动分配端口:" prop="autoAllocated">
          <el-switch v-model="formData.autoAllocated" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="服务器配置:" prop="serverConfig">
          // 此字段为json结构，可以前端自行控制展示和数据绑定模式 需绑定json的key为 formData.serverConfig 后端会按照json的类型进行存取
          {{ formData.serverConfig }}
       </el-form-item>
        <el-form-item label="服务器描述:" prop="description">
          <el-input v-model="formData.description" :clearable="true"  placeholder="请输入服务器描述" />
       </el-form-item>
        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createMcpServer,
  updateMcpServer,
  findMcpServer,
  allocatePort
} from '@/plugin/zqmcp/api/mcp_server'

import { ElMessage } from 'element-plus'

defineOptions({
    name: 'McpServerForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)
const allocatingPort = ref(false)

const type = ref('')
const mcp_server_statusOptions = ref([])
const formData = ref({
            serverName: '',
            serverPort: 0,
            serverVersion: '',
            serverStatus: '',
            autoAllocated: false,
            serverConfig: {},
            description: '',
        })
// 验证规则
const rule = reactive({
               serverName : [{
                   required: true,
                   message: '请输入服务器名称',
                   trigger: ['input','blur'],
               }],
               serverPort : [{
                   required: true,
                   message: '请输入端口号',
                   trigger: ['input','blur'],
               }],
               serverStatus : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findMcpServer({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    mcp_server_statusOptions.value = await getDictFunc('mcp_server_status')
}

init()
// 保存按钮
const save = async() => {
      btnLoading.value = true
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return btnLoading.value = false
            let res
           switch (type.value) {
             case 'create':
               res = await createMcpServer(formData.value)
               break
             case 'update':
               res = await updateMcpServer(formData.value)
               break
             default:
               res = await createMcpServer(formData.value)
               break
           }
           btnLoading.value = false
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

// 自动分配端口
const autoAllocatePort = async () => {
  try {
    allocatingPort.value = true
    const res = await allocatePort()
    if (res.code === 0) {
      formData.value.serverPort = res.data.port
      formData.value.autoAllocated = true
      ElMessage({
        type: 'success',
        message: `已分配端口：${res.data.port}`
      })
    } else {
      ElMessage({
        type: 'error',
        message: res.msg || '端口分配失败'
      })
    }
  } catch (error) {
    console.error('分配端口失败:', error)
    ElMessage({
      type: 'error',
      message: '端口分配失败'
    })
  } finally {
    allocatingPort.value = false
  }
}

</script>

<style>
</style>

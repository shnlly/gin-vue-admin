
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="所属MCP服务器ID:" prop="serverId">
        <el-select  multiple  v-model="formData.serverId" placeholder="请选择所属MCP服务器ID" style="width:100%" :clearable="true" >
          <el-option v-for="(item,key) in dataSource.serverId" :key="key" :label="item.label" :value="item.value" />
        </el-select>
       </el-form-item>
        <el-form-item label="工具名称:" prop="toolName">
          <el-input v-model="formData.toolName" :clearable="true"  placeholder="请输入工具名称" />
       </el-form-item>
        <el-form-item label="工具描述:" prop="toolDescription">
          <el-input v-model="formData.toolDescription" :clearable="true"  placeholder="请输入工具描述" />
       </el-form-item>
        <el-form-item label="工具参数模式:" prop="toolSchema">
          // 此字段为json结构，可以前端自行控制展示和数据绑定模式 需绑定json的key为 formData.toolSchema 后端会按照json的类型进行存取
          {{ formData.toolSchema }}
       </el-form-item>
        <el-form-item label="工具分类:" prop="toolCategory">
           <el-select v-model="formData.toolCategory" placeholder="请选择工具分类" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in mcp_tool_categoryOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="是否启用:" prop="isEnabled">
          <el-switch v-model="formData.isEnabled" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
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
    getMcpToolDataSource,
  createMcpTool,
  updateMcpTool,
  findMcpTool
} from '@/plugin/zqmcp/api/mcp_tool'

defineOptions({
    name: 'McpToolForm'
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

const type = ref('')
const mcp_tool_categoryOptions = ref([])
const formData = ref({
            serverId: undefined,
            toolName: '',
            toolDescription: '',
            toolSchema: {},
            toolCategory: '',
            isEnabled: false,
        })
// 验证规则
const rule = reactive({
               serverId : [{
                   required: true,
                   message: '请选择MCP服务器',
                   trigger: ['input','blur'],
               }],
               toolName : [{
                   required: true,
                   message: '请输入工具名称',
                   trigger: ['input','blur'],
               }],
               toolCategory : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()
  const dataSource = ref([])
  const getDataSourceFunc = async()=>{
    const res = await getMcpToolDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findMcpTool({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    mcp_tool_categoryOptions.value = await getDictFunc('mcp_tool_category')
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
               res = await createMcpTool(formData.value)
               break
             case 'update':
               res = await updateMcpTool(formData.value)
               break
             default:
               res = await createMcpTool(formData.value)
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

</script>

<style>
</style>

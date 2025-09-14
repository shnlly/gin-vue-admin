
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
      <el-form-item label="创建日期" prop="createdAtRange">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>
         <el-date-picker
                  v-model="searchInfo.createdAtRange"
                  class="!w-380px"
                  type="datetimerange"
                  range-separator="至"
                  start-placeholder="开始时间"
                  end-placeholder="结束时间"
                />
       </el-form-item>
      
            <el-form-item label="MCP服务器名称" prop="serverName">
  <el-input v-model="searchInfo.serverName" placeholder="搜索条件" />
</el-form-item>
           
            <el-form-item label="服务器端口号" prop="serverPort">
  <el-input v-model.number="searchInfo.serverPort" placeholder="搜索条件" />
</el-form-item>
           
            <el-form-item label="服务器状态" prop="serverStatus">
  <el-select v-model="searchInfo.serverStatus" clearable filterable placeholder="请选择" @clear="()=>{searchInfo.serverStatus=undefined}">
    <el-option v-for="(item,key) in mcp_server_statusOptions" :key="key" :label="item.label" :value="item.value" />
  </el-select>
</el-form-item>
           
        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
          <el-form-item label="服务器版本" prop="serverVersion">
  <el-input v-model="searchInfo.serverVersion" placeholder="搜索条件" />
</el-form-item>
          
          <el-form-item label="是否自动分配端口" prop="autoAllocated">
  <el-select v-model="searchInfo.autoAllocated" clearable placeholder="请选择">
    <el-option key="true" label="是" value="true"></el-option>
    <el-option key="false" label="否" value="false"></el-option>
  </el-select>
</el-form-item>
          
          <el-form-item label="服务器描述" prop="description">
  <el-input v-model="searchInfo.description" placeholder="搜索条件" />
</el-form-item>
          
          <el-form-item label="最后启动时间" prop="lastStartTime">
  <template #label>
    <span>
      最后启动时间
      <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
        <el-icon><QuestionFilled /></el-icon>
      </el-tooltip>
    </span>
  </template>
<el-date-picker class="!w-380px" v-model="searchInfo.lastStartTimeRange" type="datetimerange" range-separator="至"  start-placeholder="开始时间" end-placeholder="结束时间"></el-date-picker></el-form-item>
          
          <el-form-item label="当前连接数" prop="connectionCount">
  <el-input v-model.number="searchInfo.connectionCount" placeholder="搜索条件" />
</el-form-item>
          
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开</el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button v-auth="btnAuth.add" type="primary" icon="plus" @click="openDialog()">新增</el-button>
            <el-button v-auth="btnAuth.batchDelete" icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
            <ExportTemplate v-auth="btnAuth.exportTemplate" template-id="zqmcp_McpServer" />
            <ExportExcel v-auth="btnAuth.exportExcel" template-id="zqmcp_McpServer" filterDeleted/>
            <ImportExcel v-auth="btnAuth.importExcel" template-id="zqmcp_McpServer" @on-success="getTableData" />
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        @sort-change="sortChange"
        >
        <el-table-column type="selection" width="55" />
        
        <el-table-column sortable align="left" label="日期" prop="CreatedAt" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
            <el-table-column align="left" label="MCP服务器名称" prop="serverName" width="120" />

            <el-table-column align="left" label="服务器端口号" prop="serverPort" width="120" />

            <el-table-column align="left" label="服务器版本" prop="serverVersion" width="120" />

            <el-table-column align="left" label="服务器状态" prop="serverStatus" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.serverStatus,mcp_server_statusOptions) }}
    </template>
</el-table-column>
            <el-table-column align="left" label="是否自动分配端口" prop="autoAllocated" width="120">
    <template #default="scope">{{ formatBoolean(scope.row.autoAllocated) }}</template>
</el-table-column>
            <el-table-column sortable align="left" label="最后启动时间" prop="lastStartTime" width="180">
   <template #default="scope">{{ formatDate(scope.row.lastStartTime) }}</template>
</el-table-column>
            <el-table-column sortable align="left" label="当前连接数" prop="connectionCount" width="120" />

        <el-table-column align="left" label="操作" fixed="right" min-width="320">
            <template #default="scope">
            <el-button v-auth="btnAuth.info" type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button v-auth="btnAuth.edit" type="primary" link icon="edit" class="table-button" @click="updateMcpServerFunc(scope.row)">编辑</el-button>
            <el-button
              v-if="scope.row.serverStatus !== 'running'"
              type="success"
              link
              class="table-button"
              @click="startServer(scope.row)"
              :loading="scope.row.starting"
            >启动</el-button>
            <el-button
              v-if="scope.row.serverStatus === 'running'"
              type="warning"
              link
              class="table-button"
              @click="stopServer(scope.row)"
              :loading="scope.row.stopping"
            >停止</el-button>
            <el-button  v-auth="btnAuth.delete" type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-drawer destroy-on-close size="800" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'新增':'编辑'}}</span>
                <div>
                  <el-button :loading="btnLoading" type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
             <el-form-item label="MCP服务器名称:" prop="serverName">
    <el-input v-model="formData.serverName" :clearable="true" placeholder="请输入MCP服务器名称" />
</el-form-item>
             <el-form-item label="服务器端口号:" prop="serverPort">
    <el-input v-model.number="formData.serverPort" :clearable="true" placeholder="请输入服务器端口号" />
</el-form-item>
             <el-form-item label="服务器版本:" prop="serverVersion">
    <el-input v-model="formData.serverVersion" :clearable="true" placeholder="请输入服务器版本" />
</el-form-item>
             <el-form-item label="服务器状态:" prop="serverStatus">
    <el-select v-model="formData.serverStatus" placeholder="请选择服务器状态" style="width:100%" filterable :clearable="true">
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
    <el-input v-model="formData.description" :clearable="true" placeholder="请输入服务器描述" />
</el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                 <el-descriptions-item label="MCP服务器名称">
    {{ detailForm.serverName }}
</el-descriptions-item>
                 <el-descriptions-item label="服务器端口号">
    {{ detailForm.serverPort }}
</el-descriptions-item>
                 <el-descriptions-item label="服务器版本">
    {{ detailForm.serverVersion }}
</el-descriptions-item>
                 <el-descriptions-item label="服务器状态">
    {{ detailForm.serverStatus }}
</el-descriptions-item>
                 <el-descriptions-item label="是否自动分配端口">
    {{ detailForm.autoAllocated }}
</el-descriptions-item>
                 <el-descriptions-item label="服务器配置">
    {{ detailForm.serverConfig }}
</el-descriptions-item>
                 <el-descriptions-item label="服务器描述">
    {{ detailForm.description }}
</el-descriptions-item>
                 <el-descriptions-item label="最后启动时间">
    {{ detailForm.lastStartTime }}
</el-descriptions-item>
                 <el-descriptions-item label="当前连接数">
    {{ detailForm.connectionCount }}
</el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createMcpServer,
  deleteMcpServer,
  deleteMcpServerByIds,
  updateMcpServer,
  findMcpServer,
  getMcpServerList,
  startMcpServer,
  stopMcpServer
} from '@/plugin/zqmcp/api/mcp_server'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
// 引入按钮权限标识
import { useBtnAuth } from '@/utils/btnAuth'

// 导出组件
import ExportExcel from '@/components/exportExcel/exportExcel.vue'
// 导入组件
import ImportExcel from '@/components/exportExcel/importExcel.vue'
// 导出模板组件
import ExportTemplate from '@/components/exportExcel/exportTemplate.vue'


defineOptions({
    name: 'McpServer'
})
// 按钮权限实例化
    const btnAuth = useBtnAuth()

// 提交按钮loading
const btnLoading = ref(false)

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
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
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               serverPort : [{
                   required: true,
                   message: '请输入端口号',
                   trigger: ['input','blur'],
               },
              ],
               serverStatus : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
// 排序
const sortChange = ({ prop, order }) => {
  const sortMap = {
    CreatedAt:"created_at",
    ID:"id",
            lastStartTime: 'last_start_time',
            connectionCount: 'connection_count',
  }

  let sort = sortMap[prop]
  if(!sort){
   sort = prop.replace(/[A-Z]/g, match => `_${match.toLowerCase()}`)
  }

  searchInfo.value.sort = sort
  searchInfo.value.order = order
  getTableData()
}
// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    if (searchInfo.value.autoAllocated === ""){
        searchInfo.value.autoAllocated=null
    }
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getMcpServerList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
    mcp_server_statusOptions.value = await getDictFunc('mcp_server_status')
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteMcpServerFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const IDs = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          IDs.push(item.ID)
        })
      const res = await deleteMcpServerByIds({ IDs })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === IDs.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateMcpServerFunc = async(row) => {
    const res = await findMcpServer({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteMcpServerFunc = async (row) => {
    const res = await deleteMcpServer({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        serverName: '',
        serverPort: 0,
        serverVersion: '',
        serverStatus: '',
        autoAllocated: false,
        serverConfig: {},
        description: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
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
                closeDialog()
                getTableData()
              }
      })
}

const detailForm = ref({})

// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findMcpServer({ ID: row.ID })
  if (res.code === 0) {
    detailForm.value = res.data
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  detailForm.value = {}
}

// 启动MCP服务器
const startServer = async (row) => {
  try {
    row.starting = true
    const res = await startMcpServer(row.ID)
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '服务器启动成功'
      })
      getTableData()
    } else {
      ElMessage({
        type: 'error',
        message: res.msg || '服务器启动失败'
      })
    }
  } catch (error) {
    console.error('启动服务器失败:', error)
    ElMessage({
      type: 'error',
      message: '服务器启动失败'
    })
  } finally {
    row.starting = false
  }
}

// 停止MCP服务器
const stopServer = async (row) => {
  try {
    row.stopping = true
    const res = await stopMcpServer(row.ID)
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '服务器停止成功'
      })
      getTableData()
    } else {
      ElMessage({
        type: 'error',
        message: res.msg || '服务器停止失败'
      })
    }
  } catch (error) {
    console.error('停止服务器失败:', error)
    ElMessage({
      type: 'error',
      message: '服务器停止失败'
    })
  } finally {
    row.stopping = false
  }
}


</script>

<style>

</style>

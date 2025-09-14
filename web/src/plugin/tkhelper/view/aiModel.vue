<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="searchForm" :inline="true" :model="searchInfo">
        <el-form-item label="模型名称">
          <el-input v-model="searchInfo.name" placeholder="请输入模型名称" />
        </el-form-item>
        <el-form-item label="模型标识">
          <el-input v-model="searchInfo.modelName" placeholder="请输入模型标识" />
        </el-form-item>
        <el-form-item label="是否启用">
          <el-select v-model="searchInfo.enabled" placeholder="请选择">
            <el-option value="" label="全部" />
            <el-option :value="true" label="启用" />
            <el-option :value="false" label="禁用" />
          </el-select>
        </el-form-item>
        <el-form-item label="创建时间">
          <el-date-picker
            v-model="searchInfo.createdAt"
            type="datetimerange"
            range-separator="至"
            start-placeholder="开始时间"
            end-placeholder="结束时间"
          />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSearch">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>

    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog('add')">新增</el-button>
        <el-button
          type="danger"
          icon="delete"
          :disabled="!multipleSelection.length"
          @click="onBatchDelete"
        >
          批量删除
        </el-button>
      </div>

      <el-table
        ref="multipleTable"
        v-loading="loading"
        :data="tableData"
        style="width: 100%"
        tooltip-effect="dark"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column label="ID" prop="ID" width="80" />
        <el-table-column label="模型名称" prop="name" min-width="120" show-overflow-tooltip />
        <el-table-column label="模型标识" prop="modelName" min-width="150" show-overflow-tooltip />
        <el-table-column label="API地址" prop="baseUrl" min-width="200" show-overflow-tooltip>
          <template #default="scope">
            <el-tooltip :content="scope.row.baseUrl" placement="top">
              <span>{{ scope.row.baseUrl }}</span>
            </el-tooltip>
          </template>
        </el-table-column>
        <el-table-column label="是否启用" width="100" align="center">
          <template #default="scope">
            <el-switch
              v-model="scope.row.enabled"
              @change="handleStatusChange(scope.row)"
            />
          </template>
        </el-table-column>
        <el-table-column label="是否默认" width="100" align="center">
          <template #default="scope">
            <el-tag v-if="scope.row.isDefault" type="success">默认</el-tag>
            <el-button 
              v-else 
              type="primary" 
              link 
              size="small" 
              @click="setDefault(scope.row)"
              :disabled="!scope.row.enabled"
            >
              设为默认
            </el-button>
          </template>
        </el-table-column>
        <el-table-column label="排序" prop="sort" width="80" align="center" />
        <el-table-column label="创建时间" prop="CreatedAt" width="180">
          <template #default="scope">
            {{ formatDate(scope.row.CreatedAt) }}
          </template>
        </el-table-column>
        <el-table-column label="操作" fixed="right" min-width="200">
          <template #default="scope">
            <el-button type="primary" link icon="view" @click="openDialog('view', scope.row)">
              查看
            </el-button>
            <el-button type="primary" link icon="edit" @click="openDialog('edit', scope.row)">
              编辑
            </el-button>
            <el-button type="danger" link icon="delete" @click="deleteRow(scope.row)">
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>

      <div class="gva-pagination">
        <el-pagination
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          layout="total, sizes, prev, pager, next, jumper"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>

    <!-- 详情/编辑对话框 -->
    <el-dialog
      v-model="dialogVisible"
      :title="dialogTitle"
      width="600px"
      :close-on-click-modal="false"
    >
      <el-form
        ref="dialogForm"
        :model="formData"
        :rules="rules"
        label-width="100px"
        v-loading="dialogLoading"
      >
        <el-form-item label="模型名称" prop="name">
          <el-input 
            v-model="formData.name" 
            placeholder="请输入模型显示名称"
            :disabled="dialogType === 'view'"
          />
        </el-form-item>
        <el-form-item label="API地址" prop="baseUrl">
          <el-input 
            v-model="formData.baseUrl" 
            placeholder="请输入API基础URL，如：https://api.openai.com/v1"
            :disabled="dialogType === 'view'"
          />
        </el-form-item>
        <el-form-item label="API密钥" prop="apiKey">
          <el-input 
            v-model="formData.apiKey" 
            placeholder="请输入API密钥"
            type="password"
            show-password
            :disabled="dialogType === 'view'"
          />
        </el-form-item>
        <el-form-item label="模型标识" prop="modelName">
          <el-input 
            v-model="formData.modelName" 
            placeholder="请输入模型标识，如：gpt-3.5-turbo"
            :disabled="dialogType === 'view'"
          />
        </el-form-item>
        <el-form-item label="是否启用" prop="enabled">
          <el-switch 
            v-model="formData.enabled" 
            :disabled="dialogType === 'view'"
          />
        </el-form-item>
        <el-form-item label="是否默认" prop="isDefault" v-if="dialogType !== 'add'">
          <el-switch 
            v-model="formData.isDefault" 
            :disabled="dialogType === 'view' || !formData.enabled"
          />
        </el-form-item>
        <el-form-item label="排序" prop="sort">
          <el-input-number 
            v-model="formData.sort" 
            :min="0" 
            :max="9999" 
            placeholder="数字越小越靠前"
            :disabled="dialogType === 'view'"
          />
        </el-form-item>
        <el-form-item label="备注" prop="remark">
          <el-input 
            v-model="formData.remark" 
            type="textarea" 
            :rows="3" 
            placeholder="请输入备注信息"
            :disabled="dialogType === 'view'"
          />
        </el-form-item>
      </el-form>

      <template #footer>
        <div class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button 
            v-if="dialogType !== 'view'" 
            type="primary" 
            @click="submitForm"
            :loading="submitLoading"
          >
            确定
          </el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
import {
  getAiModelList,
  createAiModel,
  updateAiModel,
  deleteAiModel,
  deleteAiModelByIds,
  findAiModel,
  setDefaultAiModel
} from '@/plugin/tkhelper/api/aiModel'
import { formatTimeToStr } from '@/utils/date'
import { ElMessage, ElMessageBox } from 'element-plus'

export default {
  name: 'AiModel',
  data() {
    return {
      page: 1,
      pageSize: 10,
      total: 0,
      loading: false,
      tableData: [],
      multipleSelection: [],
      searchInfo: {
        name: '',
        modelName: '',
        enabled: '',
        createdAt: []
      },
      dialogVisible: false,
      dialogLoading: false,
      submitLoading: false,
      dialogType: 'add',
      dialogTitle: '新增AI模型',
      formData: this.initFormData(),
      rules: {
        name: [
          { required: true, message: '请输入模型名称', trigger: 'blur' }
        ],
        baseUrl: [
          { required: true, message: '请输入API基础URL', trigger: 'blur' },
          { type: 'url', message: '请输入有效的URL地址', trigger: 'blur' }
        ],
        apiKey: [
          { required: true, message: '请输入API密钥', trigger: 'blur' }
        ],
        modelName: [
          { required: true, message: '请输入模型标识', trigger: 'blur' }
        ]
      }
    }
  },
  created() {
    this.getTableData()
  },
  methods: {
    initFormData() {
      return {
        ID: 0,
        name: '',
        baseUrl: '',
        apiKey: '',
        modelName: '',
        enabled: true,
        isDefault: false,
        sort: 999,
        remark: ''
      }
    },

    async getTableData() {
      this.loading = true
      const params = {
        page: this.page,
        pageSize: this.pageSize
      }

      // 只添加非空的搜索条件
      if (this.searchInfo.name && this.searchInfo.name.trim()) {
        params.name = this.searchInfo.name.trim()
      }
      if (this.searchInfo.modelName && this.searchInfo.modelName.trim()) {
        params.modelName = this.searchInfo.modelName.trim()
      }
      if (this.searchInfo.enabled !== '' && this.searchInfo.enabled !== null && this.searchInfo.enabled !== undefined) {
        params.enabled = this.searchInfo.enabled
      }

      if (this.searchInfo.createdAt && this.searchInfo.createdAt.length === 2) {
        params.startCreatedAt = this.searchInfo.createdAt[0]
        params.endCreatedAt = this.searchInfo.createdAt[1]
      }

      try {
        const res = await getAiModelList(params)
        if (res.code === 0) {
          this.tableData = res.data.list || []
          this.total = res.data.total || 0
        }
      } catch (error) {
        ElMessage.error('获取数据失败')
      } finally {
        this.loading = false
      }
    },

    onSearch() {
      this.page = 1
      this.getTableData()
    },

    onReset() {
      this.searchInfo = {
        name: '',
        modelName: '',
        enabled: '',
        createdAt: []
      }
      this.onSearch()
    },

    handleCurrentChange(val) {
      this.page = val
      this.getTableData()
    },

    handleSizeChange(val) {
      this.pageSize = val
      this.getTableData()
    },

    handleSelectionChange(val) {
      this.multipleSelection = val
    },

    async handleStatusChange(row) {
      try {
        await updateAiModel(row)
        ElMessage.success('状态更新成功')
        this.getTableData()
      } catch (error) {
        ElMessage.error('状态更新失败')
        // 恢复原状态
        row.enabled = !row.enabled
      }
    },

    async setDefault(row) {
      try {
        await setDefaultAiModel({ ID: row.ID })
        ElMessage.success('设置默认模型成功')
        this.getTableData()
      } catch (error) {
        ElMessage.error('设置失败')
      }
    },

    openDialog(type, row = null) {
      this.dialogType = type
      this.dialogTitle = {
        add: '新增AI模型',
        edit: '编辑AI模型',
        view: '查看AI模型详情'
      }[type]

      if (row) {
        this.formData = { ...row }
      } else {
        this.formData = this.initFormData()
      }
      
      this.dialogVisible = true
    },

    async submitForm() {
      if (this.dialogType === 'view') return

      this.$refs.dialogForm.validate(async (valid) => {
        if (valid) {
          this.submitLoading = true
          try {
            if (this.dialogType === 'add') {
              await createAiModel(this.formData)
              ElMessage.success('创建成功')
            } else {
              await updateAiModel(this.formData)
              ElMessage.success('更新成功')
            }
            this.dialogVisible = false
            this.getTableData()
          } catch (error) {
            ElMessage.error(`${this.dialogType === 'add' ? '创建' : '更新'}失败`)
          } finally {
            this.submitLoading = false
          }
        }
      })
    },

    deleteRow(row) {
      ElMessageBox.confirm('此操作将永久删除该模型配置, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(async () => {
        try {
          await deleteAiModel({ ID: row.ID })
          ElMessage.success('删除成功')
          this.getTableData()
        } catch (error) {
          ElMessage.error('删除失败')
        }
      })
    },

    onBatchDelete() {
      const ids = this.multipleSelection.map(item => item.ID.toString())
      ElMessageBox.confirm('此操作将永久删除所选模型配置, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(async () => {
        try {
          await deleteAiModelByIds({ 'IDs[]': ids })
          ElMessage.success('批量删除成功')
          this.getTableData()
        } catch (error) {
          ElMessage.error('批量删除失败')
        }
      })
    },

    formatDate(time) {
      if (time != null && time !== '') {
        return formatTimeToStr(time, 'yyyy-MM-dd hh:mm:ss')
      } else {
        return ''
      }
    }
  }
}
</script>

<style scoped>
.gva-search-box {
  padding: 20px;
  background: white;
  border-radius: 6px;
  margin-bottom: 20px;
}

.gva-table-box {
  background: white;
  border-radius: 6px;
}

.gva-btn-list {
  padding: 20px;
  padding-bottom: 0;
}

.gva-pagination {
  padding: 20px;
  text-align: right;
}

.dialog-footer {
  text-align: right;
}

.el-table .el-table__cell {
  padding: 8px 0;
}
</style>
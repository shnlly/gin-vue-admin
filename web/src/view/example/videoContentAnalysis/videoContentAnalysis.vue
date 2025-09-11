<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="searchForm" :inline="true" :model="searchInfo">
        <el-form-item label="达人名称">
          <el-input v-model="searchInfo.creatorName" placeholder="请输入达人名称" />
        </el-form-item>
        <el-form-item label="视频标题">
          <el-input v-model="searchInfo.videoTitle" placeholder="请输入视频标题" />
        </el-form-item>
        <el-form-item label="是否下载">
          <el-select v-model="searchInfo.isDownloaded" placeholder="请选择">
            <el-option value="" label="全部" />
            <el-option :value="true" label="已下载" />
            <el-option :value="false" label="未下载" />
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
          type="success" 
          icon="upload" 
          @click="importDialog = true"
        >
          导入Excel
        </el-button>
        <el-button 
          type="info" 
          icon="pie-chart" 
          @click="statisticsDialog = true"
        >
          统计信息
        </el-button>
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
        <el-table-column label="视频标题" prop="videoTitle" min-width="200" show-overflow-tooltip>
          <template #default="scope">
            <el-link type="primary" @click="openDialog('view', scope.row)">
              {{ scope.row.videoTitle || '暂无标题' }}
            </el-link>
          </template>
        </el-table-column>
        <el-table-column label="达人名称" prop="creatorName" width="120" />
        <el-table-column label="播放量" prop="playCount" width="100" />
        <el-table-column label="点赞数" prop="likeCount" width="100" />
        <el-table-column label="销量" prop="salesVolume" width="80" />
        <el-table-column label="销售额" prop="salesAmount" width="100" />
        <el-table-column label="发布时间" prop="publishTime" width="120" />
        <el-table-column label="是否下载" width="100">
          <template #default="scope">
            <el-tag :type="scope.row.isDownloaded ? 'success' : 'info'">
              {{ scope.row.isDownloaded ? '已下载' : '未下载' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="创建时间" prop="CreatedAt" width="180">
          <template #default="scope">
            {{ formatDate(scope.row.CreatedAt) }}
          </template>
        </el-table-column>
        <el-table-column label="操作" fixed="right" min-width="200">
          <template #default="scope">
            <el-button type="primary" link icon="view" @click="openDialog('view', scope.row)">
              查看详情
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
      width="80%"
      :close-on-click-modal="false"
    >
      <el-form
        ref="dialogForm"
        :model="formData"
        :rules="rules"
        label-width="120px"
        :disabled="dialogType === 'view'"
      >
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="视频标题" prop="videoTitle">
              <el-input v-model="formData.videoTitle" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="达人名称" prop="creatorName">
              <el-input v-model="formData.creatorName" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="视频链接" prop="videoLink">
              <el-input v-model="formData.videoLink" type="textarea" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="唯一标识" prop="uniqueId">
              <el-input v-model="formData.uniqueId" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="粉丝数" prop="fansCount">
              <el-input v-model="formData.fansCount" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="播放量" prop="playCount">
              <el-input v-model="formData.playCount" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="点赞数" prop="likeCount">
              <el-input v-model="formData.likeCount" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="评论数" prop="commentCount">
              <el-input-number v-model="formData.commentCount" :min="0" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="转发数" prop="shareCount">
              <el-input-number v-model="formData.shareCount" :min="0" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="销量(件)" prop="salesVolume">
              <el-input-number v-model="formData.salesVolume" :min="0" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="销售额" prop="salesAmount">
              <el-input v-model="formData.salesAmount" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="发布时间" prop="publishTime">
              <el-input v-model="formData.publishTime" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="视频时长" prop="videoDuration">
              <el-input v-model="formData.videoDuration" />
            </el-form-item>
          </el-col>
          <el-col :span="24">
            <el-form-item label="视频描述" prop="videoDescription">
              <el-input v-model="formData.videoDescription" type="textarea" :rows="3" />
            </el-form-item>
          </el-col>
          <el-col :span="24">
            <el-form-item label="视频文案" prop="videoScript">
              <el-input v-model="formData.videoScript" type="textarea" :rows="4" />
            </el-form-item>
          </el-col>
          <el-col :span="24">
            <el-form-item label="修复后文案" prop="fixedScript">
              <el-input v-model="formData.fixedScript" type="textarea" :rows="4" />
            </el-form-item>
          </el-col>
          <el-col :span="24">
            <el-form-item label="视频画面分析" prop="videoAnalysis">
              <el-input v-model="formData.videoAnalysis" type="textarea" :rows="4" />
            </el-form-item>
          </el-col>
          <el-col :span="24">
            <el-form-item label="修复说明" prop="fixNote">
              <el-input v-model="formData.fixNote" type="textarea" :rows="2" />
            </el-form-item>
          </el-col>
          <el-col :span="24">
            <el-form-item label="处理状态">
              <el-row :gutter="10">
                <el-col :span="4">
                  <el-checkbox v-model="formData.isDownloaded">已下载</el-checkbox>
                </el-col>
                <el-col :span="4">
                  <el-checkbox v-model="formData.isAudioExtracted">已提取音频</el-checkbox>
                </el-col>
                <el-col :span="4">
                  <el-checkbox v-model="formData.isTextConverted">已转文本</el-checkbox>
                </el-col>
                <el-col :span="4">
                  <el-checkbox v-model="formData.isScriptFixed">已修复文案</el-checkbox>
                </el-col>
                <el-col :span="4">
                  <el-checkbox v-model="formData.isVideoAnalyzed">已分析画面</el-checkbox>
                </el-col>
              </el-row>
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button 
            v-if="dialogType !== 'view'" 
            type="primary" 
            @click="submitForm"
          >
            确定
          </el-button>
        </div>
      </template>
    </el-dialog>

    <!-- 导入对话框 -->
    <el-dialog v-model="importDialog" title="导入Excel" width="500px">
      <el-upload
        class="upload-demo"
        drag
        action=""
        :before-upload="beforeUpload"
        :http-request="handleUpload"
        :show-file-list="false"
      >
        <el-icon class="el-icon--upload">
          <upload-filled />
        </el-icon>
        <div class="el-upload__text">
          将Excel文件拖到此处，或<em>点击上传</em>
        </div>
        <div class="el-upload__tip text-red">
          只能上传xlsx/xls文件，且不超过10MB
        </div>
      </el-upload>
    </el-dialog>

    <!-- 统计信息对话框 -->
    <el-dialog v-model="statisticsDialog" title="统计信息" width="800px">
      <div v-loading="statisticsLoading">
        <el-row :gutter="20">
          <el-col :span="6">
            <el-statistic title="总视频数" :value="statistics.totalCount || 0" />
          </el-col>
          <el-col :span="6">
            <el-statistic title="已下载数" :value="statistics.downloadedCount || 0" />
          </el-col>
          <el-col :span="6">
            <el-statistic title="已分析数" :value="statistics.analyzedCount || 0" />
          </el-col>
          <el-col :span="6">
            <el-statistic title="总播放量" :value="statistics.totalPlayCount || 0" />
          </el-col>
        </el-row>
        <el-divider>头部达人排行</el-divider>
        <el-table :data="statistics.topCreators || []" style="width: 100%">
          <el-table-column label="达人名称" prop="creatorName" />
          <el-table-column label="视频数量" prop="videoCount" />
          <el-table-column label="总播放量" prop="totalPlay" />
        </el-table>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
  getVideoContentAnalysisList,
  createVideoContentAnalysis,
  updateVideoContentAnalysis,
  deleteVideoContentAnalysis,
  deleteVideoContentAnalysisByIds,
  findVideoContentAnalysis,
  getVideoContentAnalysisStatistics,
  importVideoContentAnalysisFromExcel
} from '@/api/example/videoContentAnalysis'
import { formatTimeToStr } from '@/utils/date'
import { ElMessage, ElMessageBox } from 'element-plus'
import service from '@/utils/request'

export default {
  name: 'VideoContentAnalysis',
  data() {
    return {
      page: 1,
      pageSize: 10,
      total: 0,
      loading: false,
      tableData: [],
      multipleSelection: [],
      searchInfo: {
        creatorName: '',
        videoTitle: '',
        isDownloaded: '',
        createdAt: []
      },
      dialogVisible: false,
      dialogType: 'add',
      dialogTitle: '新增视频内容',
      formData: this.initFormData(),
      rules: {
        videoLink: [
          { required: true, message: '请输入视频链接', trigger: 'blur' }
        ],
        creatorName: [
          { required: true, message: '请输入达人名称', trigger: 'blur' }
        ]
      },
      importDialog: false,
      statisticsDialog: false,
      statisticsLoading: false,
      statistics: {}
    }
  },
  created() {
    this.getTableData()
  },
  methods: {
    initFormData() {
      return {
        ID: 0,
        videoDescription: '',
        videoLink: '',
        creatorName: '',
        uniqueId: '',
        fansCount: '',
        playCount: '',
        likeCount: '',
        commentCount: 0,
        shareCount: 0,
        salesVolume: 0,
        salesAmount: '',
        publishTime: '',
        videoDuration: '',
        videoTitle: '',
        videoScript: '',
        fixedScript: '',
        videoAnalysis: '',
        isDownloaded: false,
        isAudioExtracted: false,
        isTextConverted: false,
        isScriptFixed: false,
        isVideoAnalyzed: false,
        fixNote: ''
      }
    },
    
    async getTableData() {
      this.loading = true
      const params = {
        page: this.page,
        pageSize: this.pageSize
      }
      
      // 只添加非空的搜索条件
      if (this.searchInfo.creatorName && this.searchInfo.creatorName.trim()) {
        params.creatorName = this.searchInfo.creatorName.trim()
      }
      if (this.searchInfo.videoTitle && this.searchInfo.videoTitle.trim()) {
        params.videoTitle = this.searchInfo.videoTitle.trim()
      }
      if (this.searchInfo.isDownloaded !== '' && this.searchInfo.isDownloaded !== null && this.searchInfo.isDownloaded !== undefined) {
        params.isDownloaded = this.searchInfo.isDownloaded
      }
      
      if (this.searchInfo.createdAt && this.searchInfo.createdAt.length === 2) {
        params.startCreatedAt = this.searchInfo.createdAt[0]
        params.endCreatedAt = this.searchInfo.createdAt[1]
      }
      
      try {
        const res = await getVideoContentAnalysisList(params)
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
        creatorName: '',
        videoTitle: '',
        isDownloaded: '',
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

    openDialog(type, row = null) {
      this.dialogType = type
      this.dialogTitle = {
        add: '新增视频内容',
        edit: '编辑视频内容',
        view: '查看视频内容详情'
      }[type]

      if (row) {
        this.formData = { ...row }
      } else {
        this.formData = this.initFormData()
      }
      
      this.dialogVisible = true
    },

    async submitForm() {
      this.$refs.dialogForm.validate(async (valid) => {
        if (valid) {
          try {
            if (this.dialogType === 'add') {
              await createVideoContentAnalysis(this.formData)
              ElMessage.success('创建成功')
            } else {
              await updateVideoContentAnalysis(this.formData)
              ElMessage.success('更新成功')
            }
            this.dialogVisible = false
            this.getTableData()
          } catch (error) {
            ElMessage.error(`${this.dialogType === 'add' ? '创建' : '更新'}失败`)
          }
        }
      })
    },

    deleteRow(row) {
      ElMessageBox.confirm('此操作将永久删除该数据, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(async () => {
        try {
          await deleteVideoContentAnalysis(row.ID)
          ElMessage.success('删除成功')
          this.getTableData()
        } catch (error) {
          ElMessage.error('删除失败')
        }
      })
    },

    onBatchDelete() {
      const ids = this.multipleSelection.map(item => item.ID.toString())
      ElMessageBox.confirm('此操作将永久删除所选数据, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(async () => {
        try {
          await deleteVideoContentAnalysisByIds(ids)
          ElMessage.success('批量删除成功')
          this.getTableData()
        } catch (error) {
          ElMessage.error('批量删除失败')
        }
      })
    },

    async getStatistics() {
      this.statisticsLoading = true
      try {
        const res = await getVideoContentAnalysisStatistics()
        if (res.code === 0) {
          this.statistics = res.data
        }
      } catch (error) {
        ElMessage.error('获取统计信息失败')
      } finally {
        this.statisticsLoading = false
      }
    },

    beforeUpload(file) {
      const isExcel = file.type === 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet' 
        || file.type === 'application/vnd.ms-excel'
      const isLt10M = file.size / 1024 / 1024 < 10

      if (!isExcel) {
        ElMessage.error('只能上传Excel文件!')
        return false
      }
      if (!isLt10M) {
        ElMessage.error('上传文件大小不能超过 10MB!')
        return false
      }
      return true
    },

    async handleUpload(params) {
      try {
        const file = params.file
        
        // 创建 FormData 对象
        const formData = new FormData()
        formData.append('file', file)
        
        // 首先上传文件到服务器
        const uploadResponse = await service({
          url: '/fileUploadAndDownload/upload',
          method: 'post',
          data: formData,
          headers: {
            'Content-Type': 'multipart/form-data'
          }
        })
        
        if (uploadResponse.code === 0) {
          // 获取上传后的文件路径
          const filePath = uploadResponse.data.file.url
          
          // 调用导入接口
          const importResponse = await importVideoContentAnalysisFromExcel({
            filePath: filePath
          })
          
          if (importResponse.code === 0) {
            const { successCount, failCount } = importResponse.data
            ElMessage.success(`导入完成！成功：${successCount}条，失败：${failCount}条`)
            this.getTableData() // 刷新数据
          } else {
            ElMessage.error('导入失败：' + importResponse.msg)
          }
        } else {
          ElMessage.error('文件上传失败：' + uploadResponse.msg)
        }
        
        this.importDialog = false
      } catch (error) {
        console.error('导入失败:', error)
        ElMessage.error('导入失败，请检查文件格式和网络连接')
        this.importDialog = false
      }
    },

    formatDate(time) {
      if (time != null && time !== '') {
        return formatTimeToStr(time, 'yyyy-MM-dd hh:mm:ss')
      } else {
        return ''
      }
    }
  },

  watch: {
    statisticsDialog(val) {
      if (val) {
        this.getStatistics()
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
</style>
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
      :width="dialogType === 'view' ? '90%' : '80%'"
      :close-on-click-modal="false"
      :class="dialogType === 'view' ? 'video-detail-dialog enhanced-view' : 'video-detail-dialog'"
    >
      <!-- 查看模式：左右两列布局 -->
      <div v-if="dialogType === 'view'" class="video-detail-view">
        <el-row :gutter="24" class="detail-layout">
          <!-- 左侧：基础信息（约1/5宽度） -->
          <el-col :span="5" class="left-panel">
            <!-- 头部信息卡片 -->
            <el-card class="detail-card header-card" shadow="hover">
              <template #header>
                <div class="card-header">
                  <span class="card-title">📹 基本信息</span>
                  <div class="header-tags">
                    <el-tag v-if="formData.isDownloaded" type="success" size="small">已下载</el-tag>
                    <el-tag v-else type="info" size="small">未下载</el-tag>
                  </div>
                </div>
              </template>
              <div class="basic-info">
                <div class="info-item">
                  <div class="info-label">🎬 标题</div>
                  <div class="info-value title-value">{{ formData.videoTitle || '暂无标题' }}</div>
                </div>
                <div class="info-item">
                  <div class="info-label">👤 达人</div>
                  <div class="info-value creator-name">{{ formData.creatorName }}</div>
                </div>
                <div class="info-item">
                  <div class="info-label">🆔 ID</div>
                  <div class="info-value">{{ formData.uniqueId || '暂无' }}</div>
                </div>
                <div class="info-item">
                  <div class="info-label">🔗 链接</div>
                  <div class="info-value link-value">
                    <el-link :href="formData.videoLink" target="_blank" type="primary" :underline="false">
                      查看视频
                    </el-link>
                  </div>
                </div>
              </div>
            </el-card>

            <!-- 数据统计卡片 -->
            <el-card class="detail-card stats-card" shadow="hover">
              <template #header>
                <div class="card-header">
                  <span class="card-title">📊 数据</span>
                  <el-button type="primary" icon="TrendCharts" size="small" @click="analyzeVideoData">
                    分析
                  </el-button>
                </div>
              </template>
              <div class="compact-stats">
                <div class="stat-row">
                  <span class="stat-label">👥 粉丝</span>
                  <span class="stat-value">{{ formData.fansCount || '0' }}</span>
                </div>
                <div class="stat-row highlight-row">
                  <span class="stat-label">▶️ 播放</span>
                  <span class="stat-value highlight">{{ formData.playCount || '0' }}</span>
                </div>
                <div class="stat-row">
                  <span class="stat-label">👍 点赞</span>
                  <span class="stat-value">{{ formData.likeCount || '0' }}</span>
                </div>
                <div class="stat-row money-row">
                  <span class="stat-label">💰 销售额</span>
                  <span class="stat-value money-value">{{ formData.salesAmount || '0' }}</span>
                </div>
                <div class="stat-row">
                  <span class="stat-label">💬 评论</span>
                  <span class="stat-value">{{ formData.commentCount || 0 }}</span>
                </div>
                <div class="stat-row">
                  <span class="stat-label">📦 销量</span>
                  <span class="stat-value">{{ formData.salesVolume || 0 }}件</span>
                </div>
              </div>
            </el-card>

            <!-- 处理状态卡片 -->
            <el-card class="detail-card status-card" shadow="hover">
              <template #header>
                <div class="card-header">
                  <span class="card-title">⚙️ 状态</span>
                </div>
              </template>
              <div class="compact-status">
                <div class="status-row" :class="{ active: formData.isDownloaded }">
                  <span class="status-icon">⬇️</span>
                  <span class="status-text">下载</span>
                  <el-tag :type="formData.isDownloaded ? 'success' : 'info'" size="small">
                    {{ formData.isDownloaded ? '✓' : '○' }}
                  </el-tag>
                </div>
                <div class="status-row" :class="{ active: formData.isAudioExtracted }">
                  <span class="status-icon">🎵</span>
                  <span class="status-text">音频</span>
                  <el-tag :type="formData.isAudioExtracted ? 'success' : 'info'" size="small">
                    {{ formData.isAudioExtracted ? '✓' : '○' }}
                  </el-tag>
                </div>
                <div class="status-row" :class="{ active: formData.isTextConverted }">
                  <span class="status-icon">📝</span>
                  <span class="status-text">文本</span>
                  <el-tag :type="formData.isTextConverted ? 'success' : 'info'" size="small">
                    {{ formData.isTextConverted ? '✓' : '○' }}
                  </el-tag>
                </div>
                <div class="status-row" :class="{ active: formData.isScriptFixed }">
                  <span class="status-icon">✨</span>
                  <span class="status-text">修复</span>
                  <el-tag :type="formData.isScriptFixed ? 'success' : 'info'" size="small">
                    {{ formData.isScriptFixed ? '✓' : '○' }}
                  </el-tag>
                </div>
                <div class="status-row" :class="{ active: formData.isVideoAnalyzed }">
                  <span class="status-icon">🔍</span>
                  <span class="status-text">分析</span>
                  <el-tag :type="formData.isVideoAnalyzed ? 'success' : 'info'" size="small">
                    {{ formData.isVideoAnalyzed ? '✓' : '○' }}
                  </el-tag>
                </div>
              </div>
            </el-card>

            <!-- 时间信息 -->
            <div class="time-info compact">
              <div class="time-item">📅 {{ formData.publishTime || '未知' }}</div>
              <div class="time-item">🕒 {{ formatDate(formData.CreatedAt) }}</div>
            </div>
          </el-col>

          <!-- 右侧：画面分析（约4/5宽度） -->
          <el-col :span="19" class="right-panel">
            <!-- 画面分析组件 -->
            <div class="video-analysis-panel">
              <el-card class="analysis-header-card" shadow="never">
                <template #header>
                  <div class="analysis-header">
                    <div class="header-info">
                      <span class="analysis-title">🎬 视频画面分析</span>
                      <span class="analysis-subtitle" v-if="parsedAnalysisData">{{ parsedAnalysisData.Summary?.one_line || '智能分析结果' }}</span>
                    </div>
                    <div class="header-actions">
                      <el-button type="success" icon="MagicStick" size="small" @click="fixVideoScript">
                        修复文案
                      </el-button>
                      <el-button type="warning" icon="VideoCamera" size="small" @click="analyzeVideoFrame">
                        重新分析
                      </el-button>
                      <el-button type="primary" icon="Refresh" size="small" @click="refreshAnalysisData">
                        刷新数据
                      </el-button>
                    </div>
                  </div>
                </template>
                <div v-if="!parsedAnalysisData" class="no-analysis">
                  <el-empty description="暂无画面分析数据">
                    <el-button type="primary" @click="analyzeVideoFrame">开始分析</el-button>
                  </el-empty>
                </div>
              </el-card>

              <!-- 分析内容 -->
              <div v-if="parsedAnalysisData" class="analysis-content">
                <!-- 分镜时间轴 -->
                <el-card class="analysis-section timeline-section" shadow="hover">
                  <template #header>
                    <div class="section-header">
                      <span class="section-title">🎞️ 分镜时间轴</span>
                      <el-tag type="info" size="small">{{ parsedAnalysisData.Raw?.length || 0 }} 个镜头</el-tag>
                    </div>
                  </template>
                  <div class="timeline-container">
                    <div 
                      v-for="(shot, index) in parsedAnalysisData.Raw" 
                      :key="index" 
                      class="timeline-item"
                      @click="activeTimelineIndex = index"
                      :class="{ active: activeTimelineIndex === index }"
                    >
                      <div class="timeline-time">{{ shot.time }}</div>
                      <div class="timeline-content">
                        <div class="scene-info">
                          <div class="scene-title">{{ shot.scene || '场景描述' }}</div>
                          <div class="scene-actions">{{ shot.actions || '动作描述' }}</div>
                        </div>
                        <div class="scene-details">
                          <el-tag v-if="shot.products" type="warning" size="small">🛍️ {{ shot.products.slice(0, 30) }}...</el-tag>
                          <el-tag v-if="shot.emotion" :type="getEmotionType(shot.emotion)" size="small">😊 {{ shot.emotion }}</el-tag>
                        </div>
                        <div v-if="shot.audio_text" class="audio-text">
                          <span class="audio-icon">🎤</span>
                          "{{ shot.audio_text }}"
                        </div>
                      </div>
                    </div>
                  </div>
                </el-card>

                <!-- 重点时刻分析 -->
                <el-card v-if="parsedAnalysisData.Highlights" class="analysis-section highlights-section" shadow="hover">
                  <template #header>
                    <div class="section-header">
                      <span class="section-title">⭐ 重点时刻</span>
                      <el-tag type="danger" size="small">{{ parsedAnalysisData.Highlights.length }} 个爆点</el-tag>
                    </div>
                  </template>
                  <div class="highlights-grid">
                    <div 
                      v-for="(highlight, index) in parsedAnalysisData.Highlights" 
                      :key="index" 
                      class="highlight-item"
                    >
                      <div class="highlight-header">
                        <el-tag :type="getCategoryType(highlight.category)" size="small">
                          {{ highlight.category }}
                        </el-tag>
                        <span class="highlight-time">{{ highlight.time }}</span>
                      </div>
                      <div class="highlight-label">{{ highlight.label }}</div>
                      <div class="highlight-description">{{ highlight.description }}</div>
                      <div class="highlight-effect">
                        <span class="effect-label">预期效果：</span>
                        <span class="effect-text">{{ highlight.effect }}</span>
                      </div>
                      <div class="highlight-drivers">
                        <el-tag 
                          v-for="driver in highlight.driver" 
                          :key="driver" 
                          :type="getDriverType(driver)" 
                          size="small"
                          class="driver-tag"
                        >
                          {{ driver }}
                        </el-tag>
                      </div>
                      <div class="highlight-metrics">
                        <el-tag :type="highlight.replicability === '高' ? 'success' : highlight.replicability === '中' ? 'warning' : 'info'" size="small">
                          可复制性: {{ highlight.replicability }}
                        </el-tag>
                        <el-tag :type="highlight.risk === '无' ? 'success' : 'danger'" size="small">
                          风险: {{ highlight.risk }}
                        </el-tag>
                      </div>
                    </div>
                  </div>
                </el-card>

                <!-- 视频总结 -->
                <el-row :gutter="16" class="summary-row">
                  <el-col :span="12">
                    <el-card class="analysis-section summary-section" shadow="hover">
                      <template #header>
                        <span class="section-title">📋 视频总结</span>
                      </template>
                      <div class="summary-content">
                        <div class="summary-item">
                          <div class="summary-label">🎯 核心观点</div>
                          <div class="summary-value">{{ parsedAnalysisData.Summary?.one_line || '暂无' }}</div>
                        </div>
                        <div class="summary-item">
                          <div class="summary-label">🔥 关键亮点</div>
                          <div class="summary-value">{{ parsedAnalysisData.Summary?.highlight_overview || '暂无' }}</div>
                        </div>
                        <div class="summary-item">
                          <div class="summary-label">🎣 引流钩子</div>
                          <div class="summary-value">{{ parsedAnalysisData.Summary?.hook || '暂无' }}</div>
                        </div>
                        <div class="summary-item">
                          <div class="summary-label">📞 行动召唤</div>
                          <div class="summary-value">{{ parsedAnalysisData.Summary?.cta || '暂无' }}</div>
                        </div>
                        <div v-if="parsedAnalysisData.Summary?.key_visuals" class="summary-item">
                          <div class="summary-label">👁️ 关键视觉</div>
                          <div class="visual-tags">
                            <el-tag 
                              v-for="visual in parsedAnalysisData.Summary.key_visuals" 
                              :key="visual" 
                              type="primary" 
                              size="small"
                              class="visual-tag"
                            >
                              {{ visual }}
                            </el-tag>
                          </div>
                        </div>
                      </div>
                    </el-card>
                  </el-col>
                  <el-col :span="12">
                    <el-card class="analysis-section product-section" shadow="hover">
                      <template #header>
                        <span class="section-title">🛍️ 产品信息</span>
                      </template>
                      <div v-if="parsedAnalysisData.ProductInfo" class="product-content">
                        <div class="product-item">
                          <div class="product-label">📦 产品名称</div>
                          <div class="product-value">{{ parsedAnalysisData.ProductInfo.name || '暂无' }}</div>
                        </div>
                        <div class="product-item">
                          <div class="product-label">✨ 卖点</div>
                          <div class="product-value">{{ parsedAnalysisData.ProductInfo.selling_points || '暂无' }}</div>
                        </div>
                        <div class="product-item">
                          <div class="product-label">🎬 展示方式</div>
                          <div class="product-value">{{ parsedAnalysisData.ProductInfo.presentation_methods || '暂无' }}</div>
                        </div>
                      </div>
                      <div v-else class="no-product-info">
                        <el-empty description="暂无产品信息" :image-size="60" />
                      </div>
                    </el-card>
                  </el-col>
                </el-row>

                <!-- 结构分析和痛点分析 -->
                <el-row :gutter="16" class="analysis-bottom-row">
                  <el-col :span="8">
                    <el-card class="analysis-section structure-section" shadow="hover">
                      <template #header>
                        <span class="section-title">🏗️ 结构分析</span>
                      </template>
                      <div v-if="parsedAnalysisData.Structure" class="structure-content">
                        <div class="structure-item">
                          <div class="structure-label">📝 类型</div>
                          <div class="structure-value">{{ parsedAnalysisData.Structure.type || '暂无' }}</div>
                        </div>
                        <div class="structure-item">
                          <div class="structure-label">⏱️ 最佳留存</div>
                          <div class="structure-value">{{ parsedAnalysisData.Structure.best_for_retention || '暂无' }}</div>
                        </div>
                        <div class="structure-item">
                          <div class="structure-label">💬 最佳互动</div>
                          <div class="structure-value">{{ parsedAnalysisData.Structure.best_for_interaction || '暂无' }}</div>
                        </div>
                        <div class="structure-item">
                          <div class="structure-label">💰 最佳转化</div>
                          <div class="structure-value">{{ parsedAnalysisData.Structure.best_for_conversion || '暂无' }}</div>
                        </div>
                      </div>
                    </el-card>
                  </el-col>
                  <el-col :span="8">
                    <el-card class="analysis-section painpoints-section" shadow="hover">
                      <template #header>
                        <span class="section-title">🎯 痛点分析</span>
                      </template>
                      <div v-if="parsedAnalysisData.PainPoints" class="painpoints-content">
                        <div class="painpoint-item">
                          <div class="painpoint-label">👥 目标用户</div>
                          <div class="painpoint-value">{{ parsedAnalysisData.PainPoints.target_audience || '暂无' }}</div>
                        </div>
                        <div class="painpoint-item">
                          <div class="painpoint-label">💡 用户需求</div>
                          <div class="painpoint-value">{{ parsedAnalysisData.PainPoints.user_needs || '暂无' }}</div>
                        </div>
                        <div class="painpoint-item">
                          <div class="painpoint-label">🔔 共鸣触发</div>
                          <div class="painpoint-value">{{ parsedAnalysisData.PainPoints.resonance_triggers || '暂无' }}</div>
                        </div>
                      </div>
                    </el-card>
                  </el-col>
                  <el-col :span="8">
                    <el-card class="analysis-section features-section" shadow="hover">
                      <template #header>
                        <span class="section-title">🎥 视频特征</span>
                      </template>
                      <div v-if="parsedAnalysisData.VideoFeatures" class="features-content">
                        <div class="feature-tags">
                          <el-tag v-if="parsedAnalysisData.VideoFeatures.is_real_person" type="success" size="small">
                            👤 真人出镜
                          </el-tag>
                          <el-tag type="info" size="small">
                            👥 {{ parsedAnalysisData.VideoFeatures.person_count || '0' }}人
                          </el-tag>
                          <el-tag type="primary" size="small">
                            🗣️ {{ parsedAnalysisData.VideoFeatures.voiceover_type || '未知' }}
                          </el-tag>
                          <el-tag type="warning" size="small">
                            🎬 {{ parsedAnalysisData.VideoFeatures.shot_type || '未知' }}
                          </el-tag>
                          <el-tag type="success" size="small">
                            🌐 {{ parsedAnalysisData.VideoFeatures.language || '未知' }}
                          </el-tag>
                        </div>
                        <div class="feature-item">
                          <div class="feature-label">📍 拍摄场景</div>
                          <div class="feature-value">{{ parsedAnalysisData.VideoFeatures.scene || '暂无' }}</div>
                        </div>
                      </div>
                    </el-card>
                  </el-col>
                </el-row>

                <!-- 评论洞察 -->
                <el-card v-if="parsedAnalysisData.CommentInsights" class="analysis-section comments-section" shadow="hover">
                  <template #header>
                    <span class="section-title">💬 评论洞察</span>
                  </template>
                  <div class="comments-grid">
                    <div class="comment-insight">
                      <div class="insight-label">🛍️ 产品反馈</div>
                      <div class="insight-content">{{ parsedAnalysisData.CommentInsights.product_feedback }}</div>
                    </div>
                    <div class="comment-insight">
                      <div class="insight-label">🎥 视频反馈</div>
                      <div class="insight-content">{{ parsedAnalysisData.CommentInsights.video_feedback }}</div>
                    </div>
                    <div class="comment-insight">
                      <div class="insight-label">💬 社交互动</div>
                      <div class="insight-content">{{ parsedAnalysisData.CommentInsights.social_interactions }}</div>
                    </div>
                    <div class="comment-insight">
                      <div class="insight-label">💰 转化信号</div>
                      <div class="insight-content">{{ parsedAnalysisData.CommentInsights.conversion_signals }}</div>
                    </div>
                    <div v-if="parsedAnalysisData.CommentInsights.risk_flags" class="comment-insight risk-insight">
                      <div class="insight-label">⚠️ 风险标识</div>
                      <div class="insight-content risk-content">{{ parsedAnalysisData.CommentInsights.risk_flags }}</div>
                    </div>
                  </div>
                </el-card>
              </div>
            </div>
          </el-col>
        </el-row>
      </div>

      <!-- 编辑模式：保持原有的表单结构 -->
      <template v-else>
        <el-form
          ref="dialogForm"
          :model="formData"
          :rules="rules"
          label-width="120px"
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
      </template>

      <template #footer>
        <div class="dialog-footer">
          <el-button @click="dialogVisible = false">{{ dialogType === 'view' ? '关闭' : '取消' }}</el-button>
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
      activeTimelineIndex: 0, // 当前激活的时间轴项
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
    },

    // 新增的分析方法
    async analyzeVideoData() {
      ElMessage.info('AI数据分析功能正在开发中，敬请期待！')
      // TODO: 实现AI数据分析逻辑
    },

    async fixVideoScript() {
      if (!this.formData.videoScript) {
        ElMessage.warning('暂无原始文案，无法进行修复')
        return
      }
      
      ElMessage.info('文案修复功能正在开发中，敬请期待！')
      // TODO: 实现文案修复逻辑
      // 示例逻辑：
      // try {
      //   const response = await fixScriptAPI({ 
      //     id: this.formData.ID, 
      //     script: this.formData.videoScript 
      //   })
      //   this.formData.fixedScript = response.data.fixedScript
      //   this.formData.fixNote = response.data.fixNote
      //   ElMessage.success('文案修复完成')
      //   this.getTableData()
      // } catch (error) {
      //   ElMessage.error('文案修复失败')
      // }
    },

    async analyzeVideoFrame() {
      if (!this.formData.videoLink) {
        ElMessage.warning('暂无视频链接，无法进行画面分析')
        return
      }
      
      ElMessage.info('视频画面分析功能正在开发中，敬请期待！')
      // TODO: 实现视频画面分析逻辑
      // 示例逻辑：
      // try {
      //   const response = await analyzeVideoAPI({ 
      //     id: this.formData.ID, 
      //     videoUrl: this.formData.videoLink 
      //   })
      //   this.formData.videoAnalysis = response.data.analysis
      //   ElMessage.success('画面分析完成')
      //   this.getTableData()
      // } catch (error) {
      //   ElMessage.error('画面分析失败')
      // }
    },

    async batchUpdateStatus() {
      if (!this.formData.ID) {
        ElMessage.warning('无效的数据ID')
        return
      }
      
      ElMessage.info('批量状态更新功能正在开发中，敬请期待！')
      // TODO: 实现批量状态更新逻辑
      // 可以打开一个新的对话框让用户选择要更新的状态
    },

    // 获取情感类型对应的标签颜色
    getEmotionType(emotion) {
      const emotionMap = {
        '开心': 'success',
        '兴奋': 'success', 
        '激动': 'success',
        '愉快': 'success',
        '惊讶': 'warning',
        '好奇': 'warning',
        '困惑': 'info',
        '平静': 'info',
        '中性': 'info',
        '担心': 'danger',
        '焦虑': 'danger',
        '悲伤': 'danger'
      };
      return emotionMap[emotion] || 'info';
    },

    // 获取分类类型对应的标签颜色
    getCategoryType(category) {
      const categoryMap = {
        '产品展示': 'primary',
        '产品介绍': 'primary',
        '使用演示': 'success',
        '效果展示': 'success',
        '情感共鸣': 'warning',
        '引流钩子': 'danger',
        '行动召唤': 'danger',
        'CTA': 'danger',
        '社交证明': 'info',
        '权威背书': 'info'
      };
      return categoryMap[category] || 'primary';
    },

    // 获取驱动因素类型对应的标签颜色
    getDriverType(driver) {
      const driverMap = {
        '好奇心': 'primary',
        '恐惧感': 'danger',
        'FOMO': 'danger',
        '权威': 'success',
        '社交证明': 'success',
        '稀缺性': 'warning',
        '紧迫性': 'warning',
        '从众心理': 'info',
        '情感共鸣': 'info'
      };
      return driverMap[driver] || 'info';
    },

    // 刷新分析数据
    async refreshAnalysisData() {
      try {
        // 重新获取当前记录的最新数据
        const res = await findVideoContentAnalysis(this.formData.ID);
        if (res.code === 0) {
          // 更新表单数据，特别是videoAnalysis字段
          this.formData = { ...this.formData, ...res.data };
          ElMessage.success('分析数据已刷新');
        } else {
          ElMessage.error('刷新失败：' + res.msg);
        }
      } catch (error) {
        console.error('刷新分析数据失败:', error);
        ElMessage.error('刷新失败，请检查网络连接');
      }
    }
  },

  computed: {
    // 解析视频分析JSON数据
    parsedAnalysisData() {
      if (!this.formData.videoAnalysis) return null;
      try {
        // 清理JSON字符串（去除markdown代码块标记）
        let cleanJson = this.formData.videoAnalysis.trim();
        if (cleanJson.startsWith('```json')) {
          cleanJson = cleanJson.replace(/```json\s*/, '').replace(/\s*```$/, '');
        }
        if (cleanJson.startsWith('```')) {
          cleanJson = cleanJson.replace(/```\s*/, '').replace(/\s*```$/, '');
        }
        
        const parsed = JSON.parse(cleanJson);
        return parsed;
      } catch (error) {
        console.error('Failed to parse analysis data:', error);
        return null;
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

/* 美化详情对话框样式 */
.video-detail-dialog.enhanced-view :deep(.el-dialog) {
  border-radius: 12px;
  overflow: hidden;
}

.video-detail-dialog.enhanced-view :deep(.el-dialog__header) {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  padding: 20px 24px;
}

.video-detail-dialog.enhanced-view :deep(.el-dialog__title) {
  font-size: 18px;
  font-weight: 600;
}

.video-detail-view {
  max-height: 70vh;
  overflow-y: auto;
  padding: 0 4px;
}

.detail-card {
  margin-bottom: 16px;
  border-radius: 8px;
  border: none;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.05);
}

.detail-card:hover {
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
  transform: translateY(-2px);
  transition: all 0.3s ease;
}

.card-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  font-weight: 600;
}

.card-title {
  font-size: 16px;
  color: #303133;
}

.header-tags, .header-actions {
  display: flex;
  gap: 8px;
}

.info-item {
  padding: 12px 0;
  border-bottom: 1px solid #f5f5f5;
}

.info-item:last-child {
  border-bottom: none;
}

.info-label {
  font-size: 14px;
  color: #909399;
  margin-bottom: 6px;
  font-weight: 500;
}

.info-value {
  font-size: 14px;
  color: #303133;
  line-height: 1.5;
  word-break: break-all;
}

.title-value {
  font-size: 16px;
  font-weight: 600;
  color: #409EFF;
}

.creator-name {
  font-weight: 600;
  color: #67C23A;
}

.link-value {
  word-break: break-all;
}

/* 统计数据样式 */
.stats-row {
  margin-bottom: 16px;
}

.stats-row:last-child {
  margin-bottom: 0;
}

.stat-item {
  display: flex;
  align-items: center;
  padding: 16px;
  background: #f8f9ff;
  border-radius: 8px;
  border-left: 4px solid #409EFF;
  transition: all 0.3s ease;
}

.stat-item:hover {
  background: #ecf5ff;
  transform: translateY(-2px);
}

.stat-item.secondary {
  background: #f9f9f9;
  border-left-color: #909399;
}

.stat-item.secondary:hover {
  background: #f5f5f5;
}

.stat-icon {
  font-size: 24px;
  margin-right: 12px;
  flex-shrink: 0;
}

.stat-content {
  flex: 1;
}

.stat-label {
  font-size: 12px;
  color: #909399;
  margin-bottom: 4px;
}

.stat-value {
  font-size: 16px;
  font-weight: 600;
  color: #303133;
}

.stat-value.highlight {
  color: #E6A23C;
  font-size: 18px;
}

.stat-value.money-value {
  color: #F56C6C;
}

/* 内容分析样式 */
.content-section {
  padding: 8px 0;
}

.content-item {
  padding: 16px 0;
  border-bottom: 1px solid #f0f0f0;
}

.content-item:last-child {
  border-bottom: none;
}

.content-label {
  font-size: 14px;
  font-weight: 600;
  color: #606266;
  margin-bottom: 8px;
}

.content-value {
  font-size: 14px;
  color: #303133;
  line-height: 1.6;
  white-space: pre-wrap;
  word-break: break-word;
  background: #f8f9fa;
  padding: 12px;
  border-radius: 6px;
  border-left: 3px solid #e0e0e0;
}

.content-value.script-content {
  background: #fff7e6;
  border-left-color: #faad14;
}

.content-value.fixed-script {
  background: #f6ffed;
  border-left-color: #52c41a;
}

.content-value.analysis-content {
  background: #e6f7ff;
  border-left-color: #1890ff;
}

.content-value.note-content {
  background: #fff2e8;
  border-left-color: #fa8c16;
}

/* 处理状态样式 */
.status-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 16px;
  padding: 8px 0;
}

.status-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 20px;
  background: #fafafa;
  border-radius: 8px;
  border: 2px solid #e8e8e8;
  transition: all 0.3s ease;
  text-align: center;
}

.status-item.active {
  background: #f0f9ff;
  border-color: #409EFF;
  box-shadow: 0 4px 12px rgba(64, 158, 255, 0.15);
}

.status-item:hover {
  transform: translateY(-3px);
  box-shadow: 0 6px 20px rgba(0, 0, 0, 0.1);
}

.status-icon {
  font-size: 32px;
  margin-bottom: 12px;
  filter: grayscale(100%);
  transition: filter 0.3s ease;
}

.status-item.active .status-icon {
  filter: none;
}

.status-text {
  font-size: 14px;
  font-weight: 600;
  color: #606266;
  margin-bottom: 8px;
}

.status-item.active .status-text {
  color: #409EFF;
}

/* 时间信息样式 */
.time-info {
  display: flex;
  justify-content: space-between;
  flex-wrap: wrap;
  gap: 16px;
  padding: 16px 20px;
  background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
  border-radius: 8px;
  margin-top: 16px;
}

.time-item {
  font-size: 13px;
  color: #606266;
  background: rgba(255, 255, 255, 0.8);
  padding: 6px 12px;
  border-radius: 4px;
  backdrop-filter: blur(5px);
}

/* 紧凑版统计样式 */
.compact-stats {
  padding: 8px 0;
}

.stat-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 8px 0;
  border-bottom: 1px solid #f5f5f5;
}

.stat-row:last-child {
  border-bottom: none;
}

.stat-row.highlight-row {
  background: linear-gradient(90deg, rgba(230, 162, 60, 0.1) 0%, transparent 100%);
  padding: 8px 12px;
  border-radius: 4px;
  border-bottom: 1px solid rgba(230, 162, 60, 0.2);
}

.stat-row.money-row {
  background: linear-gradient(90deg, rgba(245, 108, 108, 0.1) 0%, transparent 100%);
  padding: 8px 12px;
  border-radius: 4px;
  border-bottom: 1px solid rgba(245, 108, 108, 0.2);
}

.stat-label {
  font-size: 12px;
  color: #909399;
  font-weight: 500;
}

.stat-value {
  font-size: 14px;
  font-weight: 600;
  color: #303133;
}

.stat-value.highlight {
  color: #E6A23C;
  font-size: 16px;
}

.stat-value.money-value {
  color: #F56C6C;
}

/* 紧凑版状态样式 */
.compact-status {
  padding: 8px 0;
}

.status-row {
  display: flex;
  align-items: center;
  padding: 8px 0;
  border-bottom: 1px solid #f5f5f5;
  transition: all 0.3s ease;
}

.status-row:last-child {
  border-bottom: none;
}

.status-row.active {
  background: linear-gradient(90deg, rgba(64, 158, 255, 0.1) 0%, transparent 100%);
  padding: 8px 12px;
  border-radius: 4px;
  border-bottom: 1px solid rgba(64, 158, 255, 0.2);
}

.status-icon {
  font-size: 16px;
  margin-right: 8px;
  filter: grayscale(100%);
  transition: filter 0.3s ease;
}

.status-row.active .status-icon {
  filter: none;
}

.status-text {
  flex: 1;
  font-size: 13px;
  color: #606266;
  margin-right: 8px;
}

.status-row.active .status-text {
  color: #409EFF;
  font-weight: 500;
}

/* 紧凑版时间信息 */
.time-info.compact {
  display: flex;
  flex-direction: column;
  gap: 8px;
  padding: 12px 16px;
  background: linear-gradient(135deg, #f8f9fa 0%, #e9ecef 100%);
  border-radius: 6px;
  margin-top: 12px;
}

.time-item {
  font-size: 12px;
  color: #666;
  background: rgba(255, 255, 255, 0.8);
  padding: 4px 8px;
  border-radius: 4px;
  backdrop-filter: blur(3px);
}

/* 视频分析面板样式 */
.video-analysis-panel {
  background: #fafbfc;
  border-radius: 8px;
  overflow: hidden;
}

.analysis-header-card {
  margin-bottom: 16px;
  border-radius: 8px;
  border: none;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06);
}

.analysis-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.header-info {
  flex: 1;
}

.analysis-title {
  font-size: 18px;
  font-weight: 600;
  color: #303133;
  display: block;
  margin-bottom: 4px;
}

.analysis-subtitle {
  font-size: 14px;
  color: #909399;
  font-weight: normal;
}

.header-actions {
  display: flex;
  gap: 8px;
}

.no-analysis {
  text-align: center;
  padding: 60px 20px;
}

/* 分析内容区域 */
.analysis-content {
  padding: 0 4px;
}

.analysis-section {
  margin-bottom: 20px;
  border-radius: 10px;
  border: none;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);
  overflow: hidden;
}

.analysis-section:hover {
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.12);
  transform: translateY(-2px);
  transition: all 0.3s ease;
}

.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.section-title {
  font-size: 16px;
  font-weight: 600;
  color: #303133;
}

/* 时间轴样式 */
.timeline-section {
  background: linear-gradient(135deg, #fff 0%, #f8f9ff 100%);
}

.timeline-container {
  max-height: 400px;
  overflow-y: auto;
  padding-right: 8px;
}

.timeline-item {
  display: flex;
  margin-bottom: 16px;
  padding: 16px;
  background: rgba(255, 255, 255, 0.8);
  border-radius: 8px;
  border-left: 4px solid #e0e0e0;
  cursor: pointer;
  transition: all 0.3s ease;
}

.timeline-item:hover {
  background: rgba(255, 255, 255, 0.95);
  border-left-color: #409EFF;
  box-shadow: 0 2px 8px rgba(64, 158, 255, 0.15);
}

.timeline-item.active {
  background: rgba(64, 158, 255, 0.05);
  border-left-color: #409EFF;
  box-shadow: 0 4px 12px rgba(64, 158, 255, 0.2);
}

.timeline-time {
  min-width: 80px;
  font-weight: 600;
  color: #409EFF;
  font-size: 14px;
  margin-right: 16px;
  padding-top: 2px;
}

.timeline-content {
  flex: 1;
}

.scene-info {
  margin-bottom: 12px;
}

.scene-title {
  font-size: 15px;
  font-weight: 600;
  color: #303133;
  margin-bottom: 6px;
}

.scene-actions {
  font-size: 14px;
  color: #606266;
  line-height: 1.5;
}

.scene-details {
  margin-bottom: 8px;
}

.scene-details .el-tag {
  margin-right: 6px;
  margin-bottom: 4px;
}

.audio-text {
  background: #f0f9ff;
  padding: 8px 12px;
  border-radius: 6px;
  font-size: 13px;
  color: #1890ff;
  border-left: 3px solid #1890ff;
  margin-top: 8px;
}

.audio-icon {
  margin-right: 6px;
}

/* 重点时刻样式 */
.highlights-section {
  background: linear-gradient(135deg, #fff5f5 0%, #fff 100%);
}

.highlights-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 16px;
}

.highlight-item {
  background: rgba(255, 255, 255, 0.9);
  border-radius: 8px;
  padding: 16px;
  border-left: 4px solid #f56c6c;
  box-shadow: 0 2px 8px rgba(245, 108, 108, 0.1);
  transition: all 0.3s ease;
}

.highlight-item:hover {
  transform: translateY(-3px);
  box-shadow: 0 6px 20px rgba(245, 108, 108, 0.2);
}

.highlight-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
}

.highlight-time {
  font-weight: 600;
  color: #f56c6c;
  font-size: 12px;
}

.highlight-label {
  font-size: 15px;
  font-weight: 600;
  color: #303133;
  margin-bottom: 8px;
}

.highlight-description {
  font-size: 13px;
  color: #606266;
  line-height: 1.5;
  margin-bottom: 12px;
}

.highlight-effect {
  margin-bottom: 12px;
}

.effect-label {
  font-size: 12px;
  color: #909399;
  margin-right: 6px;
}

.effect-text {
  font-size: 13px;
  color: #303133;
}

.highlight-drivers {
  margin-bottom: 12px;
}

.driver-tag {
  margin-right: 6px;
  margin-bottom: 4px;
}

.highlight-metrics {
  display: flex;
  gap: 8px;
  flex-wrap: wrap;
}

/* 总结和产品信息样式 */
.summary-row {
  margin-bottom: 20px;
}

.summary-section, .product-section {
  height: 100%;
}

.summary-content, .product-content {
  padding: 8px 0;
}

.summary-item, .product-item {
  margin-bottom: 16px;
}

.summary-item:last-child, .product-item:last-child {
  margin-bottom: 0;
}

.summary-label, .product-label {
  font-size: 13px;
  font-weight: 600;
  color: #606266;
  margin-bottom: 6px;
}

.summary-value, .product-value {
  font-size: 14px;
  color: #303133;
  line-height: 1.5;
  background: #f8f9fa;
  padding: 8px 12px;
  border-radius: 6px;
  border-left: 3px solid #e0e0e0;
}

.visual-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 6px;
}

.visual-tag {
  margin: 0;
}

.no-product-info {
  text-align: center;
  padding: 40px 20px;
}

/* 底部分析行样式 */
.analysis-bottom-row {
  margin-bottom: 20px;
}

.structure-section, .painpoints-section, .features-section {
  height: 100%;
}

.structure-content, .painpoints-content, .features-content {
  padding: 8px 0;
}

.structure-item, .painpoint-item {
  margin-bottom: 12px;
}

.structure-item:last-child, .painpoint-item:last-child {
  margin-bottom: 0;
}

.structure-label, .painpoint-label, .feature-label {
  font-size: 12px;
  font-weight: 600;
  color: #606266;
  margin-bottom: 4px;
}

.structure-value, .painpoint-value, .feature-value {
  font-size: 13px;
  color: #303133;
  line-height: 1.4;
}

.feature-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 6px;
  margin-bottom: 12px;
}

.feature-item {
  margin-bottom: 8px;
}

/* 评论洞察样式 */
.comments-section {
  background: linear-gradient(135deg, #f0f9ff 0%, #fff 100%);
}

.comments-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 16px;
}

.comment-insight {
  background: rgba(255, 255, 255, 0.8);
  padding: 16px;
  border-radius: 8px;
  border-left: 4px solid #1890ff;
  box-shadow: 0 2px 8px rgba(24, 144, 255, 0.1);
}

.comment-insight.risk-insight {
  border-left-color: #f5222d;
  box-shadow: 0 2px 8px rgba(245, 34, 45, 0.1);
}

.insight-label {
  font-size: 13px;
  font-weight: 600;
  color: #1890ff;
  margin-bottom: 8px;
}

.risk-insight .insight-label {
  color: #f5222d;
}

.insight-content {
  font-size: 13px;
  color: #303133;
  line-height: 1.5;
}

.risk-content {
  color: #f5222d;
  font-weight: 500;
}

/* 响应式调整 */
@media (max-width: 768px) {
  .stats-row .el-col {
    margin-bottom: 8px;
  }
  
  .status-grid {
    grid-template-columns: repeat(auto-fit, minmax(150px, 1fr));
    gap: 12px;
  }
  
  .time-info {
    flex-direction: column;
    gap: 8px;
  }

  .highlights-grid, .comments-grid {
    grid-template-columns: 1fr;
  }
  
  .timeline-item {
    flex-direction: column;
  }
  
  .timeline-time {
    min-width: auto;
    margin-right: 0;
    margin-bottom: 8px;
  }
  
  .analysis-bottom-row .el-col {
    margin-bottom: 16px;
  }
  
  .header-actions {
    flex-wrap: wrap;
    gap: 4px;
  }
  
  .header-actions .el-button {
    font-size: 12px;
    padding: 6px 10px;
  }
}
</style>
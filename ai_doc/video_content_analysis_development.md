# 视频内容聚合分析系统开发文档

## 项目概述

基于 Excel 文件 `temp/测试聚合层文档_分析_最终_1.xlsx` 开发的视频内容聚合分析系统，用于管理和分析视频内容数据。

## 数据源分析

### Excel 文件结构
- **文件路径**: `temp/测试聚合层文档_分析_最终_1.xlsx`
- **数据记录**: 299条视频记录
- **字段数量**: 25个字段

### 核心数据字段

#### 基础信息字段
- **视频描述** (`video_description`): 视频内容描述
- **视频链接** (`video_link`): TikTok等平台的视频链接地址
- **视频标题** (`video_title`): 视频标题
- **发布时间** (`publish_time`): 视频发布时间
- **视频时长** (`video_duration`): 视频播放时长

#### 达人信息字段
- **达人名称** (`creator_name`): 视频创作者名称
- **Unique Id** (`unique_id`): 创作者唯一标识符
- **粉丝数** (`fans_count`): 达人粉丝数量

#### 互动数据字段
- **播放量** (`play_count`): 视频播放次数
- **点赞数** (`like_count`): 视频点赞数量
- **评论数** (`comment_count`): 视频评论数量
- **转发数** (`share_count`): 视频转发数量

#### 商业数据字段
- **销量(件)** (`sales_volume`): 商品销售件数
- **销售额** (`sales_amount`): 销售金额

#### 内容分析字段
- **视频文案** (`video_script`): 原始视频文案内容
- **修复后文案** (`fixed_script`): AI修复后的文案
- **视频画面分析** (`video_analysis`): 视频画面分析结果
- **修复说明** (`fix_note`): 文案修复说明

#### 处理状态字段
- **是否下载** (`is_downloaded`): 视频下载状态
- **是否提取音频** (`is_audio_extracted`): 音频提取状态
- **是否转文本** (`is_text_converted`): 文本转换状态
- **是否修复文案** (`is_script_fixed`): 文案修复状态
- **是否分析视频画面** (`is_video_analyzed`): 画面分析状态

## 技术架构

### 后端技术栈
- **框架**: Gin + GORM
- **数据库**: MySQL (通过GVA_MODEL包含ID、创建时间、更新时间、删除时间)
- **API文档**: Swagger
- **包结构**: 按照GVA框架规范的分层架构

### 前端技术栈
- **框架**: Vue 3 + Composition API
- **UI组件**: Element Plus
- **构建工具**: Vite
- **状态管理**: Pinia

## 开发实现

### 数据模型设计

#### 主表结构 (`video_content_analysis`)
```sql
CREATE TABLE video_content_analysis (
  id bigint unsigned NOT NULL AUTO_INCREMENT,
  created_at datetime(3) DEFAULT NULL,
  updated_at datetime(3) DEFAULT NULL,
  deleted_at datetime(3) DEFAULT NULL,
  video_description varchar(255) DEFAULT '' COMMENT '视频描述内容',
  video_link varchar(500) NOT NULL COMMENT '视频链接地址',
  creator_name varchar(100) NOT NULL COMMENT '视频创作者达人名称',
  unique_id varchar(100) DEFAULT '' COMMENT '视频创作者唯一标识',
  fans_count varchar(20) DEFAULT '' COMMENT '达人粉丝数量',
  play_count varchar(20) DEFAULT '' COMMENT '视频播放量',
  like_count varchar(20) DEFAULT '' COMMENT '视频点赞数',
  comment_count int DEFAULT 0 COMMENT '视频评论数',
  share_count int DEFAULT 0 COMMENT '视频转发数',
  sales_volume int DEFAULT 0 COMMENT '商品销量件数',
  sales_amount varchar(20) DEFAULT '' COMMENT '商品销售金额',
  publish_time varchar(50) DEFAULT '' COMMENT '视频发布时间',
  video_duration varchar(20) DEFAULT '' COMMENT '视频播放时长',
  video_title varchar(255) DEFAULT '' COMMENT '视频标题',
  video_script text COMMENT '原始视频文案内容',
  fixed_script text COMMENT '修复后的视频文案内容',
  video_analysis text COMMENT '视频画面分析内容',
  is_downloaded tinyint(1) DEFAULT 0 COMMENT '是否已下载视频',
  is_audio_extracted tinyint(1) DEFAULT 0 COMMENT '是否已提取音频',
  is_text_converted tinyint(1) DEFAULT 0 COMMENT '是否已转为文本',
  is_script_fixed tinyint(1) DEFAULT 0 COMMENT '是否已修复文案',
  is_video_analyzed tinyint(1) DEFAULT 0 COMMENT '是否已分析视频画面',
  fix_note varchar(500) DEFAULT '' COMMENT '修复说明备注',
  PRIMARY KEY (id),
  KEY idx_video_content_analysis_deleted_at (deleted_at)
);
```

### 后端API设计

#### API路由结构
```
/api/v1/videoContent/
├── createVideoContentAnalysis [POST]        # 创建视频内容分析
├── deleteVideoContentAnalysis [DELETE]      # 删除视频内容分析
├── deleteVideoContentAnalysisByIds [DELETE] # 批量删除
├── updateVideoContentAnalysis [PUT]         # 更新视频内容分析
├── findVideoContentAnalysis [GET]           # 根据ID查询
├── getVideoContentAnalysisList [GET]        # 分页列表查询
├── getVideoContentAnalysisStatistics [GET]  # 获取统计信息
├── importVideoContentAnalysisFromExcel [POST] # Excel导入
├── batchUpdateProcessStatus [PUT]           # 批量更新处理状态
└── getVideoContentAnalysisPublic [GET]      # 公开接口
```

#### 核心功能接口

**1. 分页列表查询**
- 支持按达人名称、视频标题模糊搜索
- 支持按播放量范围筛选
- 支持按下载状态筛选
- 支持按创建时间范围筛选

**2. 统计分析接口**
- 总视频数统计
- 各处理状态数量统计
- 头部达人排行榜（按视频数量）
- 播放量、点赞数汇总

**3. 批量操作接口**
- 批量删除视频记录
- 批量更新处理状态
- Excel数据批量导入

### 前端页面设计

#### 列表页面展示字段
**主要展示字段**:
- 视频标题 (video_title)
- 达人名称 (creator_name)
- 播放量 (play_count)
- 点赞数 (like_count)
- 销量 (sales_volume)
- 销售额 (sales_amount)
- 发布时间 (publish_time)
- 是否下载 (is_downloaded)

#### 详情页面展示字段
**完整字段展示**:
- 基础信息: 视频描述、链接、标题、时长
- 达人信息: 名称、唯一ID、粉丝数
- 互动数据: 播放量、点赞数、评论数、转发数
- 商业数据: 销量、销售额
- 内容分析: 视频文案、修复后文案、画面分析
- 处理状态: 所有状态字段
- 其他: 修复说明、时间戳

#### 功能特性
- **搜索筛选**: 支持多条件组合搜索
- **批量操作**: 支持批量删除和状态更新
- **数据导入**: 支持Excel文件上传导入
- **统计图表**: 数据可视化展示
- **详情弹窗**: 完整信息查看
- **响应式设计**: 适配不同屏幕尺寸

## 名词定义统一

| 术语 | 中文名称 | 英文标识 | 说明 |
|------|----------|----------|------|
| Video Content Analysis | 视频内容分析 | VideoContentAnalysis | 主要数据模型 |
| Creator | 达人/创作者 | creator | 视频创作者 |
| Play Count | 播放量 | playCount | 视频播放次数 |
| Like Count | 点赞数 | likeCount | 视频点赞数量 |
| Sales Volume | 销量 | salesVolume | 商品销售件数 |
| Sales Amount | 销售额 | salesAmount | 销售金额 |
| Video Script | 视频文案 | videoScript | 原始视频文案 |
| Fixed Script | 修复后文案 | fixedScript | AI修复后的文案 |
| Video Analysis | 视频画面分析 | videoAnalysis | 视频内容分析结果 |
| Process Status | 处理状态 | processStatus | 各种处理步骤的状态 |

## 权限设置

### API权限配置
- 视频内容管理相关的API需要登录权限
- 公开查询接口无需权限验证
- 管理员具有所有操作权限
- 普通用户具有查看权限

### 菜单权限配置
- 菜单路径: `/videoContent`
- 菜单名称: `视频内容分析`
- 父级菜单: 可归属于数据管理或内容管理模块
- 按钮权限: 增删改查、导入导出、统计分析

## 开发进度

- [x] Excel数据分析和需求梳理
- [x] 数据模型设计和创建
- [x] 后端API开发 (Model/Service/API层)
- [x] Swagger接口文档
- [ ] 前端页面开发
- [ ] 数据导入功能实现
- [ ] 统计图表开发
- [ ] 权限配置
- [ ] 功能测试
- [ ] 性能优化

## 注意事项

1. **数据类型处理**: 播放量、点赞数等使用字符串存储，支持K/M单位转换
2. **布尔字段**: 使用指针类型以支持null值
3. **文本内容**: 使用TEXT类型存储长文本内容
4. **索引优化**: 在常用查询字段上添加数据库索引
5. **数据验证**: 前后端都需要进行数据格式验证
6. **权限控制**: 确保敏感操作需要适当权限

## 扩展功能建议

1. **数据可视化**: 增加更多统计图表
2. **导出功能**: 支持数据导出到Excel/CSV
3. **定时任务**: 自动抓取更新视频数据
4. **标签系统**: 为视频内容添加分类标签
5. **评价系统**: 对视频质量进行评分
6. **API限流**: 防止接口被恶意调用

---

*文档生成时间: 2024-09-11*  
*开发框架: gin-vue-admin (GVA)*  
*开发者: AI Assistant*
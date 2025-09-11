# 视频内容聚合分析系统 - 部署指南

## 🚀 快速部署

### 1. 系统要求
- **操作系统**: macOS/Linux/Windows
- **Go**: 1.21+
- **Node.js**: 18+
- **MySQL**: 8.0+
- **Python**: 3.9+ (用于数据导入)
- **Conda**: 用于Python环境管理

### 2. 项目结构
```
gin-vue-admin/
├── server/                 # 后端Go服务
├── web/                   # 前端Vue应用
├── temp/                  # Excel数据文件
├── scripts/               # 数据导入脚本
├── ai_doc/               # 项目文档
└── README.md
```

## 📦 部署步骤

### Step 1: 环境准备

#### 1.1 安装Conda环境
```bash
# 创建Python环境
conda create -n gva python=3.9
conda activate gva

# 安装Python依赖
pip install pandas openpyxl requests
```

#### 1.2 配置数据库
```sql
-- 创建数据库
CREATE DATABASE gva DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- 创建用户(可选)
CREATE USER 'gva'@'localhost' IDENTIFIED BY 'your_password';
GRANT ALL PRIVILEGES ON gva.* TO 'gva'@'localhost';
FLUSH PRIVILEGES;
```

### Step 2: 后端部署

```bash
# 1. 进入后端目录
cd server/

# 2. 安装Go依赖
go mod tidy

# 3. 配置数据库连接 (修改 config.yaml)
vim config.yaml

# 配置示例:
# mysql:
#   path: "127.0.0.1:3306"
#   db-name: "gva"
#   username: "root"
#   password: "your_password"

# 4. 运行数据库迁移和启动服务
go run main.go
```

### Step 3: 前端部署

```bash
# 1. 进入前端目录
cd web/

# 2. 安装依赖
npm install

# 3. 开发环境运行
npm run dev

# 4. 生产环境构建
npm run build
```

### Step 4: 数据导入

#### 4.1 自动导入 (推荐)
```bash
# 进入scripts目录
cd scripts/

# 执行自动导入脚本
./run_import.sh
```

#### 4.2 手动导入
```bash
# 激活conda环境
conda activate gva

# 执行Python导入脚本
python import_video_data.py ../temp/测试聚合层文档_分析_最终_1.xlsx http://localhost:8888/api/v1
```

### Step 5: 权限配置

1. **访问管理后台**: http://localhost:8080
2. **登录**: 使用默认管理员账号
3. **配置权限**:
   ```
   系统管理 → 角色管理 → 选择角色
   ✅ 菜单权限: 勾选"视频内容分析" 
   ✅ API权限: 勾选"视频内容管理"分组下的所有API
   保存配置
   ```

## 🔧 配置详解

### 数据库配置 (config.yaml)
```yaml
mysql:
  path: "127.0.0.1:3306"
  port: "3306"
  db-name: "gva"
  username: "root"
  password: "your_password"
  max-idle-conns: 10
  max-open-conns: 100
  log-mode: "error"
  log-zap: false
```

### 服务器配置
```yaml
system:
  env: "develop"  # develop | test | production
  addr: 8888
  db-type: "mysql"
  oss-type: "local"
  use-multipoint: false
  use-redis: false
```

### 跨域配置
```yaml
cors:
  mode: "allow-all" # 开发环境
  whitelist:
    - allow-origin: "*"
      allow-methods: "GET,POST,PUT,DELETE,OPTIONS"
      allow-headers: "content-type,authorization"
      expose-headers: "Content-Length"
```

## 📊 数据导入说明

### Excel文件要求
- **文件格式**: .xlsx 或 .xls
- **必要字段**: 达人名称、视频链接
- **字段映射**: 
  ```
  视频描述 → video_description
  视频链接 → video_link
  达人名称 → creator_name  
  粉丝数 → fans_count
  播放量 → play_count
  ... (详见导入脚本)
  ```

### 导入过程
1. **验证文件格式**: 检查Excel文件结构
2. **数据清洗**: 处理空值、数据类型转换
3. **字段映射**: Excel列名映射到数据库字段
4. **批量导入**: 逐条插入数据库
5. **结果统计**: 显示成功/失败数量

### 常见导入问题
- **文件不存在**: 检查Excel文件路径
- **字段缺失**: 确保必要字段存在
- **数据格式错误**: 检查数值、布尔值格式
- **API连接失败**: 确保后端服务运行正常

## 🌐 访问地址

### 开发环境
- **前端**: http://localhost:3000
- **后端API**: http://localhost:8888
- **Swagger文档**: http://localhost:8888/swagger/index.html
- **管理后台**: http://localhost:8080

### 生产环境
根据实际部署配置调整端口和域名

## 🔐 权限配置

### API权限列表
| API路径 | 方法 | 描述 | 权限组 |
|---------|------|------|--------|
| `/videoContent/getVideoContentAnalysisList` | GET | 获取列表 | 视频内容管理 |
| `/videoContent/findVideoContentAnalysis` | GET | 查看详情 | 视频内容管理 |
| `/videoContent/createVideoContentAnalysis` | POST | 创建记录 | 视频内容管理 |
| `/videoContent/updateVideoContentAnalysis` | PUT | 更新记录 | 视频内容管理 |
| `/videoContent/deleteVideoContentAnalysis` | DELETE | 删除记录 | 视频内容管理 |
| `/videoContent/deleteVideoContentAnalysisByIds` | DELETE | 批量删除 | 视频内容管理 |
| `/videoContent/getVideoContentAnalysisStatistics` | GET | 统计信息 | 视频内容管理 |
| `/videoContent/importVideoContentAnalysisFromExcel` | POST | Excel导入 | 视频内容管理 |
| `/videoContent/batchUpdateProcessStatus` | PUT | 批量更新 | 视频内容管理 |
| `/videoContent/getVideoContentAnalysisPublic` | GET | 公开接口 | 无需权限 |

### 菜单权限
- **菜单名称**: 视频内容分析
- **路由路径**: videoContentAnalysis
- **组件路径**: view/example/videoContentAnalysis/videoContentAnalysis.vue
- **菜单按钮**: 新增、编辑、删除、查看、导入、统计

## 🚨 故障排查

### 常见问题

#### 1. 后端启动失败
```bash
# 检查端口占用
lsof -i :8888

# 检查数据库连接
mysql -h localhost -u root -p

# 查看后端日志
tail -f server/log/server.log
```

#### 2. 前端访问失败
```bash
# 检查Node版本
node -v

# 清除缓存重新安装
rm -rf node_modules package-lock.json
npm install

# 检查端口
lsof -i :3000
```

#### 3. 数据导入失败
```bash
# 检查Python环境
conda activate gva
python -c "import pandas, requests"

# 检查Excel文件
file temp/测试聚合层文档_分析_最终_1.xlsx

# 检查API连通性  
curl http://localhost:8888/api/v1/health
```

#### 4. 权限配置问题
- 确保角色已分配相应的API和菜单权限
- 检查用户是否属于正确的角色
- 刷新浏览器缓存或重新登录

### 日志文件位置
```
server/log/
├── server.log      # 服务器日志
├── access.log      # 访问日志
└── error.log       # 错误日志
```

## 🔄 更新升级

### 代码更新
```bash
# 拉取最新代码
git pull origin main

# 后端更新
cd server && go mod tidy && go build

# 前端更新  
cd web && npm install && npm run build
```

### 数据库升级
```bash
# 运行数据库迁移
cd server && go run main.go

# 手动执行SQL(如需要)
mysql -u root -p gva < upgrade.sql
```

## 📈 性能优化

### 后端优化
- 启用Redis缓存
- 配置数据库连接池
- 优化SQL查询索引
- 启用gzip压缩

### 前端优化
- 启用CDN加速
- 配置浏览器缓存
- 启用代码分割
- 图片懒加载

### 数据库优化
```sql
-- 添加索引
ALTER TABLE video_content_analysis ADD INDEX idx_creator_name (creator_name);
ALTER TABLE video_content_analysis ADD INDEX idx_video_title (video_title);
ALTER TABLE video_content_analysis ADD INDEX idx_publish_time (publish_time);
```

## 🛡️ 安全建议

### 生产环境配置
1. **HTTPS**: 启用SSL证书
2. **防火墙**: 限制端口访问
3. **数据库**: 使用专用数据库用户
4. **文件上传**: 限制文件类型和大小
5. **API限流**: 防止接口滥用
6. **日志审计**: 记录敏感操作

### 备份策略
```bash
# 数据库备份
mysqldump -u root -p gva > backup_$(date +%Y%m%d).sql

# 文件备份
tar -czf data_backup_$(date +%Y%m%d).tar.gz temp/ uploads/
```

---

## 📞 技术支持

### 项目信息
- **版本**: v1.0.0
- **更新时间**: 2024年9月11日
- **开发框架**: gin-vue-admin
- **技术栈**: Go + Vue3 + MySQL

### 获取帮助
- **文档**: 查看 `ai_doc/` 目录下的详细文档
- **日志**: 查看服务器日志文件排查问题
- **社区**: 参考gin-vue-admin官方文档和社区

---

**🎉 部署完成后，您就可以开始使用视频内容聚合分析系统了！**
# MCP GVA Helper 功能模块 v1

## 模块概述
MCP GVA Helper 是针对 gin-vue-admin 框架开发的智能辅助工具集，提供了完整的代码生成、模块创建和系统分析功能。该模块包含8个核心工具，支持从需求分析到代码生成的完整开发流程。

## 核心工具清单

### 1. requirement_analyzer (需求分析器) - 🚀 首选入口工具
- **功能**: 将用户自然语言需求转换为AI可理解的结构化提示词
- **优先级**: 最高优先级，所有开发工作的起点
- **工作流**: 接收用户需求 → 生成结构化分析 → 指导后续工具使用
- **输出格式**: 1xxx2xxx格式的清晰逻辑步骤

### 2. gva_auto_generate (核心执行器) - 🔧 主力生成工具
- **功能**: 接收requirement_analyzer分析结果，执行具体的模块创建操作
- **特性**:
  - 支持批量创建多个模块
  - 自动生成API权限和菜单项
  - 智能字典创建功能
  - 支持package和plugin两种类型
- **工作模式**: analyze → confirm → execute 三步骤流程

### 3. create_api (API创建器)
- **功能**: 创建后端API记录，用于权限管理
- **使用限制**: 仅在单独创建API或AI编辑器自动添加API时使用
- **重要**: 当needCreatedModules=true时，由gva_auto_generate自动处理

### 4. create_menu (菜单创建器)
- **功能**: 创建前端菜单记录，用于导航管理
- **使用限制**: 仅在单独创建菜单或AI编辑器自动添加前端页面时使用
- **重要**: 当needCreatedModules=true时，由gva_auto_generate自动处理

### 5. generate_dictionary_options (字典生成器)
- **功能**: 智能生成字典选项并自动创建字典和字典详情
- **支持场景**: 状态、性别、类型、等级、优先级、审批、角色、布尔值、订单、颜色、尺寸等
- **特性**: 为无法识别的字典类型提供通用默认选项

### 6. list_all_apis (API查询器)
- **功能**: 获取系统中所有API接口信息
- **返回数据**:
  - 数据库已注册的API列表(完整信息)
  - gin框架实际注册的路由API列表(路径+方法)
- **用途**: 帮助判断是否使用现有API还是创建新API

### 7. list_all_menus (菜单查询器)
- **功能**: 获取系统中所有菜单信息
- **返回数据**: 完整菜单树形结构、路由配置、组件路径、元数据等
- **用途**: 前端路由配置、菜单权限管理、导航组件开发、系统架构分析

### 8. query_dictionaries (字典查询器)
- **功能**: 查询系统中所有字典和字典属性
- **特性**: 支持精确查询、包含禁用项查询、仅详情查询等模式
- **用途**: AI生成逻辑时了解可用的字典选项

## 推荐工作流程

### 标准开发流程
1. **需求分析阶段**: requirement_analyzer (最高优先级)
2. **代码生成阶段**: gva_auto_generate (执行创建)
3. **辅助完善阶段**: 根据需要使用其他辅助工具

### 工具使用规则
- **批量操作**: gva_auto_generate 支持在单次操作中创建多个模块
- **自动化程度**: 当needCreatedModules=true时，API和菜单会自动生成，无需手动调用对应工具
- **字典智能化**: 使用字典类型(dictType)时，系统自动检查并创建不存在的字典

## 版本特性

### v1 核心特性
- ✅ 完整的需求分析到代码生成工作流
- ✅ 智能字典自动创建功能
- ✅ 批量模块创建支持
- ✅ API和菜单自动生成
- ✅ 多数据库支持和关联配置
- ✅ package和plugin双模式支持

### 技术支持
- **字段类型**: string, richtext, int, bool, float64, time.Time, enum, picture, pictures, video, file, json, array
- **搜索类型**: =, !=, >, >=, <, <=, NOT BETWEEN, LIKE, BETWEEN, IN, NOT IN
- **关联关系**: 一对一、一对多关联配置
- **模型支持**: GVA Model (自动ID/时间戳) 和自定义主键模式

## 使用注意事项

1. **工具优先级**: 优先使用 requirement_analyzer → gva_auto_generate 的标准流程
2. **避免重复**: 当使用gva_auto_generate且needCreatedModules=true时，不要再调用create_api和create_menu
3. **字典智能**: 利用自动字典创建功能，减少手动配置
4. **模块关联**: 正确配置dataSource实现模块间关联关系
5. **主键配置**: gvaModel=false时必须设置一个字段的primaryKey=true

## 文档版本
- **版本**: v1
- **创建时间**: 2025-09-14
- **适用范围**: gin-vue-admin框架 MCP辅助开发
- **维护状态**: 活跃维护
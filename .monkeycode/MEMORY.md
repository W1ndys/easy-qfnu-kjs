# 用户指令记忆

本文件记录了用户的指令、偏好和教导，用于在未来的交互中提供参考。

## 格式

### 用户指令条目
用户指令条目应遵循以下格式：

[用户指令摘要]
- Date: [YYYY-MM-DD]
- Context: [提及的场景或时间]
- Instructions:
  - [用户教导或指示的内容，逐行描述]

### 项目知识条目
Agent 在任务执行过程中发现的条目应遵循以下格式：

[项目知识摘要]
- Date: [YYYY-MM-DD]
- Context: Agent 在执行 [具体任务描述] 时发现
- Category: [代码结构|代码模式|代码生成|构建方法|测试方法|依赖关系|环境配置]
- Instructions:
  - [具体的知识点，逐行描述]

## 去重策略
- 添加新条目前，检查是否存在相似或相同的指令
- 若发现重复，跳过新条目或与已有条目合并
- 合并时，更新上下文或日期信息
- 这有助于避免冗余条目，保持记忆文件整洁

## 条目

### 项目技术架构
- Date: 2026-03-24
- Context: Agent 在执行前端设计改造任务时发现
- Category: 代码结构
- Instructions:
  - 前端使用 Vue 3.5 (Composition API + `<script setup>`) + Vue Router 4.5 + Axios 1.12
  - 构建工具为 Vite 7.1，样式使用 Tailwind CSS 3.4
  - 前端构建产物输出到 `../web` 目录，嵌入 Go 后端通过 `embed` 静态文件服务
  - 主色调为 `rgb(136, 79, 34)` 即 `#884F22`（棕色系）
  - 项目包含 3 个路由页面：首页、空教室查询、教室全天状态
  - 8 个组件：AppHeader, AppFooter, DateSelector, EmptyState, LoadingSpinner, QRCodeCard, StatsCard, StatusWarning
  - 3 个 Composables：useDateSelection, useSearchHistory, useSystemStatus
  - CSS 变量定义在 `src/assets/css/main.css` 中

### 用户主色调偏好
- Date: 2026-03-24
- Context: 用户在设计改造需求中明确指出
- Instructions:
  - 主体颜色必须保持为 `rgb(136, 79, 34)`，即 `#884F22`
  - 设计方案中的紫色系主色调需替换为此棕色系

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

### 速率限制中间件
- Date: 2026-04-10
- Context: Agent 在执行搜索接口限流功能时发现
- Category: 代码模式
- Instructions:
  - 速率限制中间件位于 `internal/middleware/ratelimit.go`
  - 使用 IP + User-Agent 组合作为限流 key，通过 SHA256 哈希 UA 避免 map key 过长
  - 内存 map + sync.Mutex 实现，无第三方依赖
  - 后台协程定期清理过期条目，防止内存泄漏
  - 在 main.go 中以路由级中间件方式应用，仅作用于 `/api/v1/query` 和 `/api/v1/query-full-day`

### 用户主色调偏好
- Date: 2026-03-24
- Context: 用户在设计改造需求中明确指出
- Instructions:
  - 主体颜色必须保持为 `rgb(136, 79, 34)`，即 `#884F22`
  - 设计方案中的紫色系主色调需替换为此棕色系

### 部署方式迁移要求
- Date: 2026-04-24
- Context: 用户在说明当前系统部署方案迁移时明确指出
- Instructions:
  - 系统改为使用 Docker 部署，前后端分别由 Docker 容器运行
  - 前后端通过 Docker 网络进行内网交互
  - 下游统一由 Traefik 负责负载均衡与反向代理
  - 不再使用旧的 task deploy 流程
  - 不再通过二进制文件和进程守护方式运行服务

### 运维 Taskfile 要求
- Date: 2026-04-25
- Context: 用户在补充开发运维能力时明确指出
- Instructions:
  - 需要增强 `Taskfile.yml` 以承载开发运维任务，并统一设置 `silent: true`
  - 本地安装依赖前需先切换 Go 与 npm 国内镜像
  - `task deploy` 需支持通过 CLI 变量传入 `HOST`、`PORT`、`USER`、`DIR` 等远程部署参数
  - 部署逻辑继续通过独立 `sh` 脚本实现，便于后续复用
  - 运维任务中只保留 `task deploy`，不再保留 `prod-deploy` 别名
  - 当前部署方式全面迁移到 Docker 运行，不再保留 `systemd` / `systemctl` 相关脚本、任务和示例

### Git 与回复偏好
- Date: 2026-04-25
- Context: 用户在说明后续协作与提交规范时明确指出
- Instructions:
  - 涉及 `gh` 登录时，可优先尝试复用现有 Git 凭据完成认证
  - commit message 使用专业格式：`改动类型(改动文件): 改动内容`
  - 提交前需一次性阅读所有改动，并按改动分类拆分为多次 commit
  - 提交说明需使用单行 `git commit` 命令，通过多个 `-m` 参数拼接内容
  - 异常处理需尽量完善，并将报错信息显式反馈给用户
  - 聊天回复中的链接不得使用代码块包裹，必须使用纯文本或 Markdown 链接
  - 对于开发任务，开始修改前优先创建分支；若远程与权限可用，则尽早推送分支并发起 PR，后续每完成一个阶段及时推送
  - 若用户要求提交，优先使用中文环境展示 Git 信息

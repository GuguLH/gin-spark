# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## 项目概述

基于 Gin 框架的 Go Web 应用，采用 Clean Architecture 架构，使用 Wire 依赖注入、GORM ORM 和 SQLite 数据库。

## 构建与运行

```bash
# 生成 Wire 依赖注入代码
go generate ./...

# 运行应用（端口 9090）
go run .

# 编译
go build -o app .
```

## 架构设计

```
internal/
├── domain/          # 领域模型（业务实体）
├── repository/      # 仓储接口 + 实现
│   └── dao/        # 数据访问层（GORM 实现）
├── service/         # 业务逻辑层
├── web/             # HTTP 处理器（Gin 控制器）
│   └── middleware/  # Gin 中间件（日志）
└── errs/            # 错误码
ioc/                  # 依赖注入配置（Wire）
pkg/logger/           # 日志抽象（zap 实现）
config/               # 配置文件
```

### 分层依赖

`web` → `service` → `repository` → `dao`

### 核心模式

- **接口驱动设计**：每层定义接口（`IStudentDAO`、`IStudentRepository`、`IStudentService`）
- **Domain 与 DAO 分离**：`domain.Student`（time.Time、string）vs `dao.Student`（int64 时间戳、sql.NullString）
- **错误传播**：DAO 错误（`ErrDuplicatePhone`、`ErrRecordNotFound`）通过 repository 传播至 service
- **路由注册**：处理器实现 `RegisterRoutes(server *gin.Engine)` 方法

## 依赖注入

使用 Wire 进行 DI。修改 `wire.go` 后需重新生成：

```bash
go generate ./...
```

`wire_gen.go` 是自动生成文件，请勿手动编辑。

## 配置

通过 Viper 从 `config/dev.yaml` 加载配置（默认）。可通过 `--config` 参数覆盖：

```bash
go run . --config config/prod.yaml
```

## 代码注释规范

- 注释使用中文，标点使用英文
- 文件头部注释：描述文件用途

```go
// @File    : student.go
// @Author  : GuguLH
// @Date    : 2026/3/26 10:59
// @Desc    : Student 处理器
```

- 仅在逻辑不明确时才添加注释，避免冗余说明
- 导出的函数/类型添加注释说明其用途
- 错误定义集中管理于 `internal/errs/code.go`

## 新增功能

新增实体（如 `Teacher`）的步骤：

1. 创建 `internal/domain/teacher.go` - 领域模型
2. 创建 `internal/repository/dao/teacher.go` - DAO，需实现 `ITeacherDAO` 接口
3. 创建 `internal/repository/teacher.go` - 仓储，需实现 `ITeacherRepository` 接口
4. 创建 `internal/service/teacher.go` - 服务，需实现 `ITeacherService` 接口
5. 创建 `internal/web/teacher.go` - 处理器，需实现 `RegisterRoutes` 方法
6. 更新 `wire.go`，将新的依赖集加入
7. 运行 `go generate ./...` 重新生成 wire_gen.go
8. 在 `ioc/web.go` 的 `InitWebServer` 中注册新路由

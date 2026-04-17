SHELL := /bin/bash

.PHONY: current_dir status push pull \
        init-frontend check-frontend run-frontend \
        init-backend wire jen run-backend test-backend lint-backend tidy-backend swagger \
        dev docker-up docker-down \
        podman-up podman-down podman-mysql-up podman-logs podman-ps podman-clean

current_dir: ## 显示当前目录信息
	@echo "当前目录: $$(pwd)"
	@echo "分支: $$(git rev-parse --abbrev-ref HEAD)"

status: current_dir ## 查看当前变更
	@git status

push: current_dir ## 提交并推送到 main
	@git add . \
	&& (git commit -m "update: $(shell date +'%Y-%m-%d %H:%M:%S')" || exit 0) \
	&& git push origin main

pull: ## 从 main 拉取最新代码
	@git pull origin main

# ============================================
# 开发
# ============================================
dev: ## 同时启动前后端开发服务器
	@echo "启动开发环境..."
	@make -j2 run-backend run-frontend

# ============================================
# 前端
# ============================================
init-frontend: ## 初始化前端依赖
	@cd frontend && npm install
	@echo "前端初始化完成: frontend/"

check-frontend:
	@cd frontend && npm ci && npx tsc --noEmit && npm run build

run-frontend: ## 启动前端开发服务器
	@cd frontend && npm run dev

# ============================================
# 后端
# ============================================
init-backend: ## 初始化后端依赖
	@cd backend && go mod download
	@echo "后端初始化完成: backend/"

wire: ## 重新生成 Wire 依赖注入代码
	@cd backend && wire ./cmd/server/

jen: ## 用 jen 重新生成 DAO 代码（需先编辑 backend/.model_infrax/schema.sql）
	@cd backend && jen

swagger: ## 生成 Swagger API 文档
	@cd backend && swag init -g cmd/server/main.go -o docs/

run-backend: ## 启动后端服务
	@cd backend && go run ./cmd/server --config config/config.yaml

test-backend: ## 运行后端测试
	@cd backend && go test ./... -race -count=1

lint-backend: ## 运行 golangci-lint
	@cd backend && golangci-lint run ./...

tidy-backend: ## 整理后端依赖
	@cd backend && go mod tidy

# ============================================
# Docker
# ============================================
docker-up: ## 启动所有 Docker 容器
	docker compose up -d --build

docker-down: ## 停止所有 Docker 容器
	docker compose down

# ============================================
# Podman（推荐，本仓库默认的容器运行时）
# ============================================
podman-mysql-up: ## 仅启动 mysql 容器（开发阶段 backend 本地跑时用）
	podman compose up -d mysql
	@echo "等待 mysql 变 healthy..."
	@until [ "$$(podman inspect -f '{{.State.Health.Status}}' proj_template_mysql 2>/dev/null)" = "healthy" ]; do sleep 2; done
	@echo "mysql ready on localhost:3306"

podman-up: ## 启动全栈（mysql + backend + frontend），最终验收用
	podman compose up -d --build

podman-down: ## 停止全栈（保留数据卷）
	podman compose down

podman-clean: ## 停止并删除数据卷（彻底清理）
	podman compose down -v

podman-ps: ## 查看容器状态
	podman compose ps

podman-logs: ## 跟随查看所有服务日志（Ctrl+C 退出）
	podman compose logs -f

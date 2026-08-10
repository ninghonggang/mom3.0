.PHONY: help proto proto-sync build run clean docker-up docker-down test smoke

help: ## 显示帮助
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

# ==================== Proto ====================
proto: ## 编译 proto 文件生成 Go 代码
	@echo "Generating proto code..."
	@mkdir -p gen
	@PATH="$$PATH:$$HOME/go/bin" protoc --proto_path=proto \
		--go_out=gen --go_opt=paths=source_relative \
		--go-grpc_out=gen --go-grpc_opt=paths=source_relative \
		proto/common/common.proto proto/common/enums.proto \
		proto/mdm/material.proto proto/mdm/bom.proto proto/mdm/workshop.proto proto/mdm/mdm_service.proto \
		proto/mes/order.proto proto/mes/execution.proto proto/mes/mes_service.proto \
		proto/aps/plan.proto proto/aps/schedule.proto proto/aps/aps_service.proto \
		proto/qms/inspection.proto proto/qms/ncr.proto proto/qms/qms_service.proto \
		proto/eam/equipment.proto proto/eam/eam_service.proto \
		proto/wms/inventory.proto proto/wms/transaction.proto proto/wms/wms_service.proto \
		proto/trace/trace.proto proto/trace/trace_service.proto \
		proto/andon/andon.proto proto/andon/andon_service.proto \
		proto/events/events.proto
	@cd gen && go mod tidy
	@$(MAKE) --no-print-directory proto-sync
	@echo "Proto code generated successfully."

# 网关用 @grpc/proto-loader 在运行时读 .proto，镜像里拷的是 gateway/proto 这份副本。
# 改完 proto 忘了同步，网关就会按旧契约调用（表现为 UNIMPLEMENTED / 字段丢失），
# 所以每次生成代码后自动同步一次。
proto-sync: ## 同步 proto/ 到 gateway/proto/（网关运行时加载用）
	@rsync -a --delete proto/ gateway/proto/
	@echo "proto/ -> gateway/proto/ synced."

# ==================== Go 微服务 ====================
SERVICES = mdm-service mes-service qms-service eam-service wms-service aps-service trace-service andon-service

define build_service
	@echo "=== Building $(1) ==="
	cd services/$(1) && GOWORK=off go build -o bin/$(1) ./cmd/main.go
endef

define run_service
	@echo "=== Starting $(1) ==="
	cd services/$(1) && GOWORK=off go run ./cmd/main.go &
endef

build: ## 编译所有 Go 微服务
	$(foreach svc,$(SERVICES),$(call build_service,$(svc)))

run: ## 本地启动所有微服务（后台）
	$(foreach svc,$(SERVICES),$(call run_service,$(svc)))

clean: ## 清理编译产物
	rm -rf services/*/bin
	rm -rf gen/
	rm -rf dist/

# ==================== Docker ====================
docker-up: ## 启动 Docker 基础设施（仅 DB/NATS/Redis/MinIO）
	docker compose up -d postgres nats redis minio
	@echo "Infrastructure started. Wait for healthy, then run: make docker-up-services"

docker-up-services: ## 启动微服务容器
	docker compose up -d mdm-service mes-service qms-service eam-service wms-service aps-service trace-service andon-service

docker-up-all: ## 启动完整环境（基础设施 + 微服务 + 网关 + 前端）
	docker compose up -d

docker-down: ## 停止所有容器
	docker compose down

# build-linux: 交叉编译所有 Go 微服务为 Linux amd64 二进制
build-linux: ## 编译 Linux amd64 版本（用于 Docker 构建）
	@mkdir -p dist
	$(foreach svc,$(SERVICES),\
		@echo "=== Building $(svc) for linux/amd64 ===" && \
		cd services/$(svc) && GOWORK=off GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
		go build -ldflags="-s -w" -o ../../dist/$(svc) ./cmd/main.go;)

docker-build: build-linux proto-sync ## 构建所有 Docker 镜像（先编译 linux 二进制 + 同步 proto）
	docker compose build

docker-logs: ## 查看所有容器日志
	docker compose logs -f

# ==================== 前端 ====================
web-dev: ## 启动前端开发服务器
	cd web && npm run dev

web-build: ## 构建前端
	cd web && npm run build

gateway-dev: ## 启动网关开发服务器
	cd gateway && npm run start:dev

gateway-build: ## 构建网关
	cd gateway && npm run build

# ==================== 测试 ====================
test: ## 运行所有测试
	$(foreach svc,$(SERVICES),cd services/$(svc) && GOWORK=off go test ./...;)

test-cover: ## 运行测试并生成覆盖率报告
	$(foreach svc,$(SERVICES),cd services/$(svc) && GOWORK=off go test -coverprofile=coverage.out ./...;)

smoke: ## 跑链式端到端写接口冒烟测试（需要环境已启动）
	@./scripts/smoke_write.sh

# ==================== 数据库 ====================
db-migrate: ## 运行所有迁移
	@for svc in $(SERVICES); do \
		echo "Migrating $$svc..."; \
		# psql migration command here \
	done

# ==================== 工具 ====================
fmt: ## 格式化代码
	gofmt -w services/*/.
	cd gateway && npx prettier --write 'src/**/*.ts'
	cd web && npx prettier --write 'src/**/*.{ts,tsx}'

lint: ## 代码检查
	$(foreach svc,$(SERVICES),cd services/$(svc) && GOWORK=off go vet ./...;)
	cd gateway && npx tsc --noEmit
	cd web && npx tsc --noEmit

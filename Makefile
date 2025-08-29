# Variables
BINARY_NAME=koopi-backend
BUILD_DIR=build
MAIN_FILE=cmd/server/main.go

# Colores para output
GREEN=\033[0;32m
NC=\033[0m # No Color

.PHONY: help build run test clean docker-up docker-down migrate-up migrate-down

# Comando por defecto
help: ## Mostrar esta ayuda
	@echo "Comandos disponibles:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "$(GREEN)%-20s$(NC) %s\n", $$1, $$2}'

# Desarrollo
run: ## Ejecutar en modo desarrollo
	@echo "Ejecutando en modo desarrollo..."
	go run $(MAIN_FILE)

build: ## Compilar para producción
	@echo "Compilando..."
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_FILE)

run-prod: build ## Ejecutar en producción
	@echo "Ejecutando en producción..."
	./$(BUILD_DIR)/$(BINARY_NAME)

# Testing
test: ## Ejecutar tests unitarios
	@echo "Ejecutando tests unitarios..."
	go test ./...

test-integration: ## Ejecutar tests de integración
	@echo "Ejecutando tests de integración..."
	go test -tags=integration ./...

test-coverage: ## Ejecutar tests con cobertura
	@echo "Ejecutando tests con cobertura..."
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

# Base de datos y migraciones
migrate-up: ## Ejecutar migraciones
	@echo "Ejecutando migraciones..."
	migrate -path migrations -database "postgres://koopi_user:koopi_password@localhost:5432/koopi_db?sslmode=disable" up

migrate-down: ## Revertir migraciones
	@echo "Revirtiendo migraciones..."
	migrate -path migrations -database "postgres://koopi_user:koopi_password@localhost:5432/koopi_db?sslmode=disable" down

migrate-create: ## Crear nueva migración (uso: make migrate-create name=migration_name)
	@echo "Creando migración: $(name)"
	migrate create -ext sql -dir migrations -seq $(name)

migrate-status: ## Mostrar estado de migraciones
	@echo "Estado de migraciones:"
	migrate -path migrations -database "postgres://koopi_user:koopi_password@localhost:5432/koopi_db?sslmode=disable" version

# Documentación
swagger: ## Generar documentación Swagger
	@echo "Generando documentación Swagger..."
	swag init -g $(MAIN_FILE) -o docs

# Docker
docker-up: ## Levantar servicios Docker
	@echo "Levantando servicios Docker..."
	docker-compose up -d

docker-down: ## Parar servicios Docker
	@echo "Parando servicios Docker..."
	docker-compose down

docker-logs: ## Ver logs de Docker
	@echo "Mostrando logs..."
	docker-compose logs -f

# Utilidades
clean: ## Limpiar archivos generados
	@echo "Limpiando..."
	rm -rf $(BUILD_DIR)
	rm -f coverage.out
	go clean

deps: ## Instalar dependencias
	@echo "Instalando dependencias..."
	go mod download
	go mod tidy

install-tools: ## Instalar herramientas necesarias
	@echo "Instalando herramientas..."
	go install github.com/swaggo/swag/cmd/swag@latest
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Setup inicial
setup: install-tools deps ## Configurar proyecto inicialmente
	@echo "Configuración inicial completada"

# Desarrollo con hot reload (requiere air)
dev: ## Ejecutar con hot reload (requiere air instalado)
	@echo "Ejecutando con hot reload..."
	air 
.PHONY: run-task-1 run-task-2 run-task-3 run-task-4 all clean build

# Colors for beautiful output
GREEN=\033[0;32m
BLUE=\033[0;34m
YELLOW=\033[1;33m
RED=\033[0;31m
NC=\033[0m # No Color
BOLD=\033[1m

run-task-1:
	@echo "$(BLUE)$(BOLD)========================================$(NC)"
	@echo "$(GREEN)$(BOLD)🚀 Starting Task-1$(NC)"
	@echo "$(BLUE)$(BOLD)========================================$(NC)"
	@go run Task-1/*.go
	@echo "$(GREEN)$(BOLD)✅ Task-1 completed successfully!$(NC)"
	@echo ""

run-task-2:
	@echo "$(BLUE)$(BOLD)========================================$(NC)"
	@echo "$(GREEN)$(BOLD)🚀 Starting Task-2$(NC)"
	@echo "$(BLUE)$(BOLD)========================================$(NC)"
	@if [ -f Task-2/main.go ]; then \
		go run Task-2/*.go; \
		echo "$(GREEN)$(BOLD)✅ Task-2 completed successfully!$(NC)"; \
	else \
		echo "$(YELLOW)$(BOLD)⚠️  Task-2: No Go files found, skipping...$(NC)"; \
	fi
	@echo ""

run-task-3:
	@echo "$(BLUE)$(BOLD)========================================$(NC)"
	@echo "$(GREEN)$(BOLD)🚀 Starting Task-3$(NC)"
	@echo "$(BLUE)$(BOLD)========================================$(NC)"
	@if [ -f Task-3/main.go ]; then \
		go run Task-3/*.go; \
		echo "$(GREEN)$(BOLD)✅ Task-3 completed successfully!$(NC)"; \
	else \
		echo "$(YELLOW)$(BOLD)⚠️  Task-3: No Go files found, skipping...$(NC)"; \
	fi
	@echo ""

run-task-4:
	@echo "$(BLUE)$(BOLD)========================================$(NC)"
	@echo "$(GREEN)$(BOLD)🚀 Starting Task-4$(NC)"
	@echo "$(BLUE)$(BOLD)========================================$(NC)"
	@if [ -f Task-4/main.go ]; then \
		go run Task-4/*.go; \
		echo "$(GREEN)$(BOLD)✅ Task-4 completed successfully!$(NC)"; \
	else \
		echo "$(YELLOW)$(BOLD)⚠️  Task-4: No Go files found, skipping...$(NC)"; \
	fi
	@echo ""

all:
	@echo "$(BLUE)$(BOLD)================================================$(NC)"
	@echo "$(GREEN)$(BOLD)🎯 Running All Tasks Sequentially$(NC)"
	@echo "$(BLUE)$(BOLD)================================================$(NC)"
	@echo ""
	@$(MAKE) run-task-1
	@echo "$(YELLOW)⏳ Waiting 2 seconds...$(NC)"
	@sleep 2
	@$(MAKE) run-task-2
	@echo "$(YELLOW)⏳ Waiting 2 seconds...$(NC)"
	@sleep 2
	@$(MAKE) run-task-3
	@echo "$(YELLOW)⏳ Waiting 2 seconds...$(NC)"
	@sleep 2
	@$(MAKE) run-task-4
	@echo "$(BLUE)$(BOLD)================================================$(NC)"
	@echo "$(GREEN)$(BOLD)🎉 All Tasks Completed!$(NC)"
	@echo "$(BLUE)$(BOLD)================================================$(NC)"

clean:
	@echo "$(YELLOW)$(BOLD)🧹 Cleaning build artifacts...$(NC)"
	@go clean ./...
	@echo "$(GREEN)$(BOLD)✅ Clean completed!$(NC)"

build:
	@echo "$(BLUE)$(BOLD)🔨 Building all tasks...$(NC)"
	@for task in Task-1 Task-2 Task-3 Task-4; do \
		if [ -f $$task/main.go ]; then \
			echo "$(GREEN)Building $$task...$(NC)"; \
			cd $$task && go build -o ../bin/$$task . && cd ..; \
		fi; \
	done
	@echo "$(GREEN)$(BOLD)✅ Build completed!$(NC)"
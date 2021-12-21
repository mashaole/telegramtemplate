tidy: ## Clean up dependencies
	go install
	go mod tidy
	go mod vendor

# NSFAS-TELEGRAM
#===============================================
start: ## Run USSD Telegram Bot
	go run ./main.go

# TESTS
#===============================================
test-all: ## Run all tests in project
	go test

#===============================================
.DEFAULT_GOAL := list
.PHONY: list
list: ## List all make targets (default target)
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
	| sed 's/://g' \
	| sort \
	| column -ts '##'

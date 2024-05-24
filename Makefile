.DEFAULT_GOAL := help

# INFO: 参考サイト - https://postd.cc/auto-documented-makefile/
help: ## subcommand list and description.
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
	| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

PHONY := build
build: ## go build
	@go build -o dist/gomver

PHONY := clean
clean: ## rm -rf dist
	@rm -rf dist

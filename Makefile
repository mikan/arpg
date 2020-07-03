.DEFAULT_GOAL := help

.PHONY: setup
setup: ## Resolve dependencies using Go Modules
	go mod download

.PHONY: clean
clean: ## Remove build artifact directory
	-rm -rfv arping-gui*

.PHONY: lint
lint: ## Run static code analysis
	command -v golint >/dev/null 2>&1 || { go get -u golang.org/x/lint/golint; }
	golint -set_exit_status ./...

.PHONY: run
run: ## Run app locally
	go run . -d

.PHONY: build-linux
build-linux: ## Build linux package
	command -v fyne >/dev/null 2>&1 || { go get -u fyne.io/fyne/cmd/fyne; }
	fyne package -os linux -icon icon.png -release

.PHONY: build-mac
build-mac: ## Build mac package
	command -v fyne >/dev/null 2>&1 || { go get -u fyne.io/fyne/cmd/fyne; }
	fyne package -os darwin -icon icon.png -release
	zip -r arping-gui_macos.zip arping-gui.app

.PHONY: build-win
build-win: ## Build windows package
	command -v fyne >/dev/null 2>&1 || { go get -u fyne.io/fyne/cmd/fyne; }
	fyne package -os windows -icon icon.png -release

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

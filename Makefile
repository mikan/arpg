.DEFAULT_GOAL := help

.PHONY: setup
setup: ## Resolve dependencies using Go Modules
	go mod download

.PHONY: clean
clean: ## Remove build artifact directory
	-rm -rfv arpg*

.PHONY: lint
lint: ## Run static code analysis
	command -v golint >/dev/null 2>&1 || { go install golang.org/x/lint/golint@latest; }
	golint -set_exit_status ./...

.PHONY: run
run: ## Run app locally
	go run . -d

.PHONY: build-linux
build-linux: ## Build linux package
	command -v fyne >/dev/null 2>&1 || { go install fyne.io/fyne/v2/cmd/fyne@latest; }
	fyne package -os linux -icon icon.png -release -appID com.github.mikan.arpg

.PHONY: build-mac
build-mac: ## Build mac package
	command -v fyne >/dev/null 2>&1 || { go install fyne.io/fyne/v2/cmd/fyne@latest; }
	fyne package -os darwin -icon icon.png -release -appID com.github.mikan.arpg
	zip -r arpg_macos.zip arpg.app

.PHONY: build-win
build-win: ## Build windows package
	if not exist %GOPATH%\bin\fyne go install fyne.io/fyne/v2/cmd/fyne@latest
	fyne package -os windows -icon icon.png -release -appID com.github.mikan.arpg

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

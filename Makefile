#!make
TAG=$(shell git describe --tags |cut -d- -f1)
COMMIT=$(shell git rev-parse --short HEAD)
LDFLAGS=-a -installsuffix cgo

.SILENT:
.DEFAULT_GOAL := help

help: ## Show this help
	@echo "Makefile available targets:"
	@grep -h -E '^[a-zA-Z_-].+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf " * \033[36m%-15s\033[0m %s\n", $$1, $$2}'

test: ## Run tests
	go test --short -coverprofile=cover.out -v ./...
	make test.coverage

test.coverage:
	go tool cover -func=cover.out

release: ## Git tag create and push
	git tag -s -a v${tag} -m 'chore(release): v$(tag) [skip ci]'
	git push origin v${tag}

release.revert: ## Revert git release tag
	git tag -d v${tag}
	git push --delete origin v${tag}

lint: ## Check code (used golangci-lint)
	GO111MODULE=on golangci-lint run

clean: ## Clean build directory
	rm -rf ./.bin
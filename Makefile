IS_IN_PROGRESS = "is in progress ..."

## gen: will generate mock for usecases & repositories interfaces
.PHONY: gen
gen:
	@echo "make gen ${IS_IN_PROGRESS}"
	@mockgen -source internal/users/usecase.go -destination internal/generated/mocks/usecase_mock.go -package=mocks
	@mockgen -source internal/users/repository.go -destination internal/generated/mocks/repository_mock.go -package=mocks

## setup: Set up database temporary for integration testing
.PHONY: setup
setup:
	@docker-compose -f ./infrastructure/docker-compose.yml up -d
	@sleep 10

## down: Set down database temporary for integration testing
.PHONY: down
down: 
	@docker-compose -f ./infrastructure/docker-compose.yml down -t 1

## integration-test: will test with integration tags
.PHONY: integration-test
integration-test:
	@echo "make integration-test ${IS_IN_PROGRESS}"
	@go clean -testcache
	@go test --race -timeout=90s -failfast \
		-vet= -cover -covermode=atomic -coverprofile=./.coverage/integration.out \
		-tags=integration ./internal/users/repository/...\

## unit-test: will test with unit tags
.PHONY: unit-test
unit-test:
	@echo "make unit-test ${IS_IN_PROGRESS}"
	@go clean -testcache
	@go test --race -timeout=90s -failfast \
		-vet= -cover -covermode=atomic -coverprofile=./.coverage/unit.out \
		-tags=unit ./internal/users/usecase/...\

## e2e-test: will test with e2e tags
.PHONY: e2e-test
e2e-test:
	@echo "make e2e-test ${IS_IN_PROGRESS}"
	@go clean -testcache
	@go test --race -timeout=90s -failfast \
		-vet= -cover -covermode=atomic -coverprofile=./.coverage/e2e.out \
		-tags=e2e ./internal/users/delivery/...\

## tests: run tests and any dependencies
.PHONY: tests
tests: setup e2e-test unit-test integration-test down
.PHONY: run
run: migrate
	go run -v ./cmd/apiserver

.PHONY: dev
dev: migrate
	CompileDaemon --build="go build -v ./cmd/apiserver" --command="./apiserver" --exclude-dir="frontend" --exclude-dir=".git" -include="./swagger.yaml"

.PHONY: debug
debug: build
	dlv --listen=:2345 --headless --api-version=2 --accept-multiclient --log exec ./apiserver

.PHONY: build
build: migrate
	go build -gcflags="all=-N -l" -v ./cmd/apiserver

.PHONY: migrate
migrate:
	migrate -path=./migrations/ -database "${DB_URL}" up

.PHONY: test
test:
	echo "SELECT 'CREATE DATABASE auth_test' WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'my_trello_test')\gexec" | psql "${DB_URL}"
	migrate -path=./migrations/ -database "${DB_TEST_URL}" up
	go test -v -race -timeout 30s ./...

.PHONY: check_install
check_install:
	which swagger || GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger

.PHONY: swagger
swagger: check_install
	GO111MODULE=off swagger generate spec -o ./swagger.yaml --scan-models

.PHONY: hooks
hooks:
	pre-commit run --all-files

.DEFAULT_GOAL := run

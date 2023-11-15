.PHONY: build local-db proto swagger run gen-mock unit-test open-coverage gci lint lint-consisten sec

build:
	go build -o ./tmp/server ./cmd/main.go

local-db:
	@docker-compose down
	@docker-compose up -d

proto:
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/teq.proto
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative cache/database/database.proto

swagger:
	@@hash swag 2>/dev/null || GO111MODULE=off go get -u github.com/swaggo/swag/cmd/swag
	swag init -g cmd/main.go --parseDependency --parseInternal --parseDepth 2

run:
	# go install github.com/cosmtrek/air@latest
	@air -c .air.toml

gen-mock:
	@mockery --inpackage --with-expecter --name=Repository --dir=./repository/example
	@mockery --inpackage --with-expecter --name=Repository --dir=./repository/account
	@mockery --inpackage --with-expecter --name=Repository --dir=./repository/product
	@mockery --inpackage --with-expecter --name=Repository --dir=./repository/producer
	@mockery --inpackage --with-expecter --name=Repository --dir=./repository/storage
	@mockery --inpackage --with-expecter --name=Repository --dir=./repository/checkout
	
	@mockery --inpackage --with-expecter --name=ICache --dir=./cache

unit-test:
	@mkdir coverage || true
	@godotenv -f .env go test -race -v -coverprofile=coverage/coverage.txt.tmp -count=1  ./...
	@cat coverage/coverage.txt.tmp | grep -v "mock_" > coverage/coverage.txt
	@go tool cover -func=coverage/coverage.txt
	@go tool cover -html=coverage/coverage.txt -o coverage/index.html

open-coverage:
	@open coverage/index.html

gci:
	@gci write -s Standard -s Default -s "Prefix(github.com/teq-quocbang/store)" .

lint:
	@hash golangci-lint 2>/dev/null || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.48.0
	@GO111MODULE=on CGO_ENABLED=0 golangci-lint run

lint-consistent:
	@hash go-consistent 2>/dev/null || GO111MODULE=off go get -v github.com/quasilyte/go-consistent
	@go-consistent ./...

sec:
	@curl -sfL https://raw.githubusercontent.com/securego/gosec/master/install.sh | sh -s -- -b $(GOPATH)/bin latest
	@gosec ./...

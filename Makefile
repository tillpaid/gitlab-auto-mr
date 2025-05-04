MAIN=cmd/main.go
BIN=bin/app
GLOBAL_BIN=/usr/local/bin/gitlab-auto-mr

run:
	@clear
	@go run $(MAIN)

build:
	@go build -o $(BIN) $(MAIN)
	@echo "Binary built at $(BIN)"

test:
	clear
	go test ./internal/...

testf:
	clear
	go test ./internal/... -failfast

testfw:
	clear
	gow -c test ./internal/... -failfast

build-and-replace:
	@go build -o $(BIN) $(MAIN) && sudo cp $(BIN) $(GLOBAL_BIN)
	@echo "Built and copied into $(GLOBAL_BIN)"

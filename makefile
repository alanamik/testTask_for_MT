BIN := "./encryption"
DOCKER_IMG="encryption:latest"

build:
	go build -o $(BIN) ./cmd/main.go

run: 
	build 
	$(BIN)

test:
	go test -timeout=90s -count=1 -v ./internal/...

docker-build: build
	docker build -t $(DOCKER_IMG) .

docker-run:
	docker compose up -d

docker-down:
	docker compose down

lint: 
	golangci-lint run ./...

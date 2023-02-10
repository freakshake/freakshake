MIGRATIONS_PATH=internal/schemas

POSTGRES_DB_NAME=test_db
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=1234
POSTGRES_DB_URL=postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:$(POSTGRES_PORT)/$(DB_NAME)?sslmode=disable

POSTGRES_CONTAINER=postgres_test
REDIS_CONTAINER=redis_test

tidy:
	go mod tidy

fmt:
	go fmt ./...
	swag fmt

lint:
	golangci-lint run --config "./config/.golangci.yaml"

doc:
	swag i --pd 

build-http:
	go build -o medad ./cmd/service/

run-http:
	./medad

postgres-container:
	docker run --name $(POSTGRES_CONTAINER) -p $(POSTGRES_PORT):5432 -e POSTGRES_USER=$(POSTGRES_USER) -e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) -d postgres:alpine

redis-container:
	docker run --name $(REDIS_CONTAINER) -p 6379:6379 -e REDIS_PASSWORD=redis -d redis:alpine

start-containers:
	docker start $(POSTGRES_CONTAINER)
	docker start $(REDIS_CONTAINER)

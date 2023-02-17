MIGRATIONS_PATH=./migrations/

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

swagger:
	swag init --dir ./cmd/service/ --output ./api/swagger --parseDependency 

build-http:
	go build -o storm ./cmd/service/

run-http:
	./storm

postgres-container:
	docker run --name $(POSTGRES_CONTAINER) -p $(POSTGRES_PORT):5432 -e POSTGRES_USER=$(POSTGRES_USER) -e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) -d postgres:alpine

redis-container:
	docker run --name $(REDIS_CONTAINER) -p 6379:6379 -e REDIS_PASSWORD=redis -d redis:alpine

start-containers:
	docker start $(POSTGRES_CONTAINER)
	docker start $(REDIS_CONTAINER)

stop-containers:
	docker stop $(POSTGRES_CONTAINER)
	docker stop $(REDIS_CONTAINER)

create-postgres-db:
	docker exec -it $(POSTGRES_CONTAINER) createdb --username=postgres --owner=postgres $(POSTGRES_DB_NAME)

drop-postgres-db:
	docker exec -it $(POSTGRES_CONTAINER) dropdb --username=postgres $(POSTGRES_DB_NAME)

migrateup:
	migrate -path "$(MIGRATIONS_PATH)" -database "$(POSTGRES_DB_URL)" -verbose up

migrateup-1:
	migrate -path "$(MIGRATIONS_PATH)" -database "$(POSTGRES_DB_URL)" -verbose up 1

migratedown:
	migrate -path "$(MIGRATIONS_PATH)" -database "$(POSTGRES_DB_URL)" -verbose down

migratedown-1:
	migrate -path "$(MIGRATIONS_PATH)" -database "$(POSTGRES_DB_URL)" -verbose down 1

create-migration:
	migrate create -ext sql -dir "$(MIGRATIONS_PATH)" -seq "$(MIGRATION_NAME)"

fix-migrate:
	migrate -database "$(POSTGRES_DB_URL)" -path "$(MIGRATIONS_PATH)" force $(VERSION)

pipeline: tidy swagger fmt build-http run-http
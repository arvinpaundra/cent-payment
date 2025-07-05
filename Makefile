APP_NAME := cent-payment
REST_PORT ?= 8090
GRPC_PORT ?= 8093

DB_URL ?= postgres://root:root@localhost:5432/cent_payment?sslmode=disable
MIGRATION_PATH := ./migrations

build:
	@echo "Building $(APP_NAME)"
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./bin/$(APP_NAME) .

rest:
	@echo "Running REST on $(APP_NAME)" 
	go run main.go rest -p $(REST_PORT)

grpc:
	@echo "Running gRPC on $(APP_NAME)"
	go run main.go grpc -p $(GRPC_PORT)

poller:
	@echo "Running poller on $(APP_NAME)"
	go run main.go poller

test:
	@echo "Running tests on $(APP_NAME)"
	go test -v -cover ./...

cleanup:
	@echo "removing /bin"
	rm -rf bin/

migrateadd:
	@echo "Adding new migration file $(NAME)"
	migrate create -ext sql -dir $(MIGRATION_PATH) $(NAME)

migrateup:
	@echo "Execute migrate up"
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" -verbose up

migratedown:
	@echo "Execute migrate down"
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" -verbose down

migratereset:
	@echo "Execute reset all migrations"
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" -verbose drop -f

migraterefresh: migratedown migrateup

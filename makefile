include infrastructure/local/.env
APIV1_PATH:=./api/v1

POSTGRES_URI = "postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB_NAME)?sslmode=$(POSTGRES_DB_SSL_MODE)"
GOOSE_RUN := go run github.com/pressly/goose/cmd/goose@latest
JET_RUN:= go run github.com/go-jet/jet/v2/cmd/jet@latest

.PHONY: gen-api
gen-api: 
	mkdir -p $(APIV1_PATH) 
	protoc -I $(APIV1_PATH) --go_out=. --go-grpc_out=require_unimplemented_servers=true:. $(APIV1_PATH)/*.proto

	.PHONY:up
up:
	docker compose -f ./infrastructure/local/docker-compose.yaml up

.PHONY:down
down:
	docker compose -f ./infrastructure/local/docker-compose.yaml down --remove-orphans --volumes

.PHONY: add-migration
add-migration:
	$(GOOSE_RUN) -dir db/migrations create new sql

.PHONY:fixtures-up
fixtures-up:
	$(GOOSE_RUN) -dir db/fixtures -table goose_db_fixtures_version postgres $(POSTGRES_URI) up

.PHONY:gen-db
gen-db:
	$(JET_RUN) -source postgres -dsn=$(POSTGRES_URI) -ignore-tables goose_db_version,goose_db_fixtures_version -path=db/gen

.PHONY:migrate-up
migrate-up:
	$(GOOSE_RUN) -dir db/migrations -table goose_db_version postgres $(POSTGRES_URI) up

.PHONY:start-server
start-server:
	godotenv -f infrastructure/local/.env go run cmd/server/main.go

.PHONY:start-client
start-client:
	go run cmd/client/main.go

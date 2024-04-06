.PHONY: build
build:
	@go build -o bin/migrator main.go

.PHONY: migrate-up
migrate-up:
	@./bin/migrator

.PHONY: postgres-init
postgres-init:
	@docker run --name postgres-test -p 5433:5432 -d --rm \
		-e POSTGRES_PASSWORD=foobar -e POSTGRES_USER=ptnshk postgres

.PHONY: postgres-cli
postgres-cli:
	@docker exec -it postgres-test psql -U ptnshk
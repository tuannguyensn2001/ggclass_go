install-tools:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

migrate-create:
	migrate create -ext sql -dir src/database/postgres -seq ${name}

migrate:
	@go run src/server/main.go migrate-up

migrate-down:
	@go run src/server/main.go migrate-down

migrate-refresh:
	@go run src/server/main.go migrate-refresh

build:
	@go build src/server/main.go

gen-error:
	@go run src/server/main.go gen-error ${parent} ${child} "${message}" ${statusCode}
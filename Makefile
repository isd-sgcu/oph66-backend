fmt:
	gofmt -w **/*.go

wire:
	go run github.com/google/wire/cmd/wire@latest ./...

mig:
	atlas schema apply \
		--url "postgres://postgres:123456@127.0.0.1:5432/postgres?&sslmode=disable" \
		--to "file://./migrations/init.sql" \
		--dev-url "docker://postgres/15"

start:
	. ./export-env.sh ; go run cmd/main.go

dev:
	. ./export-env.sh ; nodemon --exec go run cmd/main.go --signal SIGTERM

seed:
	./seed.sh

lint:
	golangci-lint run

lint-fix:
	golangci-lint run --fix

swag:
	swag init -g ./cmd/main.go -o ./docs

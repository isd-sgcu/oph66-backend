fmt:
	gofmt -w **/*.go

wire:
	go run github.com/google/wire/cmd/wire@latest ./...

up:
	docker-compose up -dm

run:
	. ./export-env.sh ; go run cmd/main.go

dev:
	. ./export-env.sh ; nodemon --exec go run cmd/main.go --signal SIGTERM
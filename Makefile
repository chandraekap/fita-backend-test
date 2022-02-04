build:
	go build -o .bin/main .

exec:
	./.bin/main

run:
	go run server.go wire_gen.go

test:
	go test -cover ./...

wire:
	wire gen
	wire gen ./internal/sales/test

gql:
	go run github.com/99designs/gqlgen generate
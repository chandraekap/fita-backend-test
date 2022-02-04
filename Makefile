build:
	go build -o .bin/main .

exec:
	./.bin/main

run:
	go run server.go wire_gen.go

test:
	go test -cover ./...
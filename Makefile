build:
	go build -o app cmd/server/main.go

test:
	go test -v ./...

lint:
	golangci-lint run

run:
	docker-compose up --build
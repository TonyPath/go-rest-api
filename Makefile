build:
	go build -o app cmd/server/main.go

test:
	go test -v ./...

lint:
	golangci-lint run

run:
	docker-compose up --build

export DB_USERNAME=postgres
export DB_PASSWORD=postgres
export DB_DBNAME=comments_db
export DB_HOST=db
export SSL_MODE=disable

int-test:
	docker-compose up -d db
	go test -tags=integration -v ./...
	docker-compose down db

e2e-tests:
	docker-compose up -d --build
	go test -tags=e2e -v ./...
	docker-compose down
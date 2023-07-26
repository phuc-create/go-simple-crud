run:
		go run main.go
build:
		go build main.go
test:
		go test ./...
migrate-d:
	migrate -source file://data/migrations -database "postgres://postgres:1@localhost:5432/simple_crud?sslmode=disable" -verbose down
#3
migrate-u:
	migrate -source file://data/migrations -database "postgres://postgres:1@localhost:5432/simple_crud?sslmode=disable" -verbose up

migrate-f:
	migrate -source file://data/migrations -database "postgres://postgres:1@localhost:5432/simple_crud?sslmode=disable" -verbose force 0

redo-db: migrate-f migrate-u
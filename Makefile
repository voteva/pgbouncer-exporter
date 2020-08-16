build:
	mkdir -p bin
	go build -o bin/pgbouncer-exporter ./pgbouncer-exporter.go

run:
	go run ./pgbouncer-exporter.go

test:
	go test ./...

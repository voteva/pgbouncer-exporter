build:
	mkdir -p bin
	go build -o bin/pgbouncer-exporter pgbouncer-exporter

clean:
	rm -rf bin

run:
	go run pgbouncer-exporter

test:
	go test ./...

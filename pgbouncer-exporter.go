package main

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	metricsHost = "0.0.0.0"
	metricsPort = 8989
	metricsPath = "/metrics"
	healthzPath = "/healthz"
	namespace   = "pgbouncer"
	indexHTML   = `
	<html>
		<head><title>PgBouncer Metrics Exporter</title></head>
		<body>
			<h1>PgBouncer Metrics Exporter</h1>
			<ul>
				<li><a href='` + metricsPath + `'>metrics</a></li>
				<li><a href='` + healthzPath + `'>healthz</a></li>
			</ul>
		</body>
	</html>`
)

func main() {
	// Connect to pgbouncer
	db, err := Connect("postgres://postgres:@localhost:6543/pgbouncer?sslmode=disable")
	if err != nil {
		log.Fatal("Failed to connect to pgbouncer", err)
	}

	// Create new collector
	collector := NewCollector(db, namespace)
	defer collector.Close()

	// Register collector
	r := prometheus.NewRegistry()
	r.MustRegister(collector)

	listenAddress := net.JoinHostPort(metricsHost, fmt.Sprint(metricsPort))
	mux := http.NewServeMux()
	// Add metricsPath
	mux.Handle(metricsPath, promhttp.HandlerFor(r, promhttp.HandlerOpts{}))
	// Add healthzPath
	mux.HandleFunc(healthzPath, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		if _, err := w.Write([]byte("ok")); err != nil {
			log.Fatal(err, "Unable to write to serve metrics")
		}
	})
	// Add index
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(indexHTML))
		if err != nil {
			log.Fatal(err, "Unable to write to serve metrics")
		}
	})
	err = http.ListenAndServe(listenAddress, mux)
	log.Fatal(err, "Failed to serve metrics")
}

func Connect(conn string) (*sql.DB, error) {
	connector, err := pq.NewConnector(conn)
	if err != nil {
		return nil, err
	}
	db := sql.OpenDB(connector)

	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)

	return db, nil
}

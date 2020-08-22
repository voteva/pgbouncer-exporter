package main

import (
	"database/sql"
	"github.com/lib/pq"
)

type SQLDB interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Close() error
}

func Connect(conn string) (SQLDB, error) {
	connector, err := pq.NewConnector(conn)
	if err != nil {
		return nil, err
	}
	db := sql.OpenDB(connector)

	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)

	return db, nil
}

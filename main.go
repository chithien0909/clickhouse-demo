package main

import (
	"database/sql"
	_ "github.com/mailru/go-clickhouse/v2"
	"log"
)

func main() {
	conn, err := connect()
	if err != nil {
		panic(err)
	}

	_ = conn
}

func connect() (*sql.DB, error) {
	conn, err := sql.Open("chhttp", "http://127.0.0.1:8123/default")
	if err != nil {
		log.Fatal(err)
	}
	if err := conn.Ping(); err != nil {
		log.Fatal(err)
	}

	tx, err := conn.Begin()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := conn.Query(`
		SELECT * FROM
			postgres_db.region`)

	if err != nil {
		log.Fatal(err)
	}

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var regionId int
		var regionDes string
		if err := rows.Scan(&regionId, &regionDes); err != nil {
			log.Fatal(err)
		}
		log.Printf("regionId: %d, regionDes: %s", regionId, regionDes)
	}

	return conn, nil
}

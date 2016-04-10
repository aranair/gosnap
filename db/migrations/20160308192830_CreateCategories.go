package main

import (
	"database/sql"

	_ "bitbucket.org/liamstask/goose/cmd/goose"
)

// Up is executed when this migration is applied
func Up_20160308192830(txn *sql.Tx) {
	txn.Exec(`
    CREATE TABLE categories (
      id serial NOT NULL,
      name text NOT NULL,
      url_key text NOT NULL UNIQUE,
      url text NOT NULL UNIQUE,
      created_at timestamp without time zone,
      updated_at timestamp without time zone,
      deleted_at timestamp without time zone,
      PRIMARY KEY(id)
    );
  `)
}

// Down is executed when this migration is rolled back
func Down_20160308192830(txn *sql.Tx) {
	txn.Exec(`DROP TABLE categories;`)
}

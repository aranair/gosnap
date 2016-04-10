package main

import (
	"database/sql"

	_ "bitbucket.org/liamstask/goose/cmd/goose"
)

// Up is executed when this migration is applied
func Up_20160309214641(txn *sql.Tx) {
	txn.Exec(`
    CREATE TABLE listings (
      id serial NOT NULL,
      category_id integer REFERENCES categories(id),
      url text NOT NULL UNIQUE,
      title text NOT NULL,
      created_at timestamp without time zone,
      updated_at timestamp without time zone,
      deleted_at timestamp without time zone,
      PRIMARY KEY(id)
    );
  `)
}

// Down is executed when this migration is rolled back
func Down_20160309214641(txn *sql.Tx) {
	txn.Exec(`DROP TABLE listings;`)
}

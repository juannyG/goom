package product

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const dbName = "./test.db"

const merchantTableSql = `
CREATE TABLE merchant_merchant (
    "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
    "public_id" TEXT
);`

const productTableSql = `
CREATE TABLE product_product (
    "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
    "external_product_id" TEXT,
    "autoship_enabled" INTEGER,
    "live" INTEGER,
    "merchant_id" INTEGER,
    FOREIGN KEY(merchant_id) REFERENCES merchant_merchant(id)
);
`

const merchantSql = `INSERT INTO merchant_merchant (public_id) VALUES ("TestM");`
const productSql = `
INSERT INTO product_product
    (merchant_id, external_product_id, autoship_enabled, live)
    VALUES (1, "TestP", 1, 1);
`

type DBWrapper struct {
	DB *sql.DB
}

func (dw *DBWrapper) InitDB() (*sql.DB) {
	var err error
	os.Remove(dbName)

	dw.DB, err = sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal(err.Error())
	}

	dw.createTable()
	dw.createData()
	return dw.DB
}

func (dw *DBWrapper) createTable() {
	tables := []string{merchantTableSql, productTableSql}
	for _, table := range tables {
		stmt, err := dw.DB.Prepare(table)
		if err != nil {
			log.Fatal(err.Error())
		}
		stmt.Exec()
	}
}

func (dw *DBWrapper) createData() {
	rows := []string{merchantSql, productSql}
	for _, row := range rows {
		stmt, err := dw.DB.Prepare(row)
		if err != nil {
			log.Fatal(err.Error())
		}
		_, err = stmt.Exec()
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}

func (dw *DBWrapper) tearDown() {
	os.Remove(dbName)
}

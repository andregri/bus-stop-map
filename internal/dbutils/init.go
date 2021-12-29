package dbutils

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func CreateTable(table string) {
	stmt, err := DB.Prepare(table)
	if err != nil {
		log.Fatalf("Error in table statement: %s\n", err)
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Fatalf("Error in table creation: %s\n", err)
	}
}

func InitTables() {
	CreateTable(ArrivalTable)
	log.Println("Tables created")
}

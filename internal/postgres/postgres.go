package postgres

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/andregri/bus-stop-map/internal/record"
)

const (
	Host = "localhost"
	Port = 5432
	User = "postgres"
	Password = "123"
	Dbname = "postgres"
)

type ArrivalTimeDb struct {
	Db *sql.DB
	TableName string
}

func (db *ArrivalTimeDb) MakeArrivalTimeTable() {
	stmt, err := db.Db.Prepare("create table IF NOT EXISTS students(id serial PRIMARY KEY,name VARCHAR ( 50 ) NOT NULL,roll INT NOT NULL)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		panic(err)
	}
}

func (db *ArrivalTimeDb) CreateRecord(ctx context.Context, record record.ArrivalTimeRecord) {
	// Insert data
	insertDynStmt := `insert into "students"("name", "roll") values($1, $2)`
	_, err := db.Db.Exec(insertDynStmt, "John", 2)
	if err != nil {
		panic(err)
	}
}
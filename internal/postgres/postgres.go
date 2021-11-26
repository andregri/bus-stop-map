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
	Sql *sql.DB
	TableName string
}

func (db *ArrivalTimeDb) MakeArrivalTimeTable() {
	stmt, err := db.Sql.Prepare(
		`create table IF NOT EXISTS arrival_time(
			id serial PRIMARY KEY,
			stop_code VARCHAR ( 50 ) NOT NULL,
			bus_line VARCHAR ( 50 ) NOT NULL,
			arrival_time TIMESTAMP NOT NULL)`)
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
	insertStmt := `insert into "arrival_time"("stop_code", "bus_line",
		"arrival_time") values($1, $2, $3)`
	_, err := db.Sql.Exec(insertStmt, record.StopCode, record.BusLine, record.ArrivalTime)
	if err != nil {
		panic(err)
	}
}

func (db *ArrivalTimeDb) DeleteRecord(ctx context.Context, id int) error {
	panic("Not implemented")
}
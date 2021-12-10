package postgres

import (
	"context"
	"database/sql"

	"github.com/andregri/bus-stop-map/internal/record"
	_ "github.com/lib/pq"
)

type ArrivalTimeDb struct {
	Sql       *sql.DB
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

func (db *ArrivalTimeDb) CreateRecord(ctx context.Context, record record.ArrivalTimeRecord) error {
	// Insert data
	insertStmt := `insert into "arrival_time"("stop_code", "bus_line",
		"arrival_time") values($1, $2, $3)`
	_, err := db.Sql.Exec(insertStmt, record.StopCode, record.BusLine, record.ArrivalTime)
	if err != nil {
		panic(err)
	}

	return nil
}

func (db *ArrivalTimeDb) SearchRecord(ctx context.Context, stopCode string) ([]record.ArrivalTimeRecord, error) {
	searchStmt := `select * from arrival_time where stop_code = $1`
	rows, err := db.Sql.Query(searchStmt, stopCode)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	records := []record.ArrivalTimeRecord{}

	for rows.Next() {
		var _id int
		var rec record.ArrivalTimeRecord
		if err := rows.Scan(&_id, &rec.StopCode, &rec.BusLine, &rec.ArrivalTime); err != nil {
			return nil, err
		}
		records = append(records, rec)
	}

	return records, nil
}

func (db *ArrivalTimeDb) DeleteRecord(ctx context.Context, id int) error {
	deleteStmt := `delete from arrival_time where id = $1`
	_, err := db.Sql.Exec(deleteStmt, id)
	if err != nil {
		panic(err)
	}

	return nil
}

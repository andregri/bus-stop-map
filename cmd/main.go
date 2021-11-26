package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/andregri/bus-stop-map/internal/postgres"
	"github.com/andregri/bus-stop-map/internal/record"
)

func main() {
	psqlconn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		postgres.Host, postgres.User, postgres.Password, postgres.Dbname)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("connected")

	atdb := &postgres.ArrivalTimeDb{Db: db, TableName: "arrival_time"}

	atdb.MakeArrivalTimeTable()
	atdb.CreateRecord(nil, record.ArrivalTimeRecord{StopCode: "asdf", BusLine: "sdf", ArrivalTime: "asdf"})
}
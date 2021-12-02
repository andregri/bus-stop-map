package main

import (
	"context"
	"crypto/tls"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/andregri/bus-stop-map/internal/api"
	"github.com/andregri/bus-stop-map/internal/postgres"
	"github.com/andregri/bus-stop-map/internal/record"
	_ "github.com/lib/pq"
)

const (
	Host     = "localhost"
	Port     = 5432
	User     = "postgres"
	Password = "123"
	Dbname   = "postgres"
)

func main() {
	// Connect to sql server
	psqlconn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		Host, User, Password, Dbname)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// Ping db server
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("connected")

	// Create new table
	atdb := &postgres.ArrivalTimeDb{Sql: db, TableName: "arrival_time"}
	atdb.MakeArrivalTimeTable()

	atdb.CreateRecord(
		context.Background(),
		record.ArrivalTimeRecord{StopCode: "B222", BusLine: "11", ArrivalTime: time.Now()},
	)

	records, err := atdb.SearchRecord(context.Background(), 3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(records[0].BusLine, records[0].StopCode, records[0].ArrivalTime)

	atdb.DeleteRecord(context.Background(), 1)

	// Load certificate
	cert, _ := tls.LoadX509KeyPair("localhost.crt", "localhost.key")

	// Create a server with TLS
	s := &http.Server{
		Addr:    ":9000",
		Handler: nil,
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
		},
	}

	// Handle `/` route
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello World!")
	})

	// Handle `/search/` route
	http.HandleFunc("/search/", func(rw http.ResponseWriter, r *http.Request) {
		api.SearchHandler(context.Background(), atdb, rw, r)
	})

	// Start https server
	panic(s.ListenAndServeTLS("", ""))
}

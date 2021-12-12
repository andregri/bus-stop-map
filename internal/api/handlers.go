package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/andregri/bus-stop-map/internal/record"
)

// Handler for `/search/<stop code>`
func SearchHandler(ctx context.Context, dbHandler record.ArrivalTimeHandler,
	w http.ResponseWriter, r *http.Request) {

	stopCode := r.URL.Path[len("/search/"):]

	records, err := dbHandler.SearchRecord(ctx, stopCode)
	if err != nil {
		log.Fatal(err)
	}

	for _, r := range records {
		fmt.Fprintf(w, "<div>%s - %s - %s</div>",
			r.StopCode, r.BusLine, r.ArrivalTime)
	}

}

// Handler for `/add/<stop code>/<bus line>`
func AddHandler(ctx context.Context, dbHandler record.ArrivalTimeHandler,
	w http.ResponseWriter, r *http.Request) {

	args := strings.Split(r.URL.Path[len("/add/"):], "/")

	stopCode := args[0]
	busLine := args[1]

	record := record.ArrivalTimeRecord{
		StopCode:    stopCode,
		BusLine:     busLine,
		ArrivalTime: time.Now(),
	}

	err := dbHandler.CreateRecord(ctx, record)
	if err != nil {
		log.Println(err)
	}
}

package api

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/andregri/bus-stop-map/internal/record"
)

// Handler for `/search/<stop code>`
func SearchHandler(ctx context.Context, dbHandler record.ArrivalTimeHandler,
	w http.ResponseWriter, r *http.Request) {

	//stopCode := r.URL.Path[len("/search/"):]

	records, err := dbHandler.SearchRecord(ctx, 3)
	if err != nil {
		log.Fatal(err)
	}

	for _, r := range records {
		fmt.Fprintf(w, "<div>%s</div><div>%s</div><div>%s</div>",
			r.StopCode, r.BusLine, r.ArrivalTime)
	}

}

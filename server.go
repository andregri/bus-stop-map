package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	ctx := context.Background()

	// Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Add a bus time for debugging
	busTime := &BusArrivalTime{
		StopCode:    "A123",
		BusLine:     "11",
		ArrivalTime: time.Now().String(),
	}
	err := busTime.Save(ctx, rdb)
	if err != nil {
		log.Fatal(err)
	}
	val, err := Load(ctx, rdb, busTime.StopCode)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val)

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

	// Handle `/view/` route
	http.HandleFunc("/view/", func(rw http.ResponseWriter, r *http.Request) {
		viewHandler(ctx, rdb, rw, r)
	})

	// Start https server
	log.Fatal(s.ListenAndServeTLS("", ""))
}

// Handler for `/view/<stop code>`
func viewHandler(ctx context.Context, rdb *redis.Client, w http.ResponseWriter, r *http.Request) {
	stopCode := r.URL.Path[len("/view/"):]
	bat, err := Load(ctx, rdb, stopCode)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "<div>%s</div><div>%s</div><div>%s</div>",
		bat.StopCode, bat.BusLine, bat.ArrivalTime)
}

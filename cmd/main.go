package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/andregri/bus-stop-map/internal/dbutils"
	"github.com/andregri/bus-stop-map/internal/resources"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	Port = 5432
)

func main() {
	Host := os.Getenv("POSTGRES_HOST")
	User := os.Getenv("POSTGRES_USER")
	Password := os.Getenv("POSTGRES_PASSWORD")
	Dbname := os.Getenv("POSTGRES_DB")

	// Connect to sql server
	psqlconn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		Host, User, Password, Dbname)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}
	dbutils.DB = db

	defer db.Close()

	// Ping db server
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("Connected to db")

	// Initialize tables
	dbutils.InitTables()

	//
	router := gin.Default()

	// Simple group: v1
	arrivalRouter := router.Group("/v1/arrival")
	{
		arrivalRouter.GET(":id", resources.GetArrival)
		arrivalRouter.POST("", resources.CreateArrival)
		arrivalRouter.DELETE(":id", resources.DeleteArrival)
		arrivalRouter.PATCH(":id", resources.UpdateArrival)
	}

	router.Run(":9000")
}

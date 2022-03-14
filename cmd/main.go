package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
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

	router.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, "API v1")
	})

	// Simple group: v1
	v1 := router.Group("/v1/")
	{

		v1.GET("/arrival/:id", resources.GetArrival)
		v1.POST("/arrival", resources.CreateArrival)
		v1.DELETE("/arrival/:id", resources.DeleteArrival)
		v1.PATCH("/arrival/:id", resources.UpdateArrival)
		v1.OPTIONS("/arrival", func(c *gin.Context) {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT")
			c.Header("Access-Control-Allow-Headers", "accept, content-type")
		})
	}

	router.Run(":9000")
}

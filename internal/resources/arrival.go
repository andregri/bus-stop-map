package resources

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/andregri/bus-stop-map/internal/dbutils"
	"github.com/gin-gonic/gin"
)

type Arrival struct {
	ID       int    `json:"id"`
	StopCode string `json:"stop_code"`
	BusLine  string `json:"bus_line"`
	Time     string `json:"time"`
}

// Return details of the arrival resource
func GetArrival(c *gin.Context) {
	id := c.Param("id")

	stmt, err := dbutils.DB.Prepare(`
		SELECT id, stop_code, bus_line, time
		FROM arrival
		WHERE id=$1
	`)
	if err != nil {
		log.Println("Error in select statement: ", err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		row := stmt.QueryRow(id)

		var arrival Arrival
		var t time.Time
		row.Scan(&arrival.ID, &arrival.StopCode, &arrival.BusLine, &t)
		arrival.Time = t.String()
		log.Println(arrival)

		c.JSON(http.StatusOK, arrival)
	}
}

func CreateArrival(c *gin.Context) {
	// Decode JSON request body
	decoder := json.NewDecoder(c.Request.Body)
	var arrival Arrival
	err := decoder.Decode(&arrival)
	log.Println(arrival)

	if err != nil {
		log.Println("Error in decoding JSON: ", err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {

		newID := 0
		err = dbutils.DB.QueryRow(`
			INSERT INTO arrival
			(stop_code, bus_line, time)
			VALUES ($1,$2,$3)
			RETURNING id
		`, arrival.StopCode, arrival.BusLine, arrival.Time).Scan(&newID)
		if err != nil {
			log.Println("Error in insert query: ", err)
			c.String(http.StatusInternalServerError, err.Error())
		} else {

			// Return the created ID
			c.JSON(http.StatusCreated, gin.H{
				"id": newID,
			})
		}
	}
}

func DeleteArrival(c *gin.Context) {
	panic("Not Implemented")
}

func UpdateArrival(c *gin.Context) {
	panic("Not implemented")
}

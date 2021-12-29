package resources

import (
	"github.com/gin-gonic/gin"
)

type Arrival struct {
	ID       int    `json:"id"`
	StopCode string `json:"stop_code"`
	BusLine  string `json:"bus_line"`
	Time     string `json:"time"`
}

func GetArrival(c *gin.Context) {
	panic("Not implemented")
}

func CreateArrival(c *gin.Context) {
	panic("Not implemented")
}

func DeleteArrival(c *gin.Context) {
	panic("Not Implemented")
}

func UpdateArrival(c *gin.Context) {
	panic("Not implemented")
}

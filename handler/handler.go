package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Payload struct{
	OrderID string `json:"order_id"`
	History []Location `json:"history"`

}

type Location struct{
	Latitude   float64 `json:"lat"`
	Longitude  float64 `json:"lng"`
}

//Will Create  a Data based on the Requirment
var (
	storage []Payload
	orderPresent  = make(map[string]bool)
	locationPresent = make(map[Location]bool)
)

type service struct{}

//Create a new Service Struct
func NewService() *service{
	return &service{}
}

func (s *service)Create(c* gin.Context){
	orderID := c.Param("order_id")
	 
	var location Location
	if err := c.ShouldBindJSON(&location); err!= nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"Error":err.Error(),
		})
		return
	}
	//First Entry in Storage
	if (len(storage) == 0){
		s.Create(orderID,location)
		orderPresent[orderID] = true
		locationPresent[location] =  true
		c.JSON(http.StatusOK, storage)
		return
	}
}
package helper

import (
	"net/http"

	model "json-server/model"

	"github.com/gin-gonic/gin"
)

// This will load all the data which is already available
func Get(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, model.InitialData)
}

// getting data specific data based on ID
func GetByID(c *gin.Context) {
	// this getting the reference ID from JSON that we used
	id := c.Param("id")
	// for every data inside the our hard coded data
	for _, data := range model.InitialData {
		// if the ID match with our params ID we will return that specific data
		if data.ID == id {
			c.IndentedJSON(http.StatusOK, data)
			return
		}
	}
	// if not than We are Panic
	c.IndentedJSON(http.StatusNotFound, gin.H{"Panic": "Game Not Found"})
}

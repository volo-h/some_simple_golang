package helper

import (
	model "json-server/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Post(c *gin.Context) {
	// initializing our model structure to insert new data based on MODEL
	var newGames model.Game
	// BindJSON is use to bind newDATA inside our JSON demo data
	if err := c.BindJSON(&newGames); err != nil {
		// if there is error than Panic Situation
		println("Panic: Error")
		return
	}
	// if no error than we will append or add new data to our Initial dummy data
	model.InitialData = append(model.InitialData, newGames)
	// than we will see the response
	c.IndentedJSON(http.StatusCreated, newGames)
}

package helper

import (
	model "json-server/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	// we will query the data based on the ID
	id := c.Param("id")
	// and the data model to execute new data
	var updatedGame model.Game
	// based on our dummy data we will loop through each and every data
	for i, singleGame := range model.InitialData {
		// if the ID match with our reference ID
		if singleGame.ID == id {
			// we will update all these parameters
			singleGame.ID = updatedGame.ID
			singleGame.NAME = updatedGame.NAME
			singleGame.CATEGORY = updatedGame.CATEGORY
			singleGame.PRICE = updatedGame.PRICE
			// after that we will bind this new DATA to OLD data
			if err := c.BindJSON(&updatedGame); err != nil {
				// if there is an error we will panic
				println("Panic: Error Inserting NEW DATA")
				return
			}
			// than we will append all the data as JSON
			model.InitialData = append(model.InitialData[:i], updatedGame)
			// it will give us response
			c.IndentedJSON(http.StatusCreated, updatedGame)
		}
		// println("Updated DATA: ", singleGame)
	}
}

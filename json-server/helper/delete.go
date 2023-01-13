package helper

import (
	"fmt"
	model "json-server/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	// getting the reference oof the ID
	id := c.Param("id")
	// based on  index we will get every single data
	for i, singleGame := range model.InitialData {
		// if the ID in our json match with the our parameter ID
		if singleGame.ID == id {
			// we will delete the data
			model.InitialData = append(model.InitialData[:i], model.InitialData[i+1:]...)
			// to see the response
			c.IndentedJSON(http.StatusOK, singleGame)
			// than we will print Means deleted successfully
			// fmt.Sprintf("id is %v", singleGame)
			fmt.Println("Deleted Successfully", singleGame)
		}
	}
}

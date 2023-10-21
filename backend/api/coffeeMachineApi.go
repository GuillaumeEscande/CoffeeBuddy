package api

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "coffeebuddy/service"
    "coffeebuddy/model"
)

func getCoffeeMachineByID(c *gin.Context) {
    id := c.Param("id")

	var machine, _ = service.GetCoffeeMachineByID(id)
	if machine != nil {
		c.IndentedJSON(http.StatusOK, machine) 
		return
	}
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "machine not found"})
}

func getCoffeeMachineStock(c *gin.Context) {
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "machine not found"})
}

func postCoffeeMachine(c *gin.Context) {
    var coffeeMachine model.CoffeeMachine

    // Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&coffeeMachine); err != nil {
        return
    }

	var realCoffeeMachine = service.PostCoffeeMachine(coffeeMachine)

    c.IndentedJSON(http.StatusCreated, realCoffeeMachine)
}

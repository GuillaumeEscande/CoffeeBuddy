package api

import (
    "github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {

	router := gin.Default()
	router.GET("/coffeeMachine/:id", getCoffeeMachineByID)
	router.GET("/coffeeMachine/:id/stock", getCoffeeMachineStock)
	router.POST("/coffeeMachine/:id/:user/consume-coffee", getCoffeeMachineByID)
	router.POST("/coffeeMachine/:id/:user/credit-coffee", getCoffeeMachineByID)
	router.POST("/coffeeMachine", postCoffeeMachine)
	router.GET("/user/:id", getCoffeeMachineByID)

	return router
}

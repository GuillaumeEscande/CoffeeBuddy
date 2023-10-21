package main

import (
    "coffeebuddy/api"
    "coffeebuddy/repository"
)


func main() {

    repository.InitOrm()

	router := api.NewRouter()
    
    router.Run("localhost:8080")
}
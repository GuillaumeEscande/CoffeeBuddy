package repository

import (
    "coffeebuddy/model"
)

func saveCoffeeMachine(coffeeMachine model.CoffeeMachine){
    DB.Create(coffeeMachine)
}

func getCoffeeMachineById( id string) (model.CoffeeMachine){
    var coffeeMachine model.CoffeeMachine
    DB.First(&coffeeMachine, id)
    return coffeeMachine
}

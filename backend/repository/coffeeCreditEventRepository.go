package repository

import (
    "time"
    "coffeebuddy/model"
)

var start = time.Now()

var coffeeCreditEvents = []model.CoffeeCreditEvent{
    {Datetime: start, CoffeeMachineId: "1", UserId: "1", Value: -1},
    {Datetime: start.Add(time.Hour), CoffeeMachineId: "1", UserId: "UUID", Value: 10},
    {Datetime: start.Add(time.Hour*2), CoffeeMachineId: "2", UserId: "UUID", Value: -1},
}


func saveCoffeeMachineEvent(coffeeMachineEvent model.CoffeeCreditEvent){
    DB.Create(coffeeMachineEvent)
}

func getCoffeeMachineEventByCoffeeMachineId( id string) ([]model.CoffeeMachine){
    var events []model.CoffeeMachine
    DB.Where("CoffeeMachineId <> ?", id).Find(&events)
    return events
}
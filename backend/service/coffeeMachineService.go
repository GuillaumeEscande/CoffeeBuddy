package service


import (
    "errors"
    "coffeebuddy/repository"
    "coffeebuddy/model"
)

func GetCoffeeMachineByID(id string) (*model.CoffeeMachine, error) {
    for _, coffeMachine := range repository.Machines {
        if coffeMachine.ID == id { 
            return &coffeMachine, nil
        }
    }
	return nil, errors.New("unknown machine")
}

func PostCoffeeMachine(coffeeMachine model.CoffeeMachine) (*model.CoffeeMachine) {
    return &coffeeMachine
}
package model

type CoffeeMachine struct {
    ID     string  `json:"id"`
    Name  string  `json:"name"`
    MaxCoffeeCredit int32  `json:"max"`
}

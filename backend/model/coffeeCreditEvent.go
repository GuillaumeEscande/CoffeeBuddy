package model

import (
    "time"
)

type CoffeeCreditEvent struct {
    Datetime  time.Time  `json:"datetime"`
    CoffeeMachineId  string  `json:"coffeeMachineId"`
    UserId  string  `json:"userId"`
    Value  int  `json:"value"`
}

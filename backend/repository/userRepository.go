package repository

import (
    "coffeebuddy/model"
)

func saveUser(user model.User){
    DB.Create(user)
}

func getUserById( id string) (model.User){
    var user model.User
    DB.First(&user, id)
    return user
}
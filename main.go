package main

import (
	"log"
	"startup-gin/user"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/startup_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userInput := user.RegisterUserInput{}
	userInput.Name = "Lupi"
	userInput.Email = "lupi@mail.com"
	userInput.Occupation = "Pro Gamer"
	userInput.Password = "rahasia"

	userService.RegisterUser(userInput)
}

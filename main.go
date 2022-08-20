package main

import (
	"gowork/user"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main()  {
	dsn := "root:@tcp(localhost:3306)/db_crowdf?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userInput := user.RegisterUserInput{}
	userInput.Name = "John Doe"
	userInput.Email = "john@gmail.com"
	userInput.Occupation = "Software Engineer"
	userInput.Password = "password"

	userService.RegisterUser(userInput)
}
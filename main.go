package main

import (
	"gowork/handler"
	"gowork/user"
	"log"

	// "fmt"
	// "net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(localhost:3306)/db_crowdf?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	// var address = ":9000"
	// fmt.Printf("server started at %s\n", address)

	// server := new(http.Server)
	// server.Addr = address
	// err := server.ListenAndServe()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("api/v1")

	api.POST("/user", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)

	router.Run()
	//input dari user
	///handler, mapping input dari user -> struct input
	//service : melakukan mapping dari truct input ke struct use
	// db
}

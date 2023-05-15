package main

import (
	"fmt"
	"start-up-rh/handler"
	"start-up-rh/user"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/startup_rahmat?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userByEmail, err := userRepository.FindByEmail("b@gmail.com")
	if err != nil {
		fmt.Println(err.Error())
	}

	if userByEmail.Id == 0 {
		fmt.Println("Data user tidak ditemukan")
	}
	fmt.Println(userByEmail.Name)

	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()

	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)

	router.Run()

}

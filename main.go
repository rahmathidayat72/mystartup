package main

import (
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

	//<< untuk mengcek email yang terdaftar di database
	// userByEmail, err := userRepository.FindByEmail("naro@gmail.com")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// if userByEmail.Id == 0 {
	// 	fmt.Println("Data user tidak ditemukan")
	// }
	// fmt.Println(userByEmail.Name)
	//last code

	//<< code untuk mengecek kecocokan email dan password
	// input := user.LoginInput{
	// 	Email:    "naro@gmail.com",
	// 	Password: "password",
	// }
	// user, err := userService.Login(input)
	// if err != nil {
	// 	fmt.Println("terjadi kesalahan")
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(user.Email)
	// fmt.Print(user.Name)
	//last

	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()

	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)

	router.Run()

}

package main

import (
	"fmt"
	"net/http"
	"start-up-rh/auth"
	"start-up-rh/campaign"
	"start-up-rh/handler"
	"start-up-rh/helper"
	"start-up-rh/user"
	"strings"

	"github.com/dgrijalva/jwt-go"

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
	campaignRepository := campaign.NewRepository(db)
	userService := user.NewService(userRepository)
	campaignService := campaign.NewService(campaignRepository)
	authService := auth.NewService()

	campaign, err := campaignService.FindCampaigns(22)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(campaign))
	fmt.Println(campaign)

	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()

	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checker", userHandler.CheckEmail)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatas)

	router.Run()

}

func authMiddleware(autService auth.Service, userServuce user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.ApiResponse("Unauthorization", http.StatusUnauthorized, "Error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arraySting := strings.Split(authHeader, " ")
		if len(arraySting) == 2 {
			tokenString = arraySting[1]
		}
		token, err := autService.ValidateToken(tokenString)

		if err != nil {
			response := helper.ApiResponse("Unauthorization", http.StatusUnauthorized, "Error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.ApiResponse("Unauthorization", http.StatusUnauthorized, "Error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		userId := int(claim["user_id"].(float64))

		user, err := userServuce.GetUserById(userId)
		if err != nil {
			response := helper.ApiResponse("Unauthorization", http.StatusUnauthorized, "Error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		c.Set("currentUser", user)
	}

}

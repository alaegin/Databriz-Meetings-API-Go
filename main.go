package main

import (
	"Databriz-Meetings-API-Go/configs"
	"Databriz-Meetings-API-Go/controllers"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"

	// Swagger imports
	_ "./docs" // Do NOT delete
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @Title Databriz Meetings Api
// @Description Created by mobile developers
// @BasePath /api
// @Version 0.1
func main() {
	configs.LoadConfig()
	//db.InitDatabase()
	configureRoutes()
}

// Creates Gin router and configures controllers
func configureRoutes() {
	router := gin.Default()

	router.GET("/", index)

	// Registers api groups
	api := router.Group("/api/v1")
	{
		azure := api.Group("/azure")
		{
			azureController := controllers.NewAzureController()
			azureController.RegisterRoutes(azure)
		}
	}

	// Registers swagger for documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Starts server
	if err := router.Run(":8080"); err != nil {
		fmt.Println(err)
	}
}

func index(c *gin.Context) {
	c.String(http.StatusOK, "Meeting API")
}

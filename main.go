package main

import (
	_ "./docs"
	"./src/config"
	"./src/controllers"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

// @Title Databriz Meetings Api
// @Description Created by mobile developers
// @Version 0.1
func main() {
	config.LoadConfig()
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
		azureController := controllers.NewAzureController()
		{
			azure.GET("getProjects", azureController.GetProjectsList)
			azure.GET("getProjectTeams/:projectId", azureController.GetProjectTeams)

			azure.GET("getTeamMembers/:projectId/:teamId", azureController.GetTeamMembers)
			azure.GET("getTeamIterations/:projectId/:teamId", azureController.GetTeamIterations)
			azure.GET("getMemberWorkItems/:projectId/:teamId", azureController.GetMemberWorkItems)
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

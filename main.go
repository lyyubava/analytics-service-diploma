package main

import (
	"analytics-service/controllers"
	"analytics-service/models"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func bootstrapRouter() *gin.Engine {
	r := gin.Default()
	router := r.Group("/analytics")
	router.GET("/likes", controllers.GetLikesAnalytics)
	router.GET("/comments", controllers.GetCommentsAnalytics)
	router.GET("/health", controllers.Health)
	return r
}

func main() {
	models.ConnectDatabase()
	router := bootstrapRouter()
	router.Run(fmt.Sprintf("0.0.0.0:%s", os.Getenv("API_PORT")))
}

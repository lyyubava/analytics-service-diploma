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
	r.GET("/likes", controllers.GetLikesAnalytics)
	r.GET("/comments", controllers.GetCommentsAnalytics)
	r.GET("/health", controllers.Health)
	return r
}

func main() {
	models.ConnectDatabase()
	router := bootstrapRouter()
	router.Run(fmt.Sprintf("0.0.0.0:%s", os.Getenv("API_PORT")))
}

package controllers

import (
	"analytics-service/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Analytics struct {
	Date  time.Time
	Count int
}

type AnalyticsInput struct {
	UserID string `json:"userID" binding:"required"`
}

func GetAnalytics(c *gin.Context, analytics_type string) {

	var input AnalyticsInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := input.UserID
	var analytics []Analytics
	_ = models.DB.Table(analytics_type).
		Select("DATE(created_at) AS date, COUNT(*) AS count").
		Where("created_by = ?", userID).
		Group("date").
		Order("date").
		Scan(&analytics).Error

	c.JSON(http.StatusOK, gin.H{"data": analytics})
	return

}

func GetLikesAnalytics(c *gin.Context) {
	GetAnalytics(c, "likes")
}

func GetCommentsAnalytics(c *gin.Context) {
	GetAnalytics(c, "comments")
}

package handlers

import (
	"fmt"
	"github.com/EasyPost/easypost-go/v2"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type EasyPostError struct {
	Error struct {
		Status     string `json:"Status"`
		StatusCode int    `json:"StatusCode"`
		Code       string `json:"code"`
		Message    string `json:"message"`
		Errors     []struct {
			Status     string `json:"Status"`
			StatusCode int    `json:"StatusCode"`
			Message    string `json:"message"`
			Field      string `json:"field"`
		} `json:"errors"`
	} `json:"error"`
}

func CreateTracking(c *gin.Context) {
	apiKey := os.Getenv("EASYPOST_API_KEY")
	client := easypost.New(apiKey)
	trackingCode, _ := c.GetQuery("tracking_code")
	carrier, _ := c.GetQuery("carrier")

	tracker, err := client.CreateTracker(&easypost.CreateTrackerOptions{
		TrackingCode: trackingCode,
		Carrier:      carrier,
	})
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"tracker": tracker,
		})
	}
	return
}

func GetAllTrackingPackages(c *gin.Context) {
	apiKey := os.Getenv("EASYPOST_API_KEY")
	client := easypost.New(apiKey)
	tracker, _ := client.ListTrackers(
		&easypost.ListTrackersOptions{},
	)
	c.JSON(200, gin.H{
		"tracker": tracker,
	})
}

func GetTrackingPackage(c *gin.Context) {
	apiKey := os.Getenv("EASYPOST_API_KEY")
	client := easypost.New(apiKey)
	trackerId := c.Param("tracking_id")
	tracker, _ := client.GetTracker(trackerId)
	c.JSON(200, gin.H{
		"tracker": tracker,
	})
}

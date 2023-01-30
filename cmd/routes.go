package main

import (
	"github.com/gin-gonic/gin"
	"parcel_tracker/handlers"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/create", handlers.CreateTracking)
	r.GET("/track/:id", handlers.GetTracking)
}

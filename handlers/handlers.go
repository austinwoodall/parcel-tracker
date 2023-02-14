package handlers

import (
	"fmt"
	"github.com/EasyPost/easypost-go/v2"
	"github.com/gin-gonic/gin"
	"time"
)

type CreateTracker struct {
	ID              string      `json:"id"`
	Object          string      `json:"object"`
	Mode            string      `json:"mode"`
	TrackingCode    string      `json:"tracking_code"`
	Status          string      `json:"status"`
	StatusDetail    string      `json:"status_detail"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
	SignedBy        interface{} `json:"signed_by"`
	Weight          interface{} `json:"weight"`
	EstDeliveryDate time.Time   `json:"est_delivery_date"`
	ShipmentID      interface{} `json:"shipment_id"`
	Carrier         string      `json:"carrier"`
	TrackingDetails []struct {
		Object           string      `json:"object"`
		Message          string      `json:"message"`
		Description      interface{} `json:"description"`
		Status           string      `json:"status"`
		StatusDetail     string      `json:"status_detail"`
		Datetime         time.Time   `json:"datetime"`
		Source           string      `json:"source"`
		CarrierCode      interface{} `json:"carrier_code"`
		TrackingLocation struct {
			Object  string      `json:"object"`
			City    interface{} `json:"city"`
			State   interface{} `json:"state"`
			Country interface{} `json:"country"`
			Zip     interface{} `json:"zip"`
		} `json:"tracking_location"`
	} `json:"tracking_details"`
	CarrierDetail struct {
		Object                 string      `json:"object"`
		Service                string      `json:"service"`
		ContainerType          interface{} `json:"container_type"`
		EstDeliveryDateLocal   interface{} `json:"est_delivery_date_local"`
		EstDeliveryTimeLocal   interface{} `json:"est_delivery_time_local"`
		OriginLocation         string      `json:"origin_location"`
		OriginTrackingLocation struct {
			Object  string      `json:"object"`
			City    string      `json:"city"`
			State   string      `json:"state"`
			Country interface{} `json:"country"`
			Zip     string      `json:"zip"`
		} `json:"origin_tracking_location"`
		DestinationLocation         string      `json:"destination_location"`
		DestinationTrackingLocation interface{} `json:"destination_tracking_location"`
		GuaranteedDeliveryDate      interface{} `json:"guaranteed_delivery_date"`
		AlternateIdentifier         interface{} `json:"alternate_identifier"`
		InitialDeliveryAttempt      interface{} `json:"initial_delivery_attempt"`
	} `json:"carrier_detail"`
	Finalized bool   `json:"finalized"`
	IsReturn  bool   `json:"is_return"`
	PublicURL string `json:"public_url"`
	Fees      []struct {
		Object   string `json:"object"`
		Type     string `json:"type"`
		Amount   string `json:"amount"`
		Charged  bool   `json:"charged"`
		Refunded bool   `json:"refunded"`
	} `json:"fees"`
}

func CreateTracking(c *gin.Context) {
	//apiKey := os.Getenv("EASYPOST_API_KEY")
	client := easypost.New("EZTK3ee39d13c9054b4182d398f7b5dde130TM1NPCThkMMO3c2NaSQuaQ")
	trackingCode, _ := c.GetQuery("tracking_code")
	carrier, _ := c.GetQuery("carrier")

	tracker, err := client.CreateTracker(&easypost.CreateTrackerOptions{
		TrackingCode: trackingCode,
		Carrier:      carrier,
	})
	if err != nil {
		fmt.Println(err.Error())

		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"tracker": tracker,
		})
	}
}

func GetAllTrackingPackages(c *gin.Context) {
	//apiKey := os.Getenv("EASYPOST_API_KEY")
	client := easypost.New("EZTK3ee39d13c9054b4182d398f7b5dde130TM1NPCThkMMO3c2NaSQuaQ")
	tracker, _ := client.ListTrackers(
		&easypost.ListTrackersOptions{},
	)
	c.JSON(200, gin.H{
		"tracker": tracker,
	})
}

func GetTrackingPackage(c *gin.Context) {
	//apiKey := os.Getenv("EASYPOST_API_KEY")
	client := easypost.New("EZTK3ee39d13c9054b4182d398f7b5dde130TM1NPCThkMMO3c2NaSQuaQ")
	trackerId := c.Param("id")
	tracker, _ := client.GetTracker(trackerId)
	c.JSON(200, gin.H{
		"tracker": tracker,
	})
}

package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	service "github.com/sharkbyte79/birdup/internal/service"
)

var eBirdApiKey string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load envars from .env")
	}

	eBirdApiKey = os.Getenv("EBIRD_API_KEY")
}

func main() {
	r := gin.Default()

	// Create one http client for eBird service
	hc := &http.Client{Timeout: time.Second * 10}
	s, err := service.NewEBirdService(eBirdApiKey, hc)
	if err != nil {
		return
	}

	r.GET("/observations/:region", recentObservations(s))
	r.GET("/observations/:region/notable", notableObservations(s))

	r.Run(":8080")
}

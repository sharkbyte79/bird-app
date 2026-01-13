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

var eBirdApiToken string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load envars from .env")
	}

	tok, ok := os.LookupEnv("EBIRD_API_KEY")
	if !ok {
		log.Fatal("failed to fetch eBird API key")
	}

	eBirdApiToken = tok
}

func main() {
	r := gin.Default()

	// Create one http client for eBird service
	hc := &http.Client{Timeout: time.Second * 10}
	s, err := service.NewEBirdService(eBirdApiToken, hc)
	if err != nil {
		return
	}

	r.GET("/observations/:region", recentObsHandler(s))
	r.GET("/observations/:region/notable", notableObsHandler(s))

	r.Run(":8080")
}

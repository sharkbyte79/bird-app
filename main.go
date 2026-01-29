package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	config "github.com/sharkbyte79/birdup/internal/config"
	db "github.com/sharkbyte79/birdup/internal/database"
	service "github.com/sharkbyte79/birdup/internal/service"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load envars")
	}

	r := gin.Default()

	// Create one http client for eBird service
	hc := &http.Client{Timeout: time.Second * 10}
	s, err := service.NewEBirdService(cfg.EBirdAPIToken, hc)
	if err != nil {
		log.Fatal("failed to create ebird service")
	}

	dsn := fmt.Sprintf("host=%s port=%s password=%s dbname=%s user=%s sslmode=disable",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.Password, cfg.DB.DB, cfg.DB.User)

	// TODO Pass store to user repository implementation
	_, err = db.NewStore(dsn)
	if err != nil {
		log.Fatal("Failed to open database connection")
	}

	r.GET("/observations/:region", recentObsHandler(s))
	r.GET("/observations/:region/notable", notableObsHandler(s))

	r.Run(fmt.Sprintf(":%s", cfg.Port))
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sharkbyte79/birdup/internal/config"
	db "github.com/sharkbyte79/birdup/internal/database"
	"github.com/sharkbyte79/birdup/internal/handler"
	"github.com/sharkbyte79/birdup/internal/repository"
	"github.com/sharkbyte79/birdup/internal/service"
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

	store, err := db.NewStore(dsn)
	if err != nil {
		log.Fatal("Failed to open database connection")
	}

	userRepo, err := repository.NewRepository(store)
	if err != nil {
		log.Fatal("failed to create User repository")
	}
	userService, err := service.NewUserService(userRepo)

	// TODO move this to route package, protect route for only authorized usage
	r.POST("/user/create", handler.CreateUserHandler(userService))

	r.GET("/observations/:region", recentObsHandler(s))
	r.GET("/observations/:region/notable", notableObsHandler(s))

	r.Run(fmt.Sprintf(":%s", cfg.Port))
}

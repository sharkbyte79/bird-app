package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ebs "github.com/sharkbyte79/birdup/internal/service"
)

// recentObsHandler returns a HandlerFunc that handles a GET request to
// retrieve a bundle of recent bird observations.
func recentObsHandler(s *ebs.EBirdService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		region := ctx.Param("region")

		res, err := s.RecentObsByRegion(region, 14, 30)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		ctx.IndentedJSON(http.StatusOK, res)
	}
}

// recentObsHandler returns a HandlerFunc that handles a GET request to
// retrieve a bundle of notable bird observations.
func notableObsHandler(s *ebs.EBirdService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		region := ctx.Param("region")

		res, err := s.NotableObsByRegion(region, 14, 30)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		ctx.IndentedJSON(http.StatusOK, res)
	}
}

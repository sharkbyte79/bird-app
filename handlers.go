package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ebs "github.com/sharkbyte79/birdup/internal/service"
)

// RecentObservations handles a GET request to a bundle of bird observations
// from the eBird API according to the given region code.
func recentObservations(s *ebs.EBirdService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		region := ctx.Param("region")

		res, err := s.RecentObsByRegion(region, 14, 30)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		ctx.IndentedJSON(http.StatusOK, gin.H{"data": res})
	}
}

func notableObservations(s *ebs.EBirdService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		region := ctx.Param("region")

		res, err := s.NotableObsByRegion(region, 14, 30)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		ctx.IndentedJSON(http.StatusOK, gin.H{"data": res})
	}
}

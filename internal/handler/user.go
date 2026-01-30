package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sharkbyte79/birdup/internal/dto"
	"github.com/sharkbyte79/birdup/internal/model"

	// models "github.com/sharkbyte79/birdup/internal/model"
	svc "github.com/sharkbyte79/birdup/internal/service"
)

// CreateUserHandler returns a HandlerFunc that handles a POST request to
// create a new User record.
func CreateUserHandler(s *svc.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var newUser dto.CreateUserRequest

		// Attempt to bind JSON fields to user creation dto
		if err := ctx.BindJSON(&newUser); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		u := &model.User{FirebaseId: newUser.FirebaseID, Email: newUser.Email, CreatedAt: time.Now()}
		if err := s.Create(u); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		// ctx.JSON(http.StatusCreated, gin.H{
		// 	"message": "new user was created",
		// })
		ctx.IndentedJSON(http.StatusCreated, u)
	}
}

// func GetUserHandler(s *svc.UserService) gin.HandlerFunc {}

// func DeleteUserHandler(s *svc.UserService) gin.HandlerFunc {}

// func UpdateUserHandler(s *svc.UserService) gin.HandlerFunc {}

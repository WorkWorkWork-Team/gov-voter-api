package handler

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/WorkWorkWork-Team/gov-voter-api/service"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service service.UserService
}

func NewUserHandler(getInformation service.UserService) userHandler {
	return userHandler{
		service: getInformation,
	}
}

func (g *userHandler) GetuserInfo(gi *gin.Context) {
	userInfo, err := g.service.GetUserInformation(gi.Param("CitizenID"))
	if err == nil {
		gi.JSON(http.StatusOK, gin.H{
			"info": userInfo,
		})
		return
	} else if errors.Is(err, sql.ErrNoRows) {
		gi.JSON(http.StatusNotFound, gin.H{
			"message": "Not matching data",
		})
		return
	}
	gi.JSON(http.StatusInternalServerError, gin.H{
		"message": "Something went wrong.",
	})
}

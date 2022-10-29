package handler

import (
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
	userInfo, err := g.service.CheckGetUserInformation(gi.Param("CitizenID"))
	if err == nil {
		gi.JSON(http.StatusOK, gin.H{
			"info": userInfo,
		})
	} else {
		gi.Status(http.StatusBadRequest)
	}

}

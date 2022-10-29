package handler

import (
	"net/http"

	"github.com/WorkWorkWork-Team/gov-voter-api/service"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service service.GetUserInformationService
}

func NewUserHandler(getInformation service.GetUserInformationService) userHandler {
	return userHandler{
		service: getInformation,
	}
}

func (g *userHandler) GetuserInfo(gi *gin.Context) {
	userInfo, status := g.service.CheckGetUserInformation(gi.Param("CitizenID"))
	if status {
		gi.JSON(http.StatusOK, gin.H{
			"info": userInfo,
		})
	} else {
		gi.Status(http.StatusBadRequest)
	}

}

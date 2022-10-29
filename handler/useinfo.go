package handler

import (
	"net/http"

	"github.com/WorkWorkWork-Team/gov-voter-api/service"
	"github.com/gin-gonic/gin"
)

type getinformationHandler struct {
	service service.GetUserInformationService
}

func NewGetUserInformationHandler(getInformation service.GetUserInformationService) getinformationHandler {
	return getinformationHandler{
		service: getInformation,
	}
}

func (g *getinformationHandler) GetuserInfo(gi *gin.Context) {
	userInfo, status := g.service.CheckGetUserInformation(gi.Param("CitizendID"))
	if status {
		gi.JSON(http.StatusOK, gin.H{
			"test": userInfo,
		})
	} else {
		gi.Status(http.StatusBadRequest)
	}
	
}

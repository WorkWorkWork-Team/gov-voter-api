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
	populationInfo := g.service.GetUserInformation()
	gi.JSON(http.StatusOK, gin.H{
		"test": populationInfo,
	})
}

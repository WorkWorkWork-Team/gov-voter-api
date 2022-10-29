package handler

import (
	"net/http"

	model "github.com/WorkWorkWork-Team/gov-voter-api/models"
	"github.com/WorkWorkWork-Team/gov-voter-api/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type authenticateHandler struct {
	service service.AuthenticationService
}

func NewAuthenticateHandler(service service.AuthenticationService) authenticateHandler {
	return authenticateHandler{
		service: service,
	}
}

func (a *authenticateHandler) AuthAndGenerateToken(g *gin.Context) {
	var body model.AuthenticateBody
	err := g.BindJSON(&body)
	if err != nil {
		logrus.Error("BindJson error: ", err)
		g.JSON(http.StatusBadRequest, gin.H{
			"message": "Can't parse request body. It should contain `citizenID` and `lazerID`",
		})
		return
	}

	token, err := a.service.Authenticate(body.CitizenID, body.LazerID)
	if err != nil {
		g.JSON(http.StatusUnauthorized, gin.H{
			"message": "Can't authenticate with this citizenID and lazerID",
		})
		return
	}
	g.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

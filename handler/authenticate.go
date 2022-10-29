package handler

import (
	"net/http"

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
	token, err := a.service.Authenticate(1234567891234, "1234AB")
	logrus.Error(err)
	logrus.Info(token)
	g.Status(http.StatusOK)
}

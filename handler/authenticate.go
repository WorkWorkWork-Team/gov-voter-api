package handler

import (
	"net/http"

	"github.com/WorkWorkWork-Team/gov-voter-api/service"
	"github.com/gin-gonic/gin"
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
	a.service.Authenticate("", "")
	g.Status(http.StatusOK)
}

package handler

import (
	"net/http"

	"github.com/WorkWorkWork-Team/gov-voter-api/service"
	"github.com/gin-gonic/gin"
)

type validityHandler struct {
	jwtService service.JWTService
	service    service.ValidityService
}

func NewValidityHandler(jwtService service.JWTService, service service.ValidityService) validityHandler {
	return validityHandler{
		jwtService: jwtService,
		service:    service,
	}
}

func (v *validityHandler) Validity(g *gin.Context) {
	result := v.service.CheckValidity(g.Param("CitizenID"))
	if result {
		g.Status(http.StatusOK)
		return
	}
	g.Status(http.StatusBadRequest)
}

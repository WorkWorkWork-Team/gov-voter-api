package handler

import (
	"net/http"

	"github.com/WorkWorkWork-Team/common-go/jwtservice"
	"github.com/WorkWorkWork-Team/gov-voter-api/service"
	"github.com/gin-gonic/gin"
)

type validityHandler struct {
	jwtValidator jwtservice.JWTService
	service      service.ValidityService
}

func NewValidityHandler(jwtValidator jwtservice.JWTService, service service.ValidityService) validityHandler {
	return validityHandler{
		jwtValidator: jwtValidator,
		service:      service,
	}
}

func (v *validityHandler) validity(g *gin.Context) {
	// Validate the header.
	// TODO: get token and get its claim
	_, err := v.jwtValidator.ValidateToken(g.Request.Header["Authorization"][1])
	if err != nil {
		g.Status(http.StatusForbidden)
		return
	}

	v.service.CheckValidity("")
}

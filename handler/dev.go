package handler

import (
	"net/http"

	"github.com/WorkWorkWork-Team/common-go/jwtservice"
	"github.com/gin-gonic/gin"
)

type devHandler struct {
	jwtService jwtservice.JWTService
}

func NewDevHandler(jwtService jwtservice.JWTService) devHandler {
	return devHandler{
		jwtService: jwtService,
	}
}

func (d *devHandler) NewTestToken(g *gin.Context) {
	token, err := d.jwtService.GenerateToken(g.Param("id"))
	if err != nil {
		g.Status(http.StatusInternalServerError)
		return
	}
	g.String(200, token)
}

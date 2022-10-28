package handler

import (
	"net/http"

	"github.com/WorkWorkWork-Team/gov-voter-api/service"
	"github.com/gin-gonic/gin"
)

type devHandler struct {
	jwtService service.JWTService
}

func NewDevHandler(jwtService service.JWTService) devHandler {
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
	g.JSON(200, gin.H{
		"token": token,
	})
}

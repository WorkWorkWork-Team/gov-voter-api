package handler

import (
	"net/http"

	"github.com/WorkWorkWork-Team/common-go/jwtservice"
	"github.com/gin-gonic/gin"
)

type getinfomationHandler struct {
	jwtgetuser jwtservice.JWTService
}

func NewGetUserInformationHandler(jwtgetuser jwtservice.JWTService) getinfomationHandler {
	return getinfomationHandler{
		jwtgetuser: jwtgetuser,
	}
}

func (g *getinfomationHandler) getuserInfo(gi *gin.Context) {
	_, err := g.jwtgetuser.ValidateToken(gi.Request.Header["Authorization"][1])
	if err != nil {
		gi.Status(http.StatusForbidden)
		return
	}
}

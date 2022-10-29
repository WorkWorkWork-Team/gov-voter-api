package handler

import (
	"github.com/WorkWorkWork-Team/gov-voter-api/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type voteHandler struct {
	jwtService       service.JWTService
	validityService  service.ValidityService
	applyvoteService service.ApplyvoteService
}

func NewVoteHandler(jwtService service.JWTService, validityService service.ValidityService, applyvoteService service.ApplyvoteService) voteHandler {
	return voteHandler{
		jwtService:       jwtService,
		validityService:  validityService,
		applyvoteService: applyvoteService,
	}
}

func (v *voteHandler) Validity(g *gin.Context) {
	result := v.validityService.CheckValidity(g.Param("CitizenID"))
	if result {
		g.Status(http.StatusOK)
		return
	}
	g.Status(http.StatusBadRequest)
}

func (v *voteHandler) ApplyVote(g *gin.Context) {
	err := v.applyvoteService.ApplyVote(g.Param("CitizenID"))
	if err != nil {
		g.Status(http.StatusInternalServerError)
		return
	}
	g.Status(http.StatusOK)
	return
}

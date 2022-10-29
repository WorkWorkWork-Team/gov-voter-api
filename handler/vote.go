package handler

import (
	"errors"
	"net/http"

	"github.com/WorkWorkWork-Team/gov-voter-api/service"
	"github.com/gin-gonic/gin"
)

type voteHandler struct {
	jwtService  service.JWTService
	voteService service.VoteService
}

func NewVoteHandler(jwtService service.JWTService, voteService service.VoteService) voteHandler {
	return voteHandler{
		jwtService:  jwtService,
		voteService: voteService,
	}
}

func (v *voteHandler) Validity(g *gin.Context) {
	result := v.voteService.CheckValidity(g.Param("CitizenID"))
	if result {
		g.Status(http.StatusOK)
		return
	}
	g.Status(http.StatusBadRequest)
}

func (v *voteHandler) ApplyVote(g *gin.Context) {
	err := v.voteService.ApplyVote(g.Param("CitizenID"))
	if errors.Is(err, service.ErrUserAlreadyApplied) {
		g.JSON(http.StatusBadRequest, gin.H{
			"message": "Already applied",
		})
		return
	} else if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
		})
		return
	}
	g.JSON(http.StatusOK, gin.H{
		"message": "summited",
	})
	return
}

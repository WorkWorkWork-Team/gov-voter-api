package handler

import (
	"github.com/WorkWorkWork-Team/gov-voter-api/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type applyvoteHandler struct {
	applyvoteService service.ApplyvoteService
}

func NewApplyVoteHandler(applyvoteService service.ApplyvoteService) applyvoteHandler {
	return applyvoteHandler{
		applyvoteService: applyvoteService,
	}
}

func (a *applyvoteHandler) ApplyVote(g *gin.Context) {
	err := a.applyvoteService.ApplyVote(g.Param("CitizenID"))
	if err != nil {
		g.Status(http.StatusInternalServerError)
		return
	}
	g.Status(http.StatusOK)
	return
}

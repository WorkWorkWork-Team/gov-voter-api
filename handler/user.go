package handler

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/WorkWorkWork-Team/gov-voter-api/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	populationService service.PopulationService
	jwtService        service.JWTService
	voteService       service.VoteService
}

func NewUserHandler(populationService service.PopulationService, jwtService service.JWTService, voteService service.VoteService) *UserHandler {
	return &UserHandler{
		populationService: populationService,
		jwtService:        jwtService,
		voteService:       voteService,
	}
}

func (u *UserHandler) GetUserInfo(gi *gin.Context) {
	populationInfo, err := u.populationService.GetPopulationInformation(gi.Param("CitizenID"))
	if err == nil {
		gi.JSON(http.StatusOK, populationInfo)
		return
	} else if errors.Is(err, sql.ErrNoRows) {
		gi.JSON(http.StatusNotFound, gin.H{
			"message": "Not matching data",
		})
		return
	}
	gi.JSON(http.StatusInternalServerError, gin.H{
		"message": "Something went wrong.",
	})
}

func (u *UserHandler) Validity(g *gin.Context) {
	result := u.voteService.CheckValidity(g.Param("CitizenID"))
	if result {
		g.Status(http.StatusOK)
		return
	}
	g.Status(http.StatusBadRequest)
}

func (u *UserHandler) ApplyVote(g *gin.Context) {
	err := u.voteService.ApplyVote(g.Param("CitizenID"))
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
}

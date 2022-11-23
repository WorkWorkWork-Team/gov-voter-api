package handler

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/WorkWorkWork-Team/gov-voter-api/service"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	populationService service.PopulationService
	jwtService        service.JWTService
	voteService       service.VoteService
}

func NewUserHandler(populationService service.PopulationService, jwtService service.JWTService, voteService service.VoteService) userHandler {
	return userHandler{
		populationService: populationService,
		jwtService:        jwtService,
		voteService:       voteService,
	}
}

func (u *userHandler) GetUserInfo(g *gin.Context) {
	populationInfo, err := u.populationService.GetPopulationInformation(g.Param("CitizenID"))
	if err == nil {
		g.JSON(http.StatusOK, populationInfo)
		return
	} else if errors.Is(err, sql.ErrNoRows) {
		g.JSON(http.StatusNotFound, gin.H{
			"message": "Not matching data",
		})
		return
	}
	g.JSON(http.StatusInternalServerError, gin.H{
		"message": "Something went wrong.",
	})
}

func (u *userHandler) Validity(g *gin.Context) {
	tableID, err := strconv.Atoi(g.Param("TableID"))
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can't resolve table id",
		})
		return
	}
	result, err := u.voteService.CheckValidity(g.Param("CitizenID"), tableID)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong in check validity",
		})
		return
	}
	if result {
		g.Status(http.StatusOK)
		return
	}
	g.Status(http.StatusBadRequest)
}

func (u *userHandler) ApplyVote(g *gin.Context) {
	tableID, err := strconv.Atoi(g.Param("TableID"))
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can't resolve table id",
		})
		return
	}
	err = u.voteService.ApplyVote(g.Param("CitizenID"), tableID)
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

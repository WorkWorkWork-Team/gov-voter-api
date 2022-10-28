package repository

import (
	"errors"

	model "github.com/WorkWorkWork-Team/gov-voter-api/models"
	"github.com/jmoiron/sqlx"
)

type applyVoteRepository struct {
	mysql *sqlx.DB
}

type ApplyVoteRepository interface {
	GetApplyVoteByCitizenID(citizenID string) (model.ApplyVote, error)
}

var ErrFoundMoreThanOne error = errors.New("found more than one row")

func NewApplyVoteRepository(mysql *sqlx.DB) ApplyVoteRepository {
	return &applyVoteRepository{
		mysql: mysql,
	}
}

func (a *applyVoteRepository) GetApplyVoteByCitizenID(citizenID string) (applyVote model.ApplyVote, err error) {
	var applyVoteList []model.ApplyVote
	err = a.mysql.Select(applyVoteList, "SELECT * FROM `ApplyVote` WHERE citizenID=:citizenID", citizenID)
	if err != nil {
		return applyVote, err
	}

	if len(applyVoteList) != 1 {
		return applyVote, ErrFoundMoreThanOne
	}
	return applyVoteList[0], nil
}

package repository

import (
	"errors"
	"github.com/sirupsen/logrus"

	model "github.com/WorkWorkWork-Team/gov-voter-api/models"
	"github.com/jmoiron/sqlx"
)

type applyVoteRepository struct {
	mysql *sqlx.DB
}

type ApplyVoteRepository interface {
	ApplyVote(citizenID string) error
	GetApplyVoteByCitizenID(citizenID string) (model.ApplyVote, error)
}

var ErrFoundMoreThanOne error = errors.New("found more than one row")
var ErrNotFound error = errors.New("not found")

func NewApplyVoteRepository(mysql *sqlx.DB) ApplyVoteRepository {
	return &applyVoteRepository{
		mysql: mysql,
	}
}

func (a *applyVoteRepository) ApplyVote(citizenID string) error {
	_, err := a.mysql.Query("INSERT INTO ApplyVote (CitizenID) VALUES (?)", citizenID)
	if err != nil {
		logrus.Error(err)
	}
	return err
}

func (a *applyVoteRepository) GetApplyVoteByCitizenID(citizenID string) (applyVote model.ApplyVote, err error) {
	err = a.mysql.Get(&applyVote, "SELECT * FROM `ApplyVote` WHERE CitizenID=?", citizenID)
	return applyVote, err
}

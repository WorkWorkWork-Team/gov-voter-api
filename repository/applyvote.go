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
	ApplyVoteToDB(citizenID string) error
	GetApplyVoteByCitizenID(citizenID string) (model.ApplyVote, error)
}

var ErrFoundMoreThanOne error = errors.New("found more than one row")
var ErrNotFound error = errors.New("not found")

func NewApplyVoteRepository(mysql *sqlx.DB) ApplyVoteRepository {
	return &applyVoteRepository{
		mysql: mysql,
	}
}

func (a *applyVoteRepository) ApplyVoteToDB(citizenID string) error {
	rows, err := a.mysql.Query("INSERT INTO ApplyVote (CitizenID) VALUES (?)", citizenID)
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}

func (a *applyVoteRepository) GetApplyVoteByCitizenID(citizenID string) (applyVote model.ApplyVote, err error) {
	var applyVoteList []model.ApplyVote
	err = a.mysql.Select(&applyVoteList, "SELECT * FROM `ApplyVote` WHERE citizenID=?", citizenID)
	if err != nil {
		return applyVote, err
	}

	applyVoteLength := len(applyVoteList)
	if applyVoteLength == 0 {
		return applyVote, ErrNotFound
	} else if applyVoteLength > 1 {
		return applyVote, ErrFoundMoreThanOne
	}
	return applyVoteList[0], nil
}

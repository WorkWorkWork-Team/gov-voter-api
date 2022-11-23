package repository

import (
	"errors"
	"fmt"
	model "github.com/WorkWorkWork-Team/gov-voter-api/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type applyVoteRepository struct {
	mysql *sqlx.DB
}

type ApplyVoteRepository interface {
	ApplyVote(citizenID string, tableName string) error
	GetApplyVoteByCitizenID(citizenID string, table string) (model.ApplyVote, error)
}

var ErrFoundMoreThanOne error = errors.New("found more than one row")
var ErrNotFound error = errors.New("not found")

func NewApplyVoteRepository(mysql *sqlx.DB) ApplyVoteRepository {
	return &applyVoteRepository{
		mysql: mysql,
	}
}

func (a *applyVoteRepository) ApplyVote(citizenID string, tableName string) error {
	var query = fmt.Sprintf("INSERT INTO %s (CitizenID) VALUES (%s)", tableName, citizenID)
	_, err := a.mysql.Query(query)
	if err != nil {
		logrus.Error(err)
	}
	return err
}

func (a *applyVoteRepository) GetApplyVoteByCitizenID(citizenID string, tableName string) (applyVote model.ApplyVote, err error) {
	var applyVoteList []model.ApplyVote
	var query = fmt.Sprintf("SELECT * FROM %s WHERE citizenID=%s", tableName, citizenID)
	err = a.mysql.Select(&applyVoteList, query)
	if err != nil {
		logrus.Info(err)
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

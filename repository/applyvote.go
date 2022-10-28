package repository

import (
	"github.com/jmoiron/sqlx"
)

type applyVoteRepository struct {
	mysql *sqlx.DB
}

type ApplyVoteRepository interface {
	ApplyVoteToDB(citizenID string) error
}

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

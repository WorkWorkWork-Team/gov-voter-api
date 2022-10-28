package repository

import "database/sql"

type applyVoteRepository struct {
	mysql *sql.DB
}

type ApplyVoteRepository interface {
}

func NewApplyVoteRepository(mysql *sql.DB) ApplyVoteRepository {
	return &applyVoteRepository{
		mysql: mysql,
	}
}

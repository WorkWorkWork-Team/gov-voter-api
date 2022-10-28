package repository

import (
	"github.com/jmoiron/sqlx"
)

type getUserInformation struct {
	mysql *sqlx.DB
}

type GetUserInformationRepository interface {
}

func NewGetUserInformtaionRepostory(mysql *sqlx.DB) GetUserInformationRepository {
	return &getUserInformation{
		mysql: mysql,
	}
}

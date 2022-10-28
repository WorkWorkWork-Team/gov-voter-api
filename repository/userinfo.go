package repository

import "database/sql"

type getUserInformation struct {
	mysql *sql.DB
}

type GetUserInformationRepository interface {
}

func NewGetUserInformtaionRepostory(mysql *sql.DB) GetUserInformationRepository {
	return &getUserInformation{
		mysql: mysql,
	}
}

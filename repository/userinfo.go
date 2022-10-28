package repository

import "database/sql"

type getUserInformation struct {
	mysql *sql.DB
}

type GetUserInformation interface {
}

func NewGetUserInformtaion(mysql *sql.DB) GetUserInformation {
	return &getUserInformation{
		mysql: mysql,
	}
}

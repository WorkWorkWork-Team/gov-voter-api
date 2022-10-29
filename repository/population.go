package repository

import (
	"errors"

	model "github.com/WorkWorkWork-Team/gov-voter-api/models"
	"github.com/jmoiron/sqlx"
)

type getUserInformationRepository struct {
	mysql *sqlx.DB
}

type GetUserInformationRepository interface {
	GetUserInfo(citizenID string) (model.UserInfo, error)
}

func NewGetUserInformtaionRepostory(mysql *sqlx.DB) GetUserInformationRepository {
	return &getUserInformationRepository{
		mysql: mysql,
	}
}

func (g *getUserInformationRepository) GetUserInfo(citizenID string) (userinfo model.UserInfo, err error) {
	var getUserInfoList []model.UserInfo
	err = g.mysql.Select(&getUserInfoList, "SELECT * from `Population` WHERE citizenID=?", citizenID)
	if err != nil {
		return userinfo, err
	}

	userInfoLenght := len(getUserInfoList)
	if userInfoLenght == 0 {
		return userinfo, errors.New("Not match CitizenID")
	} else if userInfoLenght > 1 {
		return userinfo, errors.New("Found more than ")
	}
	return userinfo, nil
}

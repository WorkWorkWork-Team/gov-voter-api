package repository

import (
	model "github.com/WorkWorkWork-Team/gov-voter-api/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type populationRepository struct {
	mysql *sqlx.DB
}

type PopulationRepository interface {
	GetUserInfo(citizenID string) (model.UserInfo, error)
	GetUserInfoBasedOnCitizenIDAndLazerID(citizenID string, lazerID string) (model.UserInfo, error)
}

func NewPopulationRepostory(mysql *sqlx.DB) PopulationRepository {
	return &populationRepository{
		mysql: mysql,
	}
}

func (p *populationRepository) GetUserInfo(citizenID string) (userinfo model.UserInfo, err error) {
	var getUserInfoList []model.UserInfo
	err = p.mysql.Select(&getUserInfoList, "SELECT * from `Population` WHERE citizenID=?", citizenID)
	if err != nil {
		return userinfo, err
	}

	userInfoLenght := len(getUserInfoList)
	if userInfoLenght == 0 {
		return userinfo, ErrNotFound
	} else if userInfoLenght > 1 {
		return userinfo, ErrFoundMoreThanOne
	}
	return userinfo, nil
}
func (p *populationRepository) GetUserInfoBasedOnCitizenIDAndLazerID(citizenID string, lazerID string) (userInfo model.UserInfo, err error) {
	var userInfoList model.UserInfo
	err = p.mysql.Get(&userInfoList, "SELECT * from `Population` WHERE CitizenID=? AND LazerID=?", citizenID, lazerID)
	if err != nil {
		logrus.Error("GetUserInfoBasedOnCitizenIDAndLazerID err:", err)
		return userInfo, err
	}
	return userInfoList, nil
}

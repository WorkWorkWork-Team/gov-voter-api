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
	GetUserInfo(citizenID string) (model.Population, error)
}

func NewPopulationRepository(mysql *sqlx.DB) PopulationRepository {
	return &populationRepository{
		mysql: mysql,
	}
}

func (p *populationRepository) GetUserInfo(citizenID string) (userinfo model.Population, err error) {
	var populationInfo model.Population
	err = p.mysql.Get(&populationInfo, "SELECT * FROM `Population` WHERE citizenID=?", citizenID)
	if err != nil {
		return populationInfo, err
	}
	return populationInfo, nil
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

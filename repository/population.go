package repository

import (
	model "github.com/WorkWorkWork-Team/gov-voter-api/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type getUserInformationRepository struct {
	mysql *sqlx.DB
}

type PopulationRepository interface {
	GetUserInfo(citizenID string) (model.UserInfo, error)
	GetUserInfoBasedOnCitizenIDAndLazerID(citizenID int, lazerID string) (model.UserInfo, error)
}

func NewPopulationRepostory(mysql *sqlx.DB) PopulationRepository {
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
		return userinfo, ErrNotFound
	} else if userInfoLenght > 1 {
		return userinfo, ErrFoundMoreThanOne
	}
	return userinfo, nil
}
func (g *getUserInformationRepository) GetUserInfoBasedOnCitizenIDAndLazerID(citizenID int, lazerID string) (userInfo model.UserInfo, err error) {
	var userInfoList []*model.UserInfo
	err = g.mysql.Select(&userInfoList, "SELECT * from `Population` WHERE citizenID=? AND LazerID=?", citizenID, lazerID)
	if err != nil {
		logrus.Error("GetUserInfoBasedOnCitizenIDAndLazerID err:", err)
		return userInfo, err
	}
	logrus.Info(*userInfoList[0])
	userInfoLenght := len(userInfoList)
	if userInfoLenght == 0 {
		return userInfo, ErrNotFound
	} else if userInfoLenght > 1 {
		return userInfo, ErrFoundMoreThanOne
	}
	return *userInfoList[0], nil
}

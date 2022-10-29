package repository

import (
	"fmt"
	"strconv"

	model "github.com/WorkWorkWork-Team/gov-voter-api/models"
	"github.com/jmoiron/sqlx"
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

func (g *populationRepository) GetUserInfo(citizenID string) (userInfo model.Population, err error) {
	var getUserInfoList []model.Population
	fmt.Println(citizenID)
	var id, _ = strconv.Atoi(citizenID)
	err = g.mysql.Select(&getUserInfoList, "SELECT * FROM `Population` WHERE citizenID=?", id)
	if err != nil {
		return userInfo, err
	}

	userInfoLenght := len(getUserInfoList)
	if userInfoLenght == 0 {
		return userInfo, ErrNotFound
	} else if userInfoLenght > 1 {
		return userInfo, ErrFoundMoreThanOne
	}
	return getUserInfoList[0], nil
}

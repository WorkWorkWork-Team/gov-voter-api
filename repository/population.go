package repository

import (
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
	var getUserInfoList model.Population
	err = g.mysql.Get(&getUserInfoList, "SELECT * FROM `Population` WHERE citizenID=?", citizenID)
	if err != nil {
		return userInfo, err
	}
	return getUserInfoList, nil
}

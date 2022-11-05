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
	GetPopulationInfo(citizenID string) (model.Population, error)
	GetPopulationInfoBasedOnCitizenIDAndLazerID(citizenID string, lazerID string) (populationInfo model.Population, err error)
}

func NewPopulationRepository(mysql *sqlx.DB) PopulationRepository {
	return &populationRepository{
		mysql: mysql,
	}
}

func (p *populationRepository) GetPopulationInfo(citizenID string) (populationInfo model.Population, err error) {
	err = p.mysql.Get(&populationInfo, "SELECT * FROM `Population` WHERE citizenID=?", citizenID)
	if err != nil {
		return populationInfo, err
	}
	return populationInfo, nil
}
func (p *populationRepository) GetPopulationInfoBasedOnCitizenIDAndLazerID(citizenID string, lazerID string) (populationInfo model.Population, err error) {
	err = p.mysql.Get(&populationInfo, "SELECT * from `Population` WHERE CitizenID=? AND LazerID=?", citizenID, lazerID)
	if err != nil {
		logrus.Error("GetPopulationInfoBasedOnCitizenIDAndLazerID err:", err)
		return populationInfo, err
	}
	return populationInfo, nil
}

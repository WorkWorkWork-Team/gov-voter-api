package service

import (
	model "github.com/WorkWorkWork-Team/gov-voter-api/models"
	"github.com/WorkWorkWork-Team/gov-voter-api/repository"
	"github.com/sirupsen/logrus"
)

type populationService struct {
	populationRepository repository.PopulationRepository
}

type PopulationService interface {
	GetPopulationInformation(citizenID string) (model.Population, error)
}

func NewPopulationService(populationRepository repository.PopulationRepository) PopulationService {
	return &populationService{
		populationRepository: populationRepository,
	}
}

func (g *populationService) GetPopulationInformation(citizenID string) (model.Population, error) {
	logrus.Info("Start Check Infomation")
	defer logrus.Info("Complete Checking Infomation")

	info, err := g.populationRepository.GetPopulationInfo(citizenID)
	return info, err
}

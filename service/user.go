package service

import (
	model "github.com/WorkWorkWork-Team/gov-voter-api/models"
	"github.com/WorkWorkWork-Team/gov-voter-api/repository"
	"github.com/sirupsen/logrus"
)

type userService struct {
	populationRepository repository.PopulationRepository
}

type UserService interface {
	GetUserInformation(citizenID string) (model.Population, error)
}

func NewUserService(populationRepository repository.PopulationRepository) UserService {
	return &userService{
		populationRepository: populationRepository,
	}
}

func (g *userService) GetUserInformation(citizenID string) (model.Population, error) {
	logrus.Info("Start Check Infomation")
	defer logrus.Info("Complete Checking Infomation")

	info, err := g.populationRepository.GetUserInfo(citizenID)
	return info, err
}

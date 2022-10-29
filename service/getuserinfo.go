package service

import (
	model "github.com/WorkWorkWork-Team/gov-voter-api/models"
	"github.com/WorkWorkWork-Team/gov-voter-api/repository"
	"github.com/sirupsen/logrus"
)

type getUserInformationService struct {
	getUserInfoRepository repository.PopulationRepository
}

type GetUserInformationService interface {
	CheckGetUserInformation(citizenID string) (model.Population, bool)
}

func NewGetUserInformtaionService(getUserInfoRepository repository.PopulationRepository) GetUserInformationService {
	return &getUserInformationService{
		getUserInfoRepository: getUserInfoRepository,
	}
}

func (g *getUserInformationService) CheckGetUserInformation(citizenID string) (model.Population, bool) {
	logrus.Info("Start Check Infomation")
	defer logrus.Info("Complete Checking Infomation")

	info, err := g.getUserInfoRepository.GetUserInfo(citizenID)
	if err != nil {
		return info, false
	}
	return info, true
}

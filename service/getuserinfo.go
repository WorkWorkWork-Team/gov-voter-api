package service

import "github.com/WorkWorkWork-Team/gov-voter-api/repository"

type getUserInformationService struct {
	getUserInfoRepository repository.PopulationRepository
}

type GetUserInformationService interface {
	GetUserInformation() bool
}

func NewGetUserInformtaionService(getUserInfoRepository repository.PopulationRepository) GetUserInformationService {
	return &getUserInformationService{
		getUserInfoRepository: getUserInfoRepository,
	}
}

func (g *getUserInformationService) GetUserInformation() bool {
	return false
}

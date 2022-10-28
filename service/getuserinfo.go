package service

import "github.com/WorkWorkWork-Team/gov-voter-api/repository"

type getUserInformationService struct {
	getUserInfoRepository repository.GetUserInformation
}

type GetUserInformationService interface {
}

func NewGetUserInformtaionService(getUserInfoRepository repository.GetUserInformation) GetUserInformationService {
	return &getUserInformationService{
		getUserInfoRepository: getUserInfoRepository,
	}
}

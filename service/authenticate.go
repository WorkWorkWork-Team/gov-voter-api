package service

import (
	"errors"
	"fmt"

	"github.com/WorkWorkWork-Team/gov-voter-api/repository"
	"github.com/sirupsen/logrus"
)

type authenticationService struct {
	jwtServices          JWTService
	populationRepository repository.PopulationRepository
}

type AuthenticationService interface {
	Authenticate(citizenID string, lazerID string) (string, error)
}

var ErrUserNotFound = errors.New("user is not found or malformed")

func NewAuthenticationService(jwtServices JWTService, populationRepository repository.PopulationRepository) AuthenticationService {
	return &authenticationService{
		jwtServices:          jwtServices,
		populationRepository: populationRepository,
	}
}

func (a *authenticationService) Authenticate(citizenID string, lazerID string) (string, error) {
	populationInfo, err := a.populationRepository.GetPopulationInfoBasedOnCitizenIDAndLazerID(citizenID, lazerID)
	if err != nil {
		logrus.Error(ErrUserNotFound)
		return "", ErrUserNotFound
	}

	token, err := a.jwtServices.GenerateToken(fmt.Sprint(populationInfo.CitizenID))
	return token, err
}

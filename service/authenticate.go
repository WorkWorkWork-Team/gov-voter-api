package service

import (
	"errors"
	"fmt"

	"github.com/WorkWorkWork-Team/gov-voter-api/repository"
)

type authenticationService struct {
	jwtServices          JWTService
	populationRepository repository.PopulationRepository
}

type AuthenticationService interface {
	Authenticate(citizenID int, lazerID string) (string, error)
}

func NewAuthenticationService(jwtServices JWTService, populationRepository repository.PopulationRepository) AuthenticationService {
	return &authenticationService{
		jwtServices:          jwtServices,
		populationRepository: populationRepository,
	}
}

func (a *authenticationService) Authenticate(citizenID int, lazerID string) (string, error) {
	userInfo, err := a.populationRepository.GetUserInfoBasedOnCitizenIDAndLazerID(citizenID, lazerID)
	if err != nil {
		return "", errors.New("user is not found or malformed")
	}

	token, err := a.jwtServices.GenerateToken(fmt.Sprint(userInfo.CitizenID))
	return token, err
}

package service

import "github.com/sirupsen/logrus"

type authenticationService struct {
	jwtServices JWTService
}

type AuthenticationService interface {
	Authenticate(citizenID string, lazerID string)
}

func NewAuthenticationService(jwtServices JWTService) AuthenticationService {
	return &authenticationService{
		jwtServices: jwtServices,
	}
}

func (a *authenticationService) Authenticate(citizenID string, lazerID string) {
	logrus.Info("HelloWorld")
}

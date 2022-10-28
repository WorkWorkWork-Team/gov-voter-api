package service

import (
	"errors"

	"github.com/WorkWorkWork-Team/gov-voter-api/repository"
	"github.com/sirupsen/logrus"
)

type validityService struct {
	applyVoteRepository repository.ApplyVoteRepository
}

type ValidityService interface {
	CheckValidity(citizenID string) bool
}

func NewValidityService(applyVoteRepository repository.ApplyVoteRepository) ValidityService {
	return &validityService{
		applyVoteRepository: applyVoteRepository,
	}
}

func (v *validityService) CheckValidity(citizenID string) bool {
	logrus.Info("Start CheckValidity")
	defer logrus.Info("End CheckValidity")

	_, err := v.applyVoteRepository.GetApplyVoteByCitizenID(citizenID)
	logrus.Info("GetApplyVoteByCitizenID focused err: ", err)
	return errors.Is(err, repository.ErrNotFound)
}

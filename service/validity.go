package service

import "github.com/WorkWorkWork-Team/gov-voter-api/repository"

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
	return false
}

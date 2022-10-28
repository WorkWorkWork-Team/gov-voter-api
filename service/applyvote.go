package service

import (
	"github.com/WorkWorkWork-Team/gov-voter-api/repository"
)

type applyvoteService struct {
	applyvoteRepository repository.ApplyVoteRepository
}

type ApplyvoteService interface {
	ApplyVote(citizenID string) error
}

func NewApplyvoteService(applyvoteRepository repository.ApplyVoteRepository) ApplyvoteService {
	return &applyvoteService{
		applyvoteRepository: applyvoteRepository,
	}
}

func (a *applyvoteService) ApplyVote(citizenID string) error {
	return a.applyvoteRepository.ApplyVoteToDB(citizenID)
}

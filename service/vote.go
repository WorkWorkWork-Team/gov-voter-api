package service

import (
	"errors"

	"github.com/WorkWorkWork-Team/gov-voter-api/repository"
	"github.com/sirupsen/logrus"
)

type voteService struct {
	applyVoteRepository repository.ApplyVoteRepository
}

type VoteService interface {
	ApplyVote(citizenID string) error
	CheckValidity(citizenID string) bool
}

var ErrUserAlreadyApplied error = errors.New("user already applied vote")

func NewVoteService(applyVoteRepository repository.ApplyVoteRepository) VoteService {
	return &voteService{
		applyVoteRepository: applyVoteRepository,
	}
}

func (v *voteService) ApplyVote(citizenID string) error {
	logrus.Info("Start apply vote")
	defer logrus.Info("End apply vote")

	isUserCanVoted := v.CheckValidity(citizenID)
	logrus.Info("Is user can vote: ", isUserCanVoted)
	if !isUserCanVoted {
		return ErrUserAlreadyApplied
	}
	err := v.applyVoteRepository.ApplyVote(citizenID)
	return err
}

func (v *voteService) CheckValidity(citizenID string) bool {
	logrus.Info("Start CheckValidity")
	defer logrus.Info("End CheckValidity")

	_, err := v.applyVoteRepository.GetApplyVoteByCitizenID(citizenID)
	logrus.Info("GetApplyVoteByCitizenID focused err: ", err)
	return errors.Is(err, repository.ErrNotFound)
}

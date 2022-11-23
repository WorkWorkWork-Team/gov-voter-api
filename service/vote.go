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
	ApplyVote(citizenID string, tableID int) error
	CheckValidity(citizenID string, tableID int) (bool, error)
	mapTable(tableID int) (string, error)
}

var ErrWrongTableID error = errors.New("tableID not found")
var ErrUserAlreadyApplied error = errors.New("user already applied vote")

func NewVoteService(applyVoteRepository repository.ApplyVoteRepository) VoteService {
	return &voteService{
		applyVoteRepository: applyVoteRepository,
	}
}

func (v *voteService) ApplyVote(citizenID string, tableID int) error {
	logrus.Info("Start apply vote")
	defer logrus.Info("End apply vote")

	isUserCanVoted, err := v.CheckValidity(citizenID, tableID)
	if err != nil {
		return err
	}

	tableName, _ := v.mapTable(tableID)
	logrus.Info("Is user can vote: ", isUserCanVoted)
	if !isUserCanVoted {
		return ErrUserAlreadyApplied
	}

	err = v.applyVoteRepository.ApplyVote(citizenID, tableName)
	return err
}

func (v *voteService) CheckValidity(citizenID string, tableID int) (bool, error) {
	logrus.Info("Start CheckValidity")
	defer logrus.Info("End CheckValidity")
	tableName, err := v.mapTable(tableID)
	if err != nil {
		return false, err
	}
	_, err = v.applyVoteRepository.GetApplyVoteByCitizenID(citizenID, tableName)
	if err != nil && err != repository.ErrNotFound {
		return false, err
	}
	logrus.Info("GetApplyVoteByCitizenID focused err: ", err)
	return errors.Is(err, repository.ErrNotFound), nil
}

func (v *voteService) mapTable(tableID int) (string, error) {
	logrus.Info("Start mapping table")
	defer logrus.Info("Mapping table successful")
	switch tableID {
	case 1:
		return "ApplyVoteMp", nil
	case 2:
		return "ApplyVoteParty", nil
	default:
		return "", ErrWrongTableID
	}
}

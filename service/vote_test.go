package service_test

import (
	"errors"
	model "github.com/WorkWorkWork-Team/gov-voter-api/models"
	"github.com/WorkWorkWork-Team/gov-voter-api/repository"
	"github.com/WorkWorkWork-Team/gov-voter-api/service"
	"github.com/WorkWorkWork-Team/gov-voter-api/test/mock_repository"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Vote", func() {
	var ctrl *gomock.Controller
	var mockVoteRepository *mock_repository.MockApplyVoteRepository
	var voteService service.VoteService

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockVoteRepository = mock_repository.NewMockApplyVoteRepository(ctrl)
		voteService = service.NewVoteService(mockVoteRepository)
	})

	Describe("Apply vote", func() {
		Context("With right condition", func() {
			BeforeEach(func() {
				mockVoteRepository.EXPECT().
					ApplyVote(gomock.Any()).
					Return(nil)
				mockVoteRepository.EXPECT().
					GetApplyVoteByCitizenID(gomock.Any()).
					Return(
						model.ApplyVote{
							ID: 1,
						}, repository.ErrNotFound)
			})
			It("Should not return error", func() {
				Expect(voteService.ApplyVote("")).Should(BeNil())
			})
		})

		Context("With user already voted", func() {
			BeforeEach(func() {
				mockVoteRepository.EXPECT().
					ApplyVote(gomock.Any()).
					Return(nil).Times(0)
				mockVoteRepository.EXPECT().
					GetApplyVoteByCitizenID(gomock.Any()).
					Return(
						model.ApplyVote{
							ID: 1,
						}, nil).AnyTimes()
			})
			It("Should return user already applied error", func() {
				Expect(voteService.ApplyVote("")).Should(Equal(service.ErrUserAlreadyApplied))
			})
		})

		Context("Failed to insert user to database", func() {
			BeforeEach(func() {
				mockVoteRepository.EXPECT().
					ApplyVote(gomock.Any()).
					Return(errors.New(""))
				mockVoteRepository.EXPECT().
					GetApplyVoteByCitizenID(gomock.Any()).
					Return(
						model.ApplyVote{
							ID: 1,
						}, repository.ErrNotFound).AnyTimes()
			})
			It("Should return user already applied error", func() {
				Expect(voteService.ApplyVote("")).Should(Equal(errors.New("")))
			})
		})
	})
})

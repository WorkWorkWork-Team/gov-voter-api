package service_test

import (
	"database/sql"

	model "github.com/WorkWorkWork-Team/gov-voter-api/models"
	"github.com/WorkWorkWork-Team/gov-voter-api/service"
	"github.com/WorkWorkWork-Team/gov-voter-api/test/mock_repository"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var (
	TestCitizenID string = "1234"
)

var _ = Describe("Vote", Label("unit"), func() {
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
					ApplyVote(gomock.Eq(TestCitizenID)).
					Return(nil)
				mockVoteRepository.EXPECT().
					GetApplyVoteByCitizenID(gomock.Eq(TestCitizenID)).
					Return(
						model.ApplyVote{
							ID: 1,
						}, sql.ErrNoRows).Times(1)
			})
			It("Should not return error", func() {
				Expect(voteService.ApplyVote(TestCitizenID)).Should(BeNil())
			})
		})

		Context("With user already voted", func() {
			BeforeEach(func() {
				mockVoteRepository.EXPECT().
					ApplyVote(gomock.Eq(TestCitizenID)).
					Return(nil).Times(0)
				mockVoteRepository.EXPECT().
					GetApplyVoteByCitizenID(gomock.Eq(TestCitizenID)).
					Return(
						model.ApplyVote{
							ID: 1,
						}, nil).Times(1)
			})
			It("Should return user already applied error", func() {
				Expect(voteService.ApplyVote(TestCitizenID)).Should(Equal(service.ErrUserAlreadyApplied))
			})
		})

		Context("Failed to insert user to database", func() {
			BeforeEach(func() {
				mockVoteRepository.EXPECT().
					ApplyVote(gomock.Eq(TestCitizenID)).
					Return(sql.ErrConnDone)
				mockVoteRepository.EXPECT().
					GetApplyVoteByCitizenID(gomock.Eq(TestCitizenID)).
					Return(
						model.ApplyVote{
							ID: 1,
						}, sql.ErrNoRows).Times(1)
			})
			It("Should return user already applied error", func() {
				Expect(voteService.ApplyVote(TestCitizenID)).Should(Equal(sql.ErrConnDone))
			})
		})
	})
})

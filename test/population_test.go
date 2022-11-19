package test

import (
	model "github.com/WorkWorkWork-Team/gov-voter-api/models"
	"github.com/WorkWorkWork-Team/gov-voter-api/repository"
	"github.com/WorkWorkWork-Team/gov-voter-api/service"
	"github.com/WorkWorkWork-Team/gov-voter-api/test/mock_repository"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Population", func() {
	var control *gomock.Controller
	var mockPopulationRepository *mock_repository.MockPopulationRepository
	var populationService service.PopulationService
	BeforeEach(func() {
		control = gomock.NewController(GinkgoT())
		mockPopulationRepository = mock_repository.NewMockPopulationRepository(control)
		populationService = service.NewPopulationService(mockPopulationRepository)
	})

	Describe("Get user's information", func() {
		Context("Success", func() {
			BeforeEach(func() {
				mockPopulationRepository.EXPECT().GetPopulationInfo(gomock.Any()).Return(
					model.Population{
						CitizenID: 1234,
						LazerId:   "1234",
					}, repository.ErrNotFound)
			})
			It("Should return user information with out error", func() {
				Expect(populationService.GetPopulationInformation("")).Should(BeNil())
			})
		})
	})

	Describe("Fail to find user's information", func() {
		BeforeEach(func() {
			mockPopulationRepository.EXPECT().GetPopulationInfo(gomock.Any()).Return(
				model.Population{
					CitizenID: 0,
					LazerId:   "",
				}, repository.ErrNotFound)
		})
		It("Should return base model data with error", func() {
			Expect(populationService.GetPopulationInformation("")).Should(Equal(service.ErrUserNotFound))
		})
	})
})

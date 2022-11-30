package service_test

import (
	"time"

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
						CitizenID:   1234,
						LazerID:     "1234",
						Name:        "name",
						Lastname:    "lastname",
						Birthday:    time.Now(),
						Nationality: "thai",
						DistricID:   123145,
					}, nil)
			})
			It("Should return user information with out error", func() {
				_, err := populationService.GetPopulationInformation("1234")
				Expect(err).Should(BeNil())
			})
		})
	})

	Describe("Fail to find user's information", func() {
		BeforeEach(func() {
			mockPopulationRepository.EXPECT().GetPopulationInfo(gomock.Any()).Return(
				model.Population{
					CitizenID:   1234,
					LazerID:     "1234",
					Name:        "name",
					Lastname:    "lastname",
					Birthday:    time.Now(),
					Nationality: "thai",
					DistricID:   123145,
				}, repository.ErrNotFound)
		})
		It("Should return base model data with error", func() {
			_, err := populationService.GetPopulationInformation("1234")
			Expect(err).Should(Equal(repository.ErrNotFound))
		})
	})
})

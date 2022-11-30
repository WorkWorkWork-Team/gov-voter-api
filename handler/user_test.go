package handler_test

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/WorkWorkWork-Team/gov-voter-api/handler"
	model "github.com/WorkWorkWork-Team/gov-voter-api/models"
	"github.com/WorkWorkWork-Team/gov-voter-api/repository"
	"github.com/WorkWorkWork-Team/gov-voter-api/service"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var (
	TestUserCitizenID     string = "1234567891235"
	TestUserLazerID       string = "CCAADD"
	TestJWTSecretKey      string = "key"
	TestJWTIssuer         string = "Tester"
	TestJWTTTL            int    = 10
	UserHandler           *handler.UserHandler
	AuthenticateHandler   *handler.AuthenticateHandler
	AuthenticationService service.AuthenticationService
)

var _ = Describe("User Integration Test", Label("integration"), func() {
	BeforeEach(func() {
		gin.SetMode(gin.TestMode)
		populationRepository := repository.NewPopulationRepository(MySQLConnection)
		applyVoteRepository := repository.NewApplyVoteRepository(MySQLConnection)

		// New Services

		jwtService := service.NewJWTService(TestJWTSecretKey, TestJWTIssuer, time.Duration(TestJWTTTL)*time.Second)
		voteService := service.NewVoteService(applyVoteRepository)
		AuthenticationService = service.NewAuthenticationService(jwtService, populationRepository)
		populationService := service.NewPopulationService(populationRepository)

		// New Handler
		UserHandler = handler.NewUserHandler(populationService, jwtService, voteService)
		AuthenticateHandler = handler.NewAuthenticateHandler(AuthenticationService)
	})

	Context("Validity API", func() {
		Context("Database have population data", func() {
			When("the user is not in the voted table", func() {
				It("should return success.", func() {
					// Expect no user in voted table
					var applyVoteList []model.ApplyVote
					err := MySQLConnection.Select(&applyVoteList, "SELECT * FROM `ApplyVote` WHERE citizenID=?", TestUserCitizenID)
					Expect(err).ShouldNot(HaveOccurred())
					applyVoteLength := len(applyVoteList)
					Expect(applyVoteLength).To(Equal(0))

					// Call API
					res := httptest.NewRecorder()
					c, _ := gin.CreateTestContext(res)
					c.AddParam("CitizenID", TestUserCitizenID)
					UserHandler.Validity(c)

					// Expect API return 200
					Expect(c.Writer.Status()).To(Equal(http.StatusOK))
				})
			})

			When("the user is in the voted table", func() {
				It("should return 400 Unsuccess.", func() {
					// Expect no user in voted table
					var applyVote model.ApplyVote
					err := MySQLConnection.Get(&applyVote, "SELECT * FROM `ApplyVote` WHERE CitizenID=?", TestUserCitizenID)
					Expect(err).To(Equal(sql.ErrNoRows))

					// Insert user to voted table
					_, err = MySQLConnection.Exec("INSERT INTO `ApplyVote` (CitizenID) VALUES (?)", TestUserCitizenID)
					Expect(err).To(BeNil())

					// Call API
					res := httptest.NewRecorder()
					c, _ := gin.CreateTestContext(res)
					c.AddParam("CitizenID", TestUserCitizenID)
					UserHandler.Validity(c)

					// Expect API return 400
					Expect(c.Writer.Status()).To(Equal(http.StatusBadRequest))

					// Clear user from ApplyVote table
					_, err = MySQLConnection.Exec("DELETE FROM `ApplyVote` WHERE CitizenID=?", TestUserCitizenID)
					Expect(err).To(BeNil())
				})
			})
		})
	})
})

package handler_test

import (
	"database/sql"
	"fmt"
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
	TestUserCitizenID   string = "1234567891235"
	TestJWTSecretKey    string = "key"
	TestJWTIssuer       string = "Tester"
	TestJWTTTL          int    = 10
	UserHandler         *handler.UserHandler
	AuthenticateHandler *handler.AuthenticateHandler
)

var _ = Describe("User Integration Test", Label("integration"), func() {
	BeforeEach(func() {
		populationRepository := repository.NewPopulationRepository(MySQLConnection)
		applyVoteRepository := repository.NewApplyVoteRepository(MySQLConnection)

		// New Services

		jwtService := service.NewJWTService(TestJWTSecretKey, TestJWTIssuer, time.Duration(TestJWTTTL)*time.Second)
		voteService := service.NewVoteService(applyVoteRepository)
		authenticationService := service.NewAuthenticationService(jwtService, populationRepository)
		populationService := service.NewPopulationService(populationRepository)

		// New Handler
		UserHandler = handler.NewUserHandler(populationService, jwtService, voteService)
		AuthenticateHandler = handler.NewAuthenticateHandler(authenticationService)
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
					c.Request = httptest.NewRequest(http.MethodPost, "/api", nil)
					UserHandler.Validity(c)

					// Expect API return 200
					Expect(res.Result().StatusCode).To(Equal(http.StatusOK))
				})
			})

			When("the user is in the voted table", func() {
				It("should return 400 Unsuccess.", func() {
					// Expect no user in voted table
					var applyVote model.ApplyVote
					err := MySQLConnection.Get(&applyVote, "SELECT * FROM `ApplyVote` WHERE citizenID=?", TestUserCitizenID)
					Expect(err).To(Equal(sql.ErrNoRows))

					// Insert user to voted table
					row, err := MySQLConnection.Exec("INSERT INTO `ApplyVote` (citizenID) VALUES (?)", TestUserCitizenID)
					fmt.Println(row)
					fmt.Println(err)
					Expect(err).To(BeNil())

					// Call API
					res := httptest.NewRecorder()
					c, _ := gin.CreateTestContext(res)
					c.Request = httptest.NewRequest(http.MethodPost, "/api", nil)
					UserHandler.Validity(c)

					// Expect API return 400
					Expect(res.Result().StatusCode).To(Equal(http.StatusBadRequest))
				})
			})
		})
	})
})

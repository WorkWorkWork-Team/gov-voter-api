package handler_test

import (
	"database/sql"
	"encoding/json"
	"io"
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
					//row, err := MySQLConnection.Exec("INSERT INTO `ApplyVote` (citizenID) VALUES (?)", TestUserCitizenID)
					//fmt.Println(row)
					//fmt.Println(err)
					//Expect(err).To(BeNil())

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

	Context("Applyvote API", func() {
		Context("Database have population data", func() {
			When("the user is not in the voted table", func() {
				It("Should return success", func() {
					// Check Condition
					var applyVoteList []model.ApplyVote
					var populationList []model.Population
					err := MySQLConnection.Select(&applyVoteList, "SELECT * FROM `ApplyVote` WHERE citizenID=?", TestUserCitizenID)
					Expect(err).ShouldNot(HaveOccurred())
					applyVoteLength := len(applyVoteList)
					Expect(applyVoteLength).To(Equal(0))

					err = MySQLConnection.Select(&populationList, "SELECT * FROM `Population` WHERE citizenID=?", TestUserCitizenID)
					Expect(err).ShouldNot(HaveOccurred())
					populationLength := len(populationList)
					Expect(populationLength).To(Equal(1))

					// Call API
					res := httptest.NewRecorder()
					c, _ := gin.CreateTestContext(res)
					c.AddParam("CitizenID", TestUserCitizenID)
					c.Request = httptest.NewRequest(http.MethodPost, "/api", nil)
					UserHandler.ApplyVote(c)

					Expect(res.Result().StatusCode).To(Equal(http.StatusOK))

					// Expect no user in voted table.
					err = MySQLConnection.Select(&applyVoteList, "SELECT * FROM `ApplyVote` WHERE citizenID=?", TestUserCitizenID)
					Expect(err).ShouldNot(HaveOccurred())
					applyVoteLength = len(applyVoteList)
					Expect(applyVoteLength).To(Equal(1))

					// Tear down
					_, err = MySQLConnection.Exec("DELETE FROM `ApplyVote` WHERE citizenID=?", TestUserCitizenID)
					Expect(err).ShouldNot(HaveOccurred())

					err = MySQLConnection.Select(&applyVoteList, "SELECT * FROM `ApplyVote` WHERE citizenID=?", TestUserCitizenID)
					Expect(err).ShouldNot(HaveOccurred())
					applyVoteLength = len(applyVoteList)
					Expect(applyVoteLength).To(Equal(0))
				})
			})
			When("the user is in the voted table", func() {
				It("Should return bad request", func() {
					var applyVoteList []model.ApplyVote
					var populationList []model.Population

					_, err := MySQLConnection.Exec("INSERT INTO `ApplyVote` (citizenID) VALUES (?)", TestUserCitizenID)
					Expect(err).ShouldNot(HaveOccurred())

					err = MySQLConnection.Select(&applyVoteList, "SELECT * FROM `ApplyVote` WHERE citizenID=?", TestUserCitizenID)
					Expect(err).ShouldNot(HaveOccurred())
					applyVoteLength := len(applyVoteList)
					Expect(applyVoteLength).To(Equal(1))

					err = MySQLConnection.Select(&populationList, "SELECT * FROM `Population` WHERE citizenID=?", TestUserCitizenID)
					Expect(err).ShouldNot(HaveOccurred())
					populationLength := len(populationList)
					Expect(populationLength).To(Equal(1))

					// API Calls
					res := httptest.NewRecorder()
					c, _ := gin.CreateTestContext(res)
					c.AddParam("CitizenID", TestUserCitizenID)
					c.Request = httptest.NewRequest(http.MethodPost, "/api", nil)
					UserHandler.ApplyVote(c)

					Expect(res.Result().StatusCode).To(Equal(http.StatusBadRequest))

					// Tear down
					_, err = MySQLConnection.Exec("DELETE FROM `ApplyVote` WHERE citizenID=?", TestUserCitizenID)
					Expect(err).ShouldNot(HaveOccurred())

					err = MySQLConnection.Select(&applyVoteList, "SELECT * FROM `ApplyVote` WHERE citizenID=?", TestUserCitizenID)
					Expect(err).ShouldNot(HaveOccurred())
					applyVoteLength = len(applyVoteList)
					Expect(applyVoteLength).To(Equal(0))
				})
			})
		})
	})

	Context("GetUserInfo API", func() {
		Context("Database have population data", func() {
			When("get user information", func() {
				It("Should return success with body", func() {
					var populationList []model.Population
					var populationInfo model.Population

					err := MySQLConnection.Select(&populationList, "SELECT * FROM `Population` WHERE citizenID=?", TestUserCitizenID)
					Expect(err).ShouldNot(HaveOccurred())
					populationLength := len(populationList)
					Expect(populationLength).To(Equal(1))

					// Call API
					res := httptest.NewRecorder()
					c, _ := gin.CreateTestContext(res)
					c.AddParam("CitizenID", TestUserCitizenID)
					c.Request = httptest.NewRequest(http.MethodPost, "/api", nil)
					UserHandler.GetUserInfo(c)
					Expect(res.Result().StatusCode).To(Equal(http.StatusOK))

					// Get response body
					body, err := io.ReadAll(res.Body)
					Expect(err).ShouldNot(HaveOccurred())
					err = json.Unmarshal(body, &populationInfo)
					Expect(err).ShouldNot(HaveOccurred())

					Expect(populationInfo).To(Equal(populationList[0]))

				})
			})
		})
	})
})

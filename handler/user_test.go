package handler_test

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"

	"github.com/WorkWorkWork-Team/gov-voter-api/handler"
	model "github.com/WorkWorkWork-Team/gov-voter-api/models"
	"github.com/WorkWorkWork-Team/gov-voter-api/repository"
	"github.com/WorkWorkWork-Team/gov-voter-api/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var (
	TestUserCitizenID     string = "1234567891235"
	TestNotExistCitizenID string = "1"
	TestUserLazerID       string = "CCAADD"
	TestJWTSecretKey      string = "key"
	TestJWTIssuer         string = "Tester"
	TestJWTTTL            int    = 10
	UserHandler           *handler.UserHandler
	AuthenticateHandler   *handler.AuthenticateHandler
	AuthenticationService service.AuthenticationService
	JWTService            service.JWTService
)

func NewGinTestContext() (*gin.Context, *httptest.ResponseRecorder, *gin.Engine) {
	res := httptest.NewRecorder()
	c, r := gin.CreateTestContext(res)
	return c, res, r
}

type AuthToken struct {
	Token string `json:"token"`
}

var _ = Describe("User Integration Test", Label("integration"), func() {
	BeforeEach(func() {
		gin.SetMode(gin.TestMode)
		populationRepository := repository.NewPopulationRepository(MySQLConnection)
		applyVoteRepository := repository.NewApplyVoteRepository(MySQLConnection)

		// New Services

		JWTService = service.NewJWTService(TestJWTSecretKey, TestJWTIssuer, time.Duration(TestJWTTTL)*time.Second)
		voteService := service.NewVoteService(applyVoteRepository)
		AuthenticationService = service.NewAuthenticationService(JWTService, populationRepository)
		populationService := service.NewPopulationService(populationRepository)

		// New Handler
		UserHandler = handler.NewUserHandler(populationService, JWTService, voteService)
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
					c, _, _ := NewGinTestContext()
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
					c, _, _ := NewGinTestContext()
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

	Context("Authentication API", func() {
		When("user provide incorrect user credential", func() {
			It("should return 401", func() {
				_, resultWriter, r := NewGinTestContext()
				r.POST("/api", AuthenticateHandler.AuthAndGenerateToken)
				body := fmt.Sprintf(`{
					"citizenID": "%s",
					"lazerID": "I'mSurlyFailed"
				}`, TestUserCitizenID)
				reader := strings.NewReader(body)
				req, _ := http.NewRequest("POST", "/api", reader)
				r.ServeHTTP(resultWriter, req)

				Expect(resultWriter.Result().StatusCode).To(Equal(http.StatusUnauthorized))
			})
		})
		When("user provide correct user credential", func() {
			It("should return 200", func() {
				_, resultWriter, r := NewGinTestContext()
				r.POST("/api", AuthenticateHandler.AuthAndGenerateToken)
				body := fmt.Sprintf(`{
					"citizenID": "%s",
					"lazerID": "%s"
				}`, TestUserCitizenID, TestUserLazerID)
				reader := strings.NewReader(body)
				req, _ := http.NewRequest("POST", "/api", reader)
				r.ServeHTTP(resultWriter, req)

				var result AuthToken
				resultByte, _ := io.ReadAll(resultWriter.Result().Body)
				err := json.Unmarshal(resultByte, &result)
				Expect(err).To(BeNil())

				Expect(resultWriter.Result().StatusCode).To(Equal(http.StatusOK))
				validatedToken, err := JWTService.ValidateToken(result.Token)
				Expect(err).To(BeNil())
				Expect(validatedToken.Valid).To(BeTrue())
				Expect(validatedToken.Claims.(jwt.MapClaims)["iss"]).To(Equal(TestJWTIssuer))
				Expect(validatedToken.Claims.(jwt.MapClaims)["exp"]).To(BeNumerically("==", time.Now().Add(time.Duration(TestJWTTTL)*time.Second).Unix()))
				Expect(validatedToken.Claims.(jwt.MapClaims)["CitizenID"]).To(Equal(TestUserCitizenID))
				Expect(validatedToken.Claims.(jwt.MapClaims)["iat"]).To(BeNumerically("==", time.Now().Unix()))
			})
		})

		When("user provide incorrect user credential", func() {
			It("should return 401", func() {
				_, resultWriter, r := NewGinTestContext()
				r.POST("/api", AuthenticateHandler.AuthAndGenerateToken)
				body := fmt.Sprintf(`{
					"citizenID": "%s",
					"lazerID": "I'mSurlyFailed"
				}`, TestUserCitizenID)
				reader := strings.NewReader(body)
				req, _ := http.NewRequest("POST", "/api", reader)
				r.ServeHTTP(resultWriter, req)

				Expect(resultWriter.Result().StatusCode).To(Equal(http.StatusUnauthorized))
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
					c, res, _ := NewGinTestContext()
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
					c, res, _ := NewGinTestContext()
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
					c, res, _ := NewGinTestContext()
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
			When("get user information with invalid citizen ID", func() {
				It("Should return not found", func() {
					var populationList []model.Population

					err := MySQLConnection.Select(&populationList, "SELECT * FROM `Population` WHERE citizenID=?", TestNotExistCitizenID)
					Expect(err).ShouldNot(HaveOccurred())
					populationLength := len(populationList)
					Expect(populationLength).To(Equal(0))

					// Call API
					c, res, _ := NewGinTestContext()
					c.AddParam("CitizenID", TestNotExistCitizenID)
					c.Request = httptest.NewRequest(http.MethodPost, "/api", nil)
					UserHandler.GetUserInfo(c)
					Expect(res.Result().StatusCode).To(Equal(http.StatusNotFound))
				})
			})
		})
	})
})

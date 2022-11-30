package handler_test

import (
	model "github.com/WorkWorkWork-Team/gov-voter-api/models"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var (
	TestUserCitizenID string = "1234567891234"
)

var _ = Describe("User Integration Test", Label("integration"), func() {
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

					// TODO: Call API
					Expect(true).To(BeTrue())

					// TODO: Expect API return 200

					// TODO: Expect no user in voted table.

				})
			})
		})
	})
})

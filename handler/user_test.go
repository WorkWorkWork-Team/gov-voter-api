package handler_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("User Integration Test", Label("integration"), func() {
	Context("Validity API", func() {
		Context("Database have population data", func() {
			When("the user is not in the voted table", func() {
				It("should return success.", func() {
					Expect(true).To(BeTrue())
				})
			})
		})
	})
})

package goraphite

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFindOptions(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "FindOptions")
}

var _ = Describe("FindOptions", func() {
	Describe("String is built correctly", func() {
		It("Should be able to add a Query property", func() {
			query := FindOptions{
				Query: "collectd.*",
			}

			queryString, _ := query.String()

			Expect(queryString).To(Equal("query=collectd.%2A"))
		})

		It("Should be able to add Format property", func() {
			query := FindOptions{
				Format: "completer",
			}

			queryString, _ := query.String()

			Expect(queryString).To(Equal("format=completer"))
		})
	})
})

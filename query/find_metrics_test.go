package query

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFindMetrics(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "FindMetrics")
}

var _ = Describe("FindMetrics", func() {
	Describe("String is built correctly", func() {
		It("Should be able to add a Query property", func() {
			query := FindMetrics{
				Query: "collectd.*",
			}

			queryString, _ := query.String()

			Expect(queryString).To(Equal("query=collectd.%2A"))
		})

		It("Should be able to add Format property", func() {
			query := FindMetrics{
				Format: "completer",
			}

			queryString, _ := query.String()

			Expect(queryString).To(Equal("format=completer"))
		})

		It("Should create a proper querystring when more than one value is passed", func() {
			query := FindMetrics{
				Query:  "collectd.*",
				Format: "completer",
			}

			queryString, _ := query.String()

			Expect(queryString).To(Equal("format=completer&query=collectd.%2A"))
		})
	})
})

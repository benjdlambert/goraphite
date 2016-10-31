package query

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGetMetrics(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GetMetrics")
}

var _ = Describe("GetMetrics", func() {
	Describe("String is built correctly", func() {
		It("Should be able to add a Format property", func() {
			query := GetMetrics{
				Format: "json",
			}

			queryString, _ := query.String()

			Expect(queryString).To(Equal("format=json"))
		})

		It("Should be able to add a Target property", func() {
			query := GetMetrics{
				Target: "collectd.*",
			}

			queryString, _ := query.String()

			Expect(queryString).To(Equal("target=collectd.%2A"))
		})

		It("Should create a proper querystring when more than one value is passed", func() {
			query := GetMetrics{
				Target: "collectd.*",
				Format: "json",
			}

			queryString, _ := query.String()

			Expect(queryString).To(Equal("format=json&target=collectd.%2A"))
		})
	})
})

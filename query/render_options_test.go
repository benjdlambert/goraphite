package query

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRenderOptions(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RenderOptions")
}

var _ = Describe("RenderOptions", func() {
	Describe("String is built correctly", func() {
		It("Should be able to add a Format property", func() {
			query := RenderOptions{
				Format: "json",
			}

			queryString, _ := query.String()

			Expect(queryString).To(Equal("format=json"))
		})

		It("Should be able to add a Target property", func() {
			query := RenderOptions{
				Target: "collectd.*",
			}

			queryString, _ := query.String()

			Expect(queryString).To(Equal("target=collectd.%2A"))
		})

		It("Should create a proper querystring when more than one value is passed", func() {
			query := RenderOptions{
				Target: "collectd.*",
				Format: "json",
			}

			queryString, _ := query.String()

			Expect(queryString).To(Equal("format=json&target=collectd.%2A"))
		})
	})
})

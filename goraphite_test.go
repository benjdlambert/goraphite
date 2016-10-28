package goraphite

import (
	"reflect"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGoraphite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Goraphite Suite")
}

var _ = Describe("Goraphite", func() {
	var (
		host string
		port int
	)

	BeforeEach(func() {
		host, port = "mock.host.com", 1234
	})

	Describe("Creating a Goraphite Client", func() {
		Context("With valid parameters", func() {
			It("should return an instance of a Client with no error", func() {
				client, err := NewGoraphite(host, port)
				Expect(reflect.TypeOf(client).String()).To(Equal("*goraphite.Client"))
				Expect(err).To(BeNil())
			})
		})

		Context("With invalid parameters", func() {
			It("Should return an error when port is invalid", func() {
				port = -1
				client, err := NewGoraphite(host, port)
				Expect(client).To(BeNil())
				Expect(err).Should(HaveOccurred())
			})

			It("Should return an error when host is not valid", func() {
				host = ""
				client, err := NewGoraphite(host, port)
				Expect(client).To(BeNil())
				Expect(err).Should(HaveOccurred())
			})
		})
	})
})

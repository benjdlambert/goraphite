package goraphite

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/h2non/gock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGoraphite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Goraphite Suite")
}

var _ = Describe("Goraphite", func() {
	var (
		host   string
		port   int
		client *Client
	)

	BeforeEach(func() {
		host, port = "mock.host.com", 1234
		client, _ = NewGoraphite(host, port)
	})

	AfterEach(func() {
		gock.Off()
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

	Describe("Status", func() {
		Context("Status enpoint returns 200", func() {
			BeforeEach(func() {
				gock.New(fmt.Sprintf("http://%s:%d", host, port)).
					Get("/").
					Reply(200)
			})

			It("Should return a Status struct with the response code", func() {
				status, err := client.Status()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(reflect.TypeOf(status).String()).To(Equal("*goraphite.Status"))
				Expect(status.Code).To(Equal(200))
			})
		})

		Context("Status endpoint returns non 200", func() {
			BeforeEach(func() {
				gock.New(fmt.Sprintf("http://%s:%d", host, port)).
					Get("/").
					Reply(500)
			})

			It("Should return an error", func() {
				status, err := client.Status()
				Expect(err).Should(HaveOccurred())
				Expect(status).To(BeNil())
			})
		})
	})
})

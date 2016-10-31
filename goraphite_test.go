package goraphite

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/benjdlambert/goraphite/query"
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
				Expect(err).To(HaveOccurred())
			})

			It("Should return an error when host is not valid", func() {
				host = ""
				client, err := NewGoraphite(host, port)
				Expect(client).To(BeNil())
				Expect(err).To(HaveOccurred())
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
				Expect(err).ToNot(HaveOccurred())
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
				Expect(err).To(HaveOccurred())
				Expect(status).To(BeNil())
			})
		})
	})

	Describe("FindMetrics", func() {
		Context("/metrics/find", func() {
			BeforeEach(func() {
				gock.New(fmt.Sprintf("http://%s:%d", host, port)).
					Get("/metrics/find").
					MatchParam("query", "collectd").
					Reply(200).
					JSON(`[{
                        "leaf": 0,
                        "context": {},
                        "text": "test1",
                        "expandable": 1,
                        "id": "collectd.test1",
                        "allowChildren": 1
                    }]`)
			})

			It("Should return the amount of items back correctly", func() {
				metrics, err := client.FindMetrics(
					query.FindMetrics{
						Query: "collectd",
					},
				)

				Expect(err).ToNot(HaveOccurred())
				Expect(len(*metrics)).To(Equal(1))
			})

			It("Should have unmarshalled the response properly", func() {
				metrics, _ := client.FindMetrics(
					query.FindMetrics{
						Query: "collectd",
					},
				)

				testMetric := (*metrics)[0]

				Expect(testMetric.ID).To(Equal("collectd.test1"))
				Expect(testMetric.IsLeaf()).To(Equal(false))
				Expect(testMetric.IsExpandable()).To(Equal(true))
				Expect(testMetric.AllowsChildren()).To(Equal(true))
			})
		})
	})

	Describe("GetMetrics", func() {
		Context("/render", func() {
			BeforeEach(func() {
				gock.New(fmt.Sprintf("http://%s:%d", host, port)).
					Get("/render").
					MatchParam("format", "json").
					MatchParam("target", "collectd.*").
					Reply(200).
					JSON(`[{
                        "target": "collectd.1",
                        "datapoints": [
                            [1,2], [3,4], [5,6]
                        ]
                    }]`)
			})

			It("Should return the amount of metrics back properly", func() {
				metrics, err := client.GetMetrics(
					query.GetMetrics{
						Target: "collectd.*",
					},
				)

				Expect(err).ToNot(HaveOccurred())
				Expect(len(*metrics)).To(Equal(1))
			})

			It("Should return the correct amount of datapoints", func() {
				metrics, err := client.GetMetrics(
					query.GetMetrics{
						Target: "collectd.*",
					},
				)
				testMetric := (*metrics)[0]

				Expect(err).ToNot(HaveOccurred())
				Expect(len(testMetric.Datapoints())).To(Equal(3))
			})
		})
	})
})

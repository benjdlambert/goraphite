package goraphite

import (
	"os"
	"reflect"
	"testing"

	"github.com/h2non/gock"
	"github.com/nbio/st"
)

var host string
var port int
var nilClient *Client

func setup() {
	host, port = mockGraphiteDetails()
}

func teardown() {
	gock.Off()
}

func TestMain(m *testing.M) {
	setup()
	retCode := m.Run()
	teardown()
	os.Exit(retCode)
}

func TestNewGoraphite(t *testing.T) {
	_, err := NewGoraphite(host, port)
	st.Expect(t, err, nil)
}

func TestNewGoraphiteReturnsClient(t *testing.T) {
	client, _ := NewGoraphite(host, port)
	st.Expect(t, reflect.TypeOf(client).String(), "*goraphite.Client")
}

func TestErrorWhenInvalidHostProvided(t *testing.T) {
	client, error := NewGoraphite("", port)
	st.Reject(t, error, nil)
	st.Expect(t, client, nilClient)
}

func TestErrorWhenInvalidPortIsProvided(t *testing.T) {
	client, error := NewGoraphite(host, -1)
	st.Reject(t, error, nil)
	st.Expect(t, client, nilClient)
}

func TestStatusCallToGraphite(t *testing.T) {

}

func mockGraphiteDetails() (string, int) {
	return "mock.host.com", 1234
}

package goraphite

import (
	"reflect"
	"testing"
)

var host, port = mockGraphiteDetails()

func TestNewGoraphite(t *testing.T) {
	_, err := NewGoraphite(host, port)

	if err != nil {
		t.Error(err)
	}
}

func TestNewGoraphiteReturnsClient(t *testing.T) {
	client, _ := NewGoraphite(host, port)

	if reflect.TypeOf(client).String() != "*goraphite.Client" {
		t.Errorf("Type is wrong: %s", reflect.TypeOf(client))
	}
}

func TestErrorWhenInvalidHostProvided(t *testing.T) {
	client, error := NewGoraphite("", port)

	if error == nil {
		t.Error("No error was thrown")
	}

	if client != nil {
		t.Error("Client is not nil")
	}
}

func TestErrorWhenInvalidPortIsProvided(t *testing.T) {
	client, error := NewGoraphite(host, 0)

	if error == nil {
		t.Error("No error was thrown")
	}

	if client != nil {
		t.Error("Client is not nil")
	}
}

func mockGraphiteDetails() (string, int) {
	return "mock.host.com", 1234
}

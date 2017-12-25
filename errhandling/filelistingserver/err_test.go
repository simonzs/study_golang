package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func panicErr(writer http.ResponseWriter,
	request *http.Request) error {
	panic("panic message")
}

type testError string

func (e testError) Error() string {
	return e.Message()
}

func (e testError) Message() string {
	return string(e)
}

func userErr(writer http.ResponseWriter,
	request *http.Request) error {
	return testError("this is a user error")
}

func io404Err(writer http.ResponseWriter,
	request *http.Request) error {
	return os.ErrNotExist
}

func io403Err(writer http.ResponseWriter,
	request *http.Request) error {
	return os.ErrPermission
}

func unknownErr(writer http.ResponseWriter,
	request *http.Request) error {
	return errors.New("unknown")
}

var tests = []struct {
	name         string
	handler      appHandler
	expectedMsg  string
	expectedCode int
}{
	{"panic", panicErr, "Internal Server Error", 500},
	{"404", io404Err, "Not Found", 404},
}

func TestErrWrapper(t *testing.T) {
	for _, tt := range tests {
		request := httptest.NewRequest(
			http.MethodGet,
			"http://www.imooc.com", nil)
		resp := httptest.NewRecorder()
		errWrapper(tt.handler)(resp, request)
		b, _ := ioutil.ReadAll(resp.Body)
		actualMsg := strings.Trim(string(b), "\n")
		actualCode := resp.Code
		if actualMsg != tt.expectedMsg ||
			actualCode != tt.expectedCode {
			t.Errorf("Test %s: expect (%s, %d); "+
				"got (%s, %d)",
				tt.name, tt.expectedMsg,
				tt.expectedCode, actualMsg,
				actualCode)
		}
	}
}

func TestErrWrapperServer(t *testing.T) {
	for _, tt := range tests {
		server := httptest.NewServer(
			http.HandlerFunc(
				errWrapper(tt.handler)))

		resp, _ := http.Get(server.URL)
		b, _ := ioutil.ReadAll(resp.Body)
		actualMsg := strings.Trim(string(b), "\n")
		actualCode := resp.StatusCode
		if actualMsg != tt.expectedMsg ||
			actualCode != tt.expectedCode {
			t.Errorf("Test %s: expect (%s, %d); "+
				"got (%s, %d)",
				tt.name, tt.expectedMsg,
				tt.expectedCode, actualMsg,
				actualCode)
		}
	}
}

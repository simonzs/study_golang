package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func errPanic(writer http.ResponseWriter, 
	_ *http.Request) error {
	panic(123)
}

type testingUserError string

func (e testingUserError) Error() string {
	return e.Message()
}

func (e testingUserError) Message() string {
	return string(e)
}

func errUserError(_ http.ResponseWriter,
	_ *http.Request) error {
	return testingUserError("User Error")
}

func errNotFound(_ http.ResponseWriter,
	_ *http.Request) error {
	return os.ErrNotExist
}

func errNoPermission(_ http.ResponseWriter,
	_ *http.Request) error {
	return os.ErrPermission
}

func errUnknow(_ http.ResponseWriter,
	_ *http.Request) error {
	return errors.New("unknown error")
}

func noError(writer http.ResponseWriter,
	_ *http.Request) error {
	fmt.Fprintln(writer, "no error")
	return nil
}


var tests = []struct {
	h apphandler
	code int
	message string
}{
	{errPanic, 500, "Internal Server Error"},
	{errUserError, 400, "User Error"},
	{errNotFound, 404, "Not Found"},
	{errNoPermission, 403, "Forbidden"},
	{errUnknow, 500, "Internal Server Error"},
	{noError, 200, "no error"},
}

// TestErrWrapper 使用假的Response Request 测试      测试的是一段代码
func TestErrWrapper(t *testing.T) {
	for _, tt := range tests {
		
		f := errWrapper(tt.h)
		
		response := httptest.NewRecorder()
		request := httptest.NewRequest(
			http.MethodGet,
			"http://www.imooc.com", nil)
		f(response, request)


		verfyResponse(response.Result(), tt.code, tt.message, t)

	}
	
}

// TestErrWrapperInServer 使用真正的WEB服务器测试      测试的是整个服务器
func TestErrWrapperInServer(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		server := httptest.NewServer(http.HandlerFunc(f))
		resp, _ := http.Get(server.URL)
		verfyResponse(resp, tt.code, tt.message, t)
	}
}

func verfyResponse(resp *http.Response, 
	exceptedCode int, exceptedMsg string,
	t *testing.T){
		b, _ := ioutil.ReadAll(resp.Body)
		body := strings.Trim(string(b), "\n")
		if resp.StatusCode != exceptedCode ||
			body != exceptedMsg {
			t.Errorf("expect (%d, %s); " + 
				"got (%d, %s)",
				exceptedCode, exceptedMsg,	
				resp.StatusCode, body)
		}

}
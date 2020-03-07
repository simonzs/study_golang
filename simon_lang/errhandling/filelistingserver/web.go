package main

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/", errWrapper(HandleFilelist))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}

type apphandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler apphandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {

		// panic
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic: %v", err)
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()

		err := handler(writer, request)

		if err != nil {
			log.Printf("Error occurred "+
				"request: %s",
				err.Error())

			// user Error
			if userError, ok := err.(userError); ok {
				http.Error(writer,
					userError.Message(),
					http.StatusBadRequest)
				return
			}

			// log.Warn("Error handling request: %s", err.Error())

			// system error
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

type userError interface {
	error
	Message() string
}

const prefix = "/list/"

// HandleFilelist ...
func HandleFilelist(writer http.ResponseWriter,
	request *http.Request) error {
	if strings.Index(request.URL.Path, prefix) != 0 {
		return errors.New("Path must start with " +
			"with " + prefix)
	}
	path := request.URL.Path[len("/list/"):] // /list/web.go
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	writer.Write(all)
	return nil
}

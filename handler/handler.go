package handler

import (
	"log"
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {

	return func(writer http.ResponseWriter, request *http.Request) {

		err := handler(writer, request)

		defer func() {
			r := recover()
			log.Printf("Panic: %v", r)
			http.Error(writer, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
		}()

		if err != nil {
			log.Printf("Error handleing request: %s", err.Error())

			if userError, ok := err.(userError); ok {
				http.Error(writer,
					userError.Message(),
					http.StatusBadRequest)

				return
			}

			code := http.StatusOK

			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			}
		}
	}
}

type userError interface {
	error
	Message() string
}

package middleware

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// LoggerMiddleware - http middleware for log request body
type LoggerMiddleware struct {
	Read    func(reader io.Reader) ([]byte, error)
	Handler http.Handler
}

func (middleware LoggerMiddleware) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	requestBody, err := middleware.Read(request.Body)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}
	request.Body = ioutil.NopCloser(bytes.NewReader(requestBody))
	defer log.Println(string(requestBody))

	middleware.Handler.ServeHTTP(responseWriter, request)
}

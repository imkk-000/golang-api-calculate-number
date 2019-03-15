package middleware

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

type LoggerMiddleware struct {
	Handler http.Handler
}

func (middleware LoggerMiddleware) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	requestBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}
	request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))
	defer log.Println(string(requestBody))

	middleware.Handler.ServeHTTP(responseWriter, request)
}

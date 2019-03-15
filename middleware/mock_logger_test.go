package middleware_test

import (
	"errors"
	"io"
	"net/http"
)

type mockAPI struct{}

func (mockAPI) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
}

func mockReadError(reader io.Reader) ([]byte, error) {
	return nil, errors.New("mock read error")
}

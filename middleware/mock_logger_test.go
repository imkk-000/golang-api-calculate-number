package middleware_test

import "net/http"

type mockAPI struct{}

func (mockAPI) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
}

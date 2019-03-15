package api

import "net/http"

func CalculateAPI(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Header().Add("Content-Type", "application/json")
	responseWriter.Write([]byte(`{"result": 3}`))
}

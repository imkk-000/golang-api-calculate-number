package middleware_test

import (
	"bytes"
	. "calculate-number/middleware"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLoggerMiddlewareReadSuccess(t *testing.T) {
	expectedLogRequestBody := `{"number1": 3, "number2": 3}`
	requestBody := `{"number1": 3, "number2": 3}`
	request := httptest.NewRequest(http.MethodPost, "/calculate", bytes.NewBufferString(requestBody))
	loggerMiddleware := LoggerMiddleware{
		Handler: mockAPI{},
		Read:    ioutil.ReadAll,
	}
	var buffer bytes.Buffer
	log.SetOutput(&buffer)

	loggerMiddleware.ServeHTTP(nil, request)

	if !strings.Contains(buffer.String(), expectedLogRequestBody) {
		t.Errorf("expected '%s' but it got '%s'", expectedLogRequestBody, buffer.String())
	}
}

func TestLoggerMiddlewareReadError(t *testing.T) {
	expectedStatusCode := http.StatusInternalServerError
	request := httptest.NewRequest(http.MethodPost, "/calculate", nil)
	responseRecorder := httptest.NewRecorder()
	loggerMiddleware := LoggerMiddleware{
		Handler: mockAPI{},
		Read:    mockReadError,
	}

	loggerMiddleware.ServeHTTP(responseRecorder, request)
	response := responseRecorder.Result()

	if expectedStatusCode != response.StatusCode {
		t.Errorf("expected '%d' but it got '%d'", expectedStatusCode, response.StatusCode)
	}
}

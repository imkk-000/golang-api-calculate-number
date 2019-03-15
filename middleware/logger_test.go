package middleware_test

import (
	"bytes"
	. "calculate-number/middleware"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLoggerMiddleware(t *testing.T) {
	expectedLogRequestBody := `{"number1": 3, "number2": 3}`
	requestBody := `{"number1": 3, "number2": 3}`
	request := httptest.NewRequest(http.MethodPost, "/calculate", bytes.NewBufferString(requestBody))
	loggerMiddleware := LoggerMiddleware{
		Handler: mockAPI{},
	}
	var buffer bytes.Buffer
	log.SetOutput(&buffer)

	loggerMiddleware.ServeHTTP(nil, request)

	if !strings.Contains(buffer.String(), expectedLogRequestBody) {
		t.Errorf("expected '%s' but it got '%s'", expectedLogRequestBody, buffer.String())
	}
}

package api_test

import (
	"bytes"
	. "calculate-number/api"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCalculateAPIRequestMethodPostAndJsonNumber1Is1Number2Is2ShouldBeResponseJsonResultIs3(t *testing.T) {
	expectedResultJson := `{"result":3}`
	requestBody := `{"number1": 1, "number2": 2}`
	request := httptest.NewRequest(http.MethodPost, "/calculate", bytes.NewBufferString(requestBody))
	responseRecorder := httptest.NewRecorder()

	CalculateAPI(responseRecorder, request)
	response := responseRecorder.Result()
	actualResponseJson, _ := ioutil.ReadAll(response.Body)

	if expectedResultJson != string(actualResponseJson) {
		t.Errorf("expect '%s' but it got '%s'", expectedResultJson, string(actualResponseJson))
	}
}

func TestCalculateAPIRequestMethodPostAndJsonNumber1IsMinus1Number1Is2ShouldBeResponseJsonResultIs0(t *testing.T) {
	expectedResultJson := `{"result":0}`
	requestBody := `{"number1": -1, "number2": 1}`
	request := httptest.NewRequest(http.MethodPost, "/calculate", bytes.NewBufferString(requestBody))
	responseRecorder := httptest.NewRecorder()

	CalculateAPI(responseRecorder, request)
	response := responseRecorder.Result()
	actualResponseJson, _ := ioutil.ReadAll(response.Body)

	if expectedResultJson != string(actualResponseJson) {
		t.Errorf("expect '%s' but it got '%s'", expectedResultJson, string(actualResponseJson))
	}
}

func TestCalculateAPIRequestMethodPostAndJsonNumberAIs1ShouldBeResponseStatusCodeBadRequest(t *testing.T) {
	expectedStatusCode := http.StatusBadRequest
	requestBody := `{"number1": A}`
	request := httptest.NewRequest(http.MethodPost, "/calculate", bytes.NewBufferString(requestBody))
	responseRecorder := httptest.NewRecorder()

	CalculateAPI(responseRecorder, request)
	response := responseRecorder.Result()

	if expectedStatusCode != response.StatusCode {
		t.Errorf("expect '%d' but it got '%d'", expectedStatusCode, response.StatusCode)
	}
}

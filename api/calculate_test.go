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
	expectedResultJson := `{"result": 3}`
	requestBody := []byte(`{"number1": 1, "number2": 2}`)
	request := httptest.NewRequest(http.MethodPost, "/calculate", bytes.NewBuffer(requestBody))
	responseRecorder := httptest.NewRecorder()

	CalculateAPI(responseRecorder, request)
	response := responseRecorder.Result()
	actualResponseJson, _ := ioutil.ReadAll(response.Body)

	if expectedResultJson != string(actualResponseJson) {
		t.Errorf("expect '%s' but it got '%s'", expectedResultJson, string(actualResponseJson))
	}
}

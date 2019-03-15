package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type requestBodyStruct struct {
	Number1 int `json:"number1"`
	Number2 int `json:"number2"`
}

type responseBodyStruct struct {
	Result int `json:"result"`
}

func CalculateAPI(responseWriter http.ResponseWriter, request *http.Request) {
	requestBody, _ := ioutil.ReadAll(request.Body)
	body := requestBodyStruct{}
	json.Unmarshal(requestBody, &body)
	response := responseBodyStruct{
		Result: body.Number1 + body.Number2,
	}
	responseBodyJSON, _ := json.Marshal(response)

	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Header().Add("Content-Type", "application/json")
	responseWriter.Write(responseBodyJSON)
}

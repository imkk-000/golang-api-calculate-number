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
	requestBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}
	body := requestBodyStruct{}
	err = json.Unmarshal(requestBody, &body)
	if err != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
		return
	}
	response := responseBodyStruct{
		Result: body.Number1 + body.Number2,
	}
	responseBodyJSON, err := json.Marshal(response)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Header().Add("Content-Type", "application/json")
	responseWriter.Write(responseBodyJSON)
}

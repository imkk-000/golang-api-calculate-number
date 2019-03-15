package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type requestBodyStruct struct {
	Number1 int `json:"number1"`
	Number2 int `json:"number2"`
}

func CalculateAPI(responseWriter http.ResponseWriter, request *http.Request) {
	requestBody, _ := ioutil.ReadAll(request.Body)
	body := requestBodyStruct{}
	json.Unmarshal(requestBody, &body)
	responseBodyJson := fmt.Sprintf(`{"result": %d}`, body.Number1+body.Number2)

	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Header().Add("Content-Type", "application/json")
	responseWriter.Write([]byte(responseBodyJson))
}

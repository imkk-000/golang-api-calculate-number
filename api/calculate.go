package api

import (
	"io"
	"net/http"
)

type CalculateAPI struct {
	Read      func(reader io.Reader) ([]byte, error)
	Unmarshal func(data []byte, v interface{}) error
	Marshal   func(v interface{}) ([]byte, error)
}

type requestBodyStruct struct {
	Number1 int `json:"number1"`
	Number2 int `json:"number2"`
}

type responseBodyStruct struct {
	Result int `json:"result"`
}

func (api CalculateAPI) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		responseWriter.WriteHeader(http.StatusNotFound)
		return
	}
	requestBody, err := api.Read(request.Body)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}
	body := requestBodyStruct{}
	err = api.Unmarshal(requestBody, &body)
	if err != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
		return
	}
	response := responseBodyStruct{
		Result: body.Number1 + body.Number2,
	}
	responseBodyJSON, err := api.Marshal(response)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Header().Add("Content-Type", "application/json")
	responseWriter.Write(responseBodyJSON)
}

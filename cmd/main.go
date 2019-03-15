package main

import (
	"calculate-number/api"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func main() {
	myAPI := api.API{
		Read:      ioutil.ReadAll,
		Marshal:   json.Marshal,
		Unmarshal: json.Unmarshal,
	}
	http.HandleFunc("/calculate", myAPI.CalculateAPI)
	http.ListenAndServe(":8888", nil)
}

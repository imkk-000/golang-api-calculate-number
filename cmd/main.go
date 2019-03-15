package main

import (
	"calculate-number/api"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	var port uint
	flag.UintVar(&port, "port", 8888, "Set port for api listening")

	myAPI := api.API{
		Read:      ioutil.ReadAll,
		Marshal:   json.Marshal,
		Unmarshal: json.Unmarshal,
	}
	http.HandleFunc("/calculate", myAPI.CalculateAPI)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	log.Fatal(err)
}

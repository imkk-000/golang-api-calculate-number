package main

import (
	"calculate-number/api"
	"calculate-number/middleware"
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

	calculateAPI := api.CalculateAPI{
		Read:      ioutil.ReadAll,
		Marshal:   json.Marshal,
		Unmarshal: json.Unmarshal,
	}
	loggerMiddleware := middleware.LoggerMiddleware{
		Read:    ioutil.ReadAll,
		Handler: calculateAPI,
	}
	http.Handle("/calculate", loggerMiddleware)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	log.Fatal(err)
}

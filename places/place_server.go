package main

import (
	"encoding/json"
	_ "fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var API_KEY string

func initapi() error {
	key, err := ioutil.ReadFile("keys/plats.key")
	checkError(err)
	API_KEY = string(key)
	return err
}

func main() {
	//read api key here
	checkError(initapi())
	//set up routes
	http.HandleFunc("/place", placeHandler)
	//start the api server
	log.Println("Server started and listens at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func placeHandler(resWritter http.ResponseWriter, req *http.Request) {
	log.Println("Incoming request to /place via", req.URL)
	parameters := req.URL.Query()

	if len(parameters) == 0 || parameters.Get("search") == "" {
		log.Println("The wrong parameters was passed to the api. Parameters:", parameters)
		resWritter.Header().Set("Content-Type", "application/json; charset=UTF-8")
		resWritter.WriteHeader(400) // unprocessable entity
		payl := ErrorPayload{"wrong parameters. should be place/?search="}
		json.NewEncoder(resWritter).Encode(payl)
		return
	}
	searchString := parameters.Get("search")
	SL_URL := "http://api.sl.se/api2/typeahead.json?key=" + API_KEY + "&searchstring=" + searchString
	log.Println("SL:", SL_URL)

	resp, errGet := http.Get(SL_URL) // call to SL-API.
	checkError(errGet)
	defer resp.Body.Close()

	body, errRead := ioutil.ReadAll(
		io.LimitReader(resp.Body, 1048576))
	checkError(errRead)

	var answer Answer
	if err := json.Unmarshal(body, &answer); err != nil && answer.StatusCode == 0 {
		resWritter.Header().Set("Content-Type", "application/json; charset=UTF-8")
		resWritter.WriteHeader(422) // unprocessable entity
		payl := ErrorPayload{"Something wrong with request to SL. :" + string(body)}
		json.NewEncoder(resWritter).Encode(payl)
		return //panic(err)
	}

	resWritter.Header().Set("Content-Type", "application/json; charset=UTF-8")
	resWritter.WriteHeader(http.StatusOK)

	payload := Payload{ConvertSitesToPlaces(answer.ResponseData)}

	json.NewEncoder(resWritter).Encode(payload)
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
		panic(err)
	}
}

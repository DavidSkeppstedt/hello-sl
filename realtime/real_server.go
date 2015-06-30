package main

import (
	"encoding/json"
	_ "fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var API_KEY string

func initapi() error {
	key, err := ioutil.ReadFile("keys/realtid.key")
	checkError(err)
	API_KEY = string(key)
	return err
}

func main() {
	//read api key here
	checkError(initapi())
	//set up routes
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/client/", http.StripPrefix("/client/", fs))
	http.HandleFunc("/realtime", realtimeHandler)
	//start the api server
	log.Println("Server started, listens at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func realtimeHandler(resWritter http.ResponseWriter, req *http.Request) {
	log.Println("Incoming request to /realtime via", req.URL)
	resWritter.Header().Set("Content-Type", "application/json; charset=UTF-8")

	parameters := req.URL.Query()
	if len(parameters) == 0 || parameters.Get("site") == "" {
		log.Println("The wrong parameters was passed to the api. Parameters:", parameters)
		resWritter.WriteHeader(http.StatusBadRequest) // unprocessable entity
		payl := ErrorPayload{"wrong parameters. should be realtime/?site="}
		json.NewEncoder(resWritter).Encode(payl)
		return
	}
	searchString := parameters.Get("site")
	searchString = strings.Replace(searchString, " ", "", -1) //hack to fix sl broken api

	SL_URL := "http://api.sl.se/api2/realtimedepartures.json?key=" + API_KEY + "&siteid=" + searchString + "&timewindow=60"
	//log.Println("SL:", SL_URL)

	resp, errGet := http.Get(SL_URL) // call to SL-API.
	defer resp.Body.Close()
	checkError(errGet)

	body, errRead := ioutil.ReadAll(
		io.LimitReader(resp.Body, 1048576))
	checkError(errRead)
	var realtime Realtime
	if err := json.Unmarshal(body, &realtime); err != nil && realtime.StatusCode == 0 {
		resWritter.WriteHeader(422) // unprocessable entity
		payl := ErrorPayload{"Something wrong with request to SL. :" + string(body)}
		json.NewEncoder(resWritter).Encode(payl)
		log.Println("Error unmarshal:", err)
		return //panic(err)
	}

	//Everyting is okay.
	resWritter.WriteHeader(http.StatusOK)

	json.NewEncoder(resWritter).Encode(realtime)
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
		panic(err)
	}
}

package main

type Realtime struct {
	StatusCode      int
	Message         string
	ExecutationTime int
	ResponseData    Departure
}

type Departure struct {
	Buses []Ride
}

type Ride struct {
	LineNumer   string
	Destination string
	DisplayTime string
}

type ErrorPayload struct {
	ErrorMessage string
}

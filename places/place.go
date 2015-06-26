package main

type Answer struct {
	StatusCode    int
	Message       string
	ExecutionTime int64
	ResponseData  Sites
}
type Sites []Site

type Site struct {
	Name   string
	SiteId string
	Type   string
	X      string
	Y      string
}

func ConvertSitesToPlaces(sites Sites) Places {
	places := make(Places, 0)
	for _, v := range sites {
		place := ConvertSiteToString(v)
		places = append(places, place)
	}
	return places
}

func ConvertSiteToString(site Site) Place {
	p := Place{site.Name, site.SiteId}
	return p
}

type Payload struct {
	Places Places `json:"places"`
}
type Places []Place

type Place struct {
	Name   string
	SiteId string
}

type ErrorPayload struct {
	ErrorMessage string
}

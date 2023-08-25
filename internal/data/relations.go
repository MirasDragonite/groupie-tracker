package data

import (
	"encoding/json"
	"log"
	"net/http"
)

type DateLocation struct {
	Date string `json:"date"`
}

type DatesLocations map[string][]string

type IndexItem struct {
	ID             int            `json:"id"`
	DatesLocations DatesLocations `json:"datesLocations"`
}

type Index struct {
	Items []IndexItem `json:"index"`
}

func GetLocationsAndDates() Index {
	jsonData, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		log.Fatal(err)
	}
	var index Index
	err = json.NewDecoder(jsonData.Body).Decode(&index)

	if err != nil {
		log.Fatal(err)
	}

	return index
}

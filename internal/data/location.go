package data

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type Index1 struct {
	Index []Location `json:"index"`
}

type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

func GetLocations() Index1 {
	data, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		log.Fatal(err)
	}
	var index Index1

	err = json.NewDecoder(data.Body).Decode(&index)

	if err != nil {
		log.Fatal(err)
	}
	return index
}

func GetLocationById(id int) ([]string, error) {
	err := errors.New("Wrong id")
	locations := GetLocations()
	for _, ch := range locations.Index {
		if ch.ID == id {
			return ch.Locations, nil
		}
	}
	return nil, err
}

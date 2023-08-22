package data

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const mapboxAPIURL = "https://api.mapbox.com/geocoding/v5/mapbox.places/"

type Feature struct {
	ID string `json:"id"`

	Center []float64 `json:"center"`
}

type FeatureCollection struct {
	Features []Feature `json:"features"`
}

func GetCoordinates(location string) []float64 {
	// fmt.Println(location)
	apiKey := "pk.eyJ1IjoibW90b2JlIiwiYSI6ImNsbG0xdnQzYTJqZG8zZ21neTJuN28wemoifQ.U2Zp2USylpY-WQyXi8TYfw"
	result := make([]float64, 2)
	// Create the URL for the API request
	url := fmt.Sprintf("%s%s.json?access_token=%s", mapboxAPIURL, location, apiKey)
	// fmt.Println(url)
	// Make the HTTP GET request
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making request:", err)
		return result
	}
	var coordinate FeatureCollection
	err = json.NewDecoder(response.Body).Decode(&coordinate)

	if err != nil {
		log.Fatal(err)
	}

	result[0] = float64(coordinate.Features[1].Center[0])
	result[1] = float64(coordinate.Features[1].Center[1])

	return result
}

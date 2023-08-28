package data

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

const mapboxAPIURL = "https://maps.googleapis.com/maps/api/geocode/"

// https://maps.googleapis.com/maps/api/geocode/json?address=1600+Amphitheatre+Parkway,+Mountain+View,+CA&key=YOUR_API_KEY
type Locations struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Geometry struct {
	Locations Locations `json:"location"`
}

type Result struct {
	Geometry Geometry `json:"geometry"`
}

type GeocodeResponse struct {
	Results []Result `json:"results"`
}

func GetCoordinates(location, apiKey string) []float64 {
	result := make([]float64, 2)

	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?address=%s&key=%s", location, apiKey)
	fmt.Println(url)
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making request:", err)
		return result
	}
	defer response.Body.Close()

	var geocodeResponse GeocodeResponse
	err = json.NewDecoder(response.Body).Decode(&geocodeResponse)
	if err != nil {
		log.Fatal(err)
	}

	result[1] = geocodeResponse.Results[0].Geometry.Locations.Lat
	result[0] = geocodeResponse.Results[0].Geometry.Locations.Lng

	return result
}

func GetCoordinatesBatch(locations []string) [][]float64 {
	apiKey := "AIzaSyBXH2PYMwKrL18rjTE-O5OdtEgUZywoIgo"
	coordinateCh := make(chan []float64, len(locations))
	var wg sync.WaitGroup

	for _, location := range locations {
		wg.Add(1)
		go func(loc string) {
			defer wg.Done()
			coordinates := GetCoordinates(loc, apiKey)
			coordinateCh <- coordinates
		}(location)
	}

	go func() {
		wg.Wait()
		close(coordinateCh)
	}()

	var coor [][]float64
	for coords := range coordinateCh {
		coor = append(coor, coords)
	}

	return coor
}

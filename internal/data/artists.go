package data

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

func GetArtists() []Artist {
	data, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Fatal(err)
	}

	var artists []Artist

	err = json.NewDecoder(data.Body).Decode(&artists)

	if err != nil {
		log.Fatal(err)
	}
	return artists
}

func GetData(id int) (Artist, IndexItem, error) {
	var empty1 Artist
	var empty2 IndexItem
	err := errors.New("Wrong number")
	artists := GetArtists()
	locADate := getLocationsAndDates()

	for _, ch := range artists {
		for _, el := range locADate.Items {
			if ch.Id == id && el.ID == id {
				return ch, el, nil
			}
		}
	}

	return empty1, empty2, err
}

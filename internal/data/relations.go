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

func getLocationsAndDates() Index {
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

// func main() {
// 	jsonData, _ := os.ReadFile("relation.json")

// 	var indexData Index
// 	if err := json.Unmarshal(jsonData, &indexData); err != nil {
// 		fmt.Println("Error unmarshaling JSON:", err)
// 		return
// 	}

// 	// Now you can work with the unmarshaled data
// 	for _, item := range indexData.Items {
// 		fmt.Printf("ID: %d\n", item.ID)
// 		for location, dates := range item.DatesLocations {
// 			fmt.Printf("Location: %s\n", location)
// 			for _, date := range dates {
// 				fmt.Printf("Date: %s\n", date)
// 			}
// 		}
// 		fmt.Println()
// 	}
// }

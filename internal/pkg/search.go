package pkg

import (
	"fmt"
	"strconv"
	"strings"

	"groupie-tracker/internal/data"
)

func Search(s string) []data.Artist {
	artist := data.GetArtists()
	locations := data.GetLocations().Index
	var result []data.Artist
	if strings.TrimSpace(s) == "" {
		return artist
	}
	for i, ch := range artist {
		if Contains(strings.ToLower(strings.TrimSpace(ch.Name)), strings.ToLower(strings.TrimSpace(s))) {
			if !ArtistContains(result, ch) {
				result = append(result, ch)
				fmt.Println("NAME")
			}
		} else if Contains(strings.ToLower(strings.TrimSpace(strconv.Itoa(ch.CreationDate))), strings.ToLower(strings.TrimSpace(s))) {
			if !ArtistContains(result, ch) {
				result = append(result, ch)
				fmt.Println("CreationDate")
			}
		}
		if Contains(strings.ToLower(strings.TrimSpace(ch.FirstAlbum)), strings.ToLower(strings.TrimSpace(s))) {
			if !ArtistContains(result, ch) {
				result = append(result, ch)
				fmt.Println("FirstDate")
			}
		}
		for _, el := range ch.Members {
			if Contains(strings.ToLower(el), strings.ToLower(s)) {
				if !ArtistContains(result, ch) {
					result = append(result, ch)
					fmt.Println("members")
				}
			}
		}

		for _, loc := range locations[i].Locations {
			if Contains(strings.ToLower(strings.TrimSpace(loc)), strings.ToLower(strings.TrimSpace(s))) {
				if !ArtistContains(result, artist[i]) {
					fmt.Println("location")
					result = append(result, artist[i])
				}
			}
		}

	}

	return result
}

func Contains(s, str string) bool {
	if strings.Contains(s, str) {
		return true
	} else {
		return false
	}
}

func ArtistContains(artist []data.Artist, el data.Artist) bool {
	for _, ch := range artist {
		if ch.Id == el.Id {
			return true
		}
	}
	return false
}

func IsTextValid(text string) bool {
	for _, ch := range text {
		if (int(ch) < 32 || int(ch) > 126) && int(ch) != 10 && int(ch) != 9 {
			return false
		}
	}
	return true
}

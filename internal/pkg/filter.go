package pkg

import (
	"groupie-tracker/internal/data"
)

func Filter(min, max int) []data.Artist {
	artists := data.GetArtists()
	var result []data.Artist
	if min >= max {
		return result
	}
	for _, ch := range artists {
		if ch.CreationDate > min && ch.CreationDate < max {
			result = append(result, ch)
		}
	}

	return result
}

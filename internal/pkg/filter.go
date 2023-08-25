package pkg

import (
	"strconv"

	"groupie-tracker/internal/data"
)

func Filter(minCD, maxCD, minFA, maxFA int, loc string, members []int) []data.Artist {
	artists := data.GetArtists()
	location := data.GetLocations()
	var result []data.Artist
	if minCD >= maxCD || minFA >= maxFA {
		return result
	}
	for i, ch := range artists {
		counter := 0
		for range ch.Members {
			counter++
		}
		located := false
		member := false
		if loc != "none" {
			for j, ch := range location.Index {
				if i == j {
					for _, el := range ch.Locations {
						if el == loc {
							located = true
						}
					}
				}
			}
		} else {
			located = true
		}

		for _, ch := range members {
			if ch == counter {
				member = true
			}
		}
		artFA, _ := strconv.Atoi(ch.FirstAlbum[6:])
		if (ch.CreationDate >= minCD && ch.CreationDate <= maxCD) && (artFA >= minFA && artFA <= maxFA) && located && member {
			result = append(result, ch)
		}
	}

	return result
}

func AddToSlice(slic []int, n1, n2, n3, n4, n5, n6, n7, n8 int) []int {
	if n1 == 0 && n2 == 0 && n3 == 0 && n4 == 0 && n5 == 0 && n6 == 0 && n7 == 0 && n8 == 0 {
		slic = append(slic, 1)
		slic = append(slic, 2)
		slic = append(slic, 3)
		slic = append(slic, 4)
		slic = append(slic, 5)
		slic = append(slic, 6)
		slic = append(slic, 7)
		slic = append(slic, 8)
		return slic
	}

	slic = append(slic, n1)
	slic = append(slic, n2)
	slic = append(slic, n3)
	slic = append(slic, n4)
	slic = append(slic, n5)
	slic = append(slic, n6)
	slic = append(slic, n7)
	slic = append(slic, n8)
	return slic
}

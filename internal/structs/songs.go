package structs

import "fmt"

type Song struct {
	Title             string
	Artist            string
	ReleaseYear       uint
	DurationInSeconds uint
}

func (s *Song) DisplaySong() {
	fmt.Println("Title:", s.Title)
	fmt.Println("Artist:", s.Artist)
	fmt.Println("Release Year:", s.ReleaseYear)
	fmt.Println("Duration in Seconds:", s.DurationInSeconds)
}

func songs(show bool) {
	if show {
		song := Song{
			Title:             "Dream On",
			Artist:            "Aerosmith",
			ReleaseYear:       1973,
			DurationInSeconds: 268,
		}

		song.DisplaySong()

		// in Go I can continue using song in Rust I cannot
		fmt.Println("Song Title:", song.Title)
	}
}

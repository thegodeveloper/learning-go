package structs

import "fmt"

type Song struct {
	Title             string
	Artist            string
	ReleaseYear       uint
	DurationInSeconds uint
}

// this is the way to define methods for a struct
func (s *Song) DisplaySong() {
	fmt.Println("Title:", s.Title)
	fmt.Println("Artist:", s.Artist)
	fmt.Println("Release Year:", s.ReleaseYear)
	fmt.Println("Duration in Seconds:", s.DurationInSeconds)
}

func (s *Song) DoubleLenght() {
	s.DurationInSeconds *= 2
}

func Songs(show bool) {
	if show {
		song := Song{
			Title:             "Dream On",
			Artist:            "Aerosmith",
			ReleaseYear:       1973,
			DurationInSeconds: 268,
		}

		// I am calling the method in the object and the instance is passed as reference
		song.DisplaySong()

		// in Go I can continue using song in Rust I cannot
		fmt.Println("Song Title:", song.Title)

		// exactly the same here, the instance is passed as reference to the method
		song.DoubleLenght()
		fmt.Println("Duration in Seconds:", song.DurationInSeconds)
	}
}

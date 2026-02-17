package arrays

import "fmt"

func ArrayStrings(show bool) {
	if show {
		var germanNumbers = [10]string{"eins", "zwei", "drei", "vier", "funf", "sechs", "sieben", "acth", "neun", "zehn"}

		for num, value := range germanNumbers {
			fmt.Printf("number %d in German is %s\n ", num+1, value)
		}
	}
}

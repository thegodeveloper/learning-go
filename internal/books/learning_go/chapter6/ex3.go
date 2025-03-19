package chapter6

func ex3(show bool) {
	if show {
		var people []Person

		for i := 0; i < 10_000_000; i++ {
			people = append(people, MakePerson("Fred", "Williamson", 25))
		}
	}
}

package strain

// Implement the "Keep" and "Discard" function in this file.

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1

func Keep[T any](c []T, predicate func(T) bool) []T {
	var res []T
	for _, v := range c {
		if predicate(v) {
			res = append(res, v)
		}
	} 
	return res
}

func Discard[T any](c []T, predicate func(T) bool) []T {
	var res []T
	for _, v := range c {
		if !predicate(v) {
			res = append(res, v)
		}
	}
	return res
}

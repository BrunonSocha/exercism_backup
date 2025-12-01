package stringset

import "slices"

// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

// Define the Set type here.

type Set struct {
	Values []string
}

func New() Set {
	return Set{}
}

func NewFromSlice(l []string) Set {
	var newSli []string
	for _, l := range l {
		if !slices.Contains(newSli, l) {
			newSli = append(newSli, l)
		}
	}
	return Set{newSli}
}

func (s Set) String() string {
	res := "{"
	for _, l := range s.Values {
		res += "\"" + l + "\", "
	}
	if len(res) > 2 {
		return res[:len(res)-2] + "}"
	}
	return res + "}"
}

func (s Set) IsEmpty() bool {
	if len(s.Values) == 0 {
		return true
	}
	return false
}

func (s Set) Has(elem string) bool {
	if slices.Contains(s.Values, elem){
		return true
	}
	return false
}

func (s *Set) Add(elem string) {
	if !slices.Contains(s.Values, elem) {
		s.Values = append(s.Values, elem)
	}
}

func Subset(s1, s2 Set) bool {
	for _, l := range s1.Values {
		if !slices.Contains(s2.Values, l) {
			return false
		}
	}
	return true
}

func Disjoint(s1, s2 Set) bool {
	for _, l := range s2.Values {
		if slices.Contains(s1.Values, l) {
			return false
		}
	}
	return true
}

func Equal(s1, s2 Set) bool {
	if len(s1.Values) != len(s2.Values) {
		return false
	}
	for _, l := range s1.Values {
		if !slices.Contains(s2.Values, l) {
			return false
		}
	}
	return true
}

func Intersection(s1, s2 Set) Set {
	var setSli []string
	for _, l := range s1.Values {
		if slices.Contains(s2.Values, l) {
			setSli = append(setSli, l)
		}
	}
	return Set{setSli}
}

func Difference(s1, s2 Set) Set {
	var setSli []string
	for _, l := range s1.Values {
		if !slices.Contains(s2.Values, l) {
			setSli = append(setSli, l)
		}
	}
	return Set{setSli}
}

func Union(s1, s2 Set) Set {
	res := s1
	for _, l := range s2.Values {
		res.Add(l)
	}
	return res
}

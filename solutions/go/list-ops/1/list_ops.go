package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	res := initial
	for _, v := range s {
		res = fn(res, v)
	}
	return res
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	res := initial
	for i := s.Length()-1; i > -1; i-- {
		res = fn(s[i], res)
	}
	return res
}

func (s IntList) Filter(fn func(int) bool) IntList {
	if s.Length() == 0 {
		return s
	}
	var res IntList
	for _, v := range s {
		if fn(v) {
			res = res.Append(IntList{v})
		}
	}
	return res
}

func (s IntList) Length() int {
	res := 0
	for i, _ := range s {
		res = i+1
	}
	return res
	
}

func (s IntList) Map(fn func(int) int) IntList {
	res := make(IntList, 0, len(s))
	for _, v := range s {
		res = res.Append(IntList{fn(v)})
	}
	return res
}

func (s IntList) Reverse() IntList {
	res := make(IntList, s.Length())
	j := 0
	for i := s.Length()-1; i > -1; i-- {
		res[j] = s[i]
		j++
	}
	return res
}

func (s IntList) Append(lst IntList) IntList {
	res := make(IntList, s.Length()+lst.Length())
	for i, _ := range s {
		res[i] = s[i]
	}
	j := 0
	for i := s.Length(); i < (s.Length() + lst.Length()); i++ {
		res[i] = lst[j]
		j++
	}
	return res
}

func (s IntList) Concat(lists []IntList) IntList {
	res := s
	for _, v := range lists {
		res = res.Append(v)	
	}
	return res
}

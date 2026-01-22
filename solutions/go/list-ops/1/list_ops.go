package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	for _, n := range s {
		initial = fn(initial, n)
	}

	return initial
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	for i := len(s) - 1; i >= 0; i-- {
		initial = fn(s[i], initial)
	}

	return initial
}

func (s IntList) Filter(fn func(int) bool) IntList {
	out := IntList{}

	for _, n := range s {
		if fn(n) {
			out = append(out, n)
		}
	}

	return out
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
	out := IntList{}

	for _, n := range s {
		out = append(out, fn(n))
	}

	return out
}

func (s IntList) Reverse() IntList {
	out := IntList{}

	for i := range s {
		out = append(out, s[len(s)-i-1])
	}

	return out
}

func (s IntList) Append(lst IntList) IntList {
	out := IntList{}

	out = append(out, s...)
	out = append(out, lst...)

	return out
}

func (s IntList) Concat(lists []IntList) IntList {
	out := IntList{}

	out = append(out, s...)

	for _, l := range lists {
		out = append(out, l...)
	}

	return out
}

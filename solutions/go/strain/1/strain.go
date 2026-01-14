package strain

func Keep[T any](arr []T, f func(T) bool) []T {
	var out []T

	for _, v := range arr {
		if f(v) {
			out = append(out, v)
		}
	}

	return out

}

func Discard[T any](arr []T, f func(T) bool) []T {
	var out []T

	for _, v := range arr {
		if !f(v) {
			out = append(out, v)
		}
	}

	return out
}

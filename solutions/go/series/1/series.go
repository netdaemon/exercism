package series

func All(n int, s string) []string {
	out := []string{}

	for i := 0; i <= len(s)-n; i++ {
		out = append(out, s[i:n+i])
	}

	return out
}

func UnsafeFirst(n int, s string) string {
	return s[:n]
}

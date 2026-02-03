package flatten

func Flatten(nested any) []any {
	flattened := []any{}

	switch i := nested.(type) {
	case []any:
		for _, val := range i {
			flattened = append(flattened, Flatten(val)...)
		}
	case any:
		flattened = append(flattened, i)
	}

	return flattened
}

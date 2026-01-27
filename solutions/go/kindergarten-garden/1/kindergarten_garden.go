package kindergarten

import (
	"errors"
	"slices"
	"strings"
)

type Garden map[string][]string

var plants = map[byte]string{
	'G': "grass",
	'C': "clover",
	'R': "radishes",
	'V': "violets",
}

// The diagram argument starts each row with a '\n'.  This allows Go's
// raw string literals to present diagrams in source code nicely as two
// rows flush left, for example,
//
//     diagram := `
//     VVCCGG
//     VVCCGG`

func NewGarden(diagram string, children []string) (*Garden, error) {
	if len(children) == 0 {
		return nil, errors.New("No children")
	}

	rows := strings.Split(diagram, "\n")

	if len(rows) < 3 || len(rows[1]) != len(rows[2]) || len(rows[1])%2 != 0 {
		return nil, errors.New("Invalid diagram")
	}

	garden := Garden{}

	slices.Sort(children)

	for i, child := range children {
		_, recordExists := garden[child]

		if recordExists {
			return nil, errors.New("Duplicate child")
		}

		plant1, exists1 := plants[rows[1][2*i]]
		plant2, exists2 := plants[rows[1][2*i+1]]
		plant3, exists3 := plants[rows[2][2*i]]
		plant4, exists4 := plants[rows[2][2*i+1]]

		if !exists1 || !exists2 || !exists3 || !exists4 {
			return nil, errors.New("Invalid cup code")
		}

		garden[child] = append(garden[child], plant1, plant2, plant3, plant4)
	}

	return &garden, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	plants, found := (*g)[child]

	return plants, found
}

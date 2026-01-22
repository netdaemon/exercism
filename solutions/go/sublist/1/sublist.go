package sublist

import "reflect"

// Relation type is defined in relations.go file.

func Sublist(l1, l2 []int) Relation {
	if len(l1) == len(l2) && reflect.DeepEqual(l1, l2) {
		return RelationEqual
	}

	if len(l1) > len(l2) {
		for i := 0; i <= len(l1)-len(l2); i++ {
			slice := l1[i : len(l2)+i]

			if reflect.DeepEqual(slice, l2) {
				return RelationSuperlist
			}
		}
	}

	if len(l2) > len(l1) {
		for i := 0; i <= len(l2)-len(l1); i++ {
			slice := l2[i : len(l1)+i]

			if reflect.DeepEqual(slice, l1) {
				return RelationSublist
			}
		}
	}

	return RelationUnequal
}

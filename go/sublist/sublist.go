package sublist

import "reflect"

// Relation is the comprasion between lists
type Relation string

// Possible relations
const (
	RelationEqual     Relation = "equal"
	RelationSublist   Relation = "sublist"
	RelationSuperlist Relation = "superlist"
	RelationUnequal   Relation = "unequal"
)

// Sublist checks difference of two lists and
// returns equal, sublist, superlist or unequal according
// to their relation to each other.
func Sublist(listOne, listTwo []int) Relation {
	if reflect.DeepEqual(listOne, listTwo) {
		return RelationEqual
	}
	if len(listOne) > len(listTwo) && checkSubList(listOne, listTwo) {
		return RelationSuperlist
	}
	if len(listOne) < len(listTwo) && checkSubList(listTwo, listOne) {
		return RelationSublist
	}
	return RelationUnequal
}

func checkSubList(listOne, listTwo []int) bool {
	if len(listTwo) == 0 {
		return true
	}
	for i, v := range listOne {
		if v == listTwo[0] {
			if reflect.DeepEqual(listOne[i:(i+len(listTwo))], listTwo) {
				return true
			}
		}
	}
	return false
}

package sublist

import "slices"

// Relation type is defined in relations.go file.

func Sublist(l1, l2 []int) Relation {
	if len(l1) == 0 {
		if len(l2) == 0 {
			return RelationEqual
		}
		return RelationSublist
	} else if len(l2) == 0 {
		return RelationSuperlist
	}
	if slices.Compare(l1, l2) == 0 {
		return RelationEqual
	}
	if len(l1) > len(l2) {
		found := false
		for i, v := range l1 {
			if v == l2[0] {
				found = true
				for j := 0; j < len(l2); j++ {
					if l2[j] != l1[i+j] {
						found = false
						break
					}
				}
				if found {
					return RelationSuperlist
				}
			}
			
		}
		return RelationUnequal
	}
	if len(l1) < len(l2) {
		found := false
		for i, v := range l2 {
			if v == l1[0] {
				found = true
				for j := 0; j < len(l1); j++ {
					if l1[j] != l2[i+j] {
						found = false
						break
					}
				}
				if found {
					return RelationSublist
				}
				
			}
		}
		return RelationUnequal
	}	

	return RelationUnequal
}

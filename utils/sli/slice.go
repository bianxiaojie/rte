package sli

import "slices"

func EqualIgnoreOrder[S ~[]E, E comparable](s1, s2 S) bool {
	if len(s1) != len(s2) {
		return false
	}
	r := make([]bool, len(s2))
	for _, e1 := range s1 {
		matched := false
		for j, e2 := range s2 {
			if r[j] {
				continue
			}
			if e1 == e2 {
				matched = true
				r[j] = true
			}
		}
		if !matched {
			return false
		}
	}
	return true
}

func EqualIgnoreOrderFunc[S1 ~[]E1, S2 ~[]E2, E1, E2 any](s1 S1, s2 S2, eq func(E1, E2) bool) bool {
	if len(s1) != len(s2) {
		return false
	}
	r := make([]bool, len(s2))
	for _, e1 := range s1 {
		matched := false
		for j, e2 := range s2 {
			if r[j] {
				continue
			}
			if eq(e1, e2) {
				matched = true
				r[j] = true
			}
		}
		if !matched {
			return false
		}
	}
	return true
}

func Delete[S ~[]E, E comparable](s S, v E) S {
	return slices.DeleteFunc(s, func(c E) bool {
		return c == v
	})
}

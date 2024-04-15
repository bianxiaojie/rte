package str

import "regexp"

func MatchAndFindGroups(pattern string, input string) ([]string, bool) {
	r, err := regexp.Compile(pattern)
	if err != nil {
		return nil, false
	}

	if !r.MatchString(input) {
		return nil, false
	}

	return r.FindStringSubmatch(input), true
}

package main

import "regexp"

func checkIntegerFormat(s string) bool {
	matched, err := regexp.MatchString("\\b[0-9]+\\b", s)
	if err != nil {
		return false
	}

	return matched
}

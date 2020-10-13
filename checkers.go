package main

import "regexp"

func checkIntegerFormat(s string) bool {
	m, err := regexp.MatchString("\\b[0-9]+\\b", s)
	if err != nil {
		return false
	}
	return m
}

func checkMultivectorFormat(s string) bool {
	m, err := regexp.MatchString("\\b[0-9]+e0\\+[0-9]+e1\\+[0-9]+e2\\+[0-9]+e3\\+[0-9]+e12\\+[0-9]+e13\\+[0-9]+e23\\+[0-9]+e123\\b", s)
	if err != nil {
		return false
	}
	return m
}

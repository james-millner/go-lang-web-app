package web

import "strings"

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func split(str string, sep string) []string {
	return strings.Split(str, sep)
}
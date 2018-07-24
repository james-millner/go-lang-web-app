package web

import (
	"math/rand"
	"strings"
	"time"
)

//ArrayContains method for determining if a String exists within a string array.
func ArrayContains(arr []string, str string) bool {
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

//RandomSleep method
func RandomSleep(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

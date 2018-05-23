package web

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSplit(t *testing.T) {
	str := "abc,123,baby,you,and,me,girl"

	array := split(str, ",")

	assert.Equal(t, 7, len(array))
}

func TestContains(t *testing.T) {
	strArray := []string{
		"https://thenextweb.com/contributors/2018/05/18/15-online-trends-to-watch-for-in-2018-and-beyond/",
		"https://thenextweb.com/section/insights/",
		"https://thenextweb.com/section/insights.pdf",
		"www.google.co.uk"}

	assert.Equal(t, true, contains(strArray, "www.google.co.uk"))
	assert.Equal(t, false, contains(strArray, "www.hiphop.co.uk"))
}
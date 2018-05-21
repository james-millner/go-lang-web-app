package web

import (
	"testing"
	"github.com/stretchr/testify/assert"

)

func TestLinkHasSuffix(t *testing.T) {
	strArray := []string{"https://thenextweb.com/contributors/2018/05/18/15-online-trends-to-watch-for-in-2018-and-beyond/", "https://thenextweb.com/section/insights/", "https://thenextweb.com/section/insights.pdf"}

	assert.Equal(t, true, CheckLinkHasSuffix(strArray[0], "/"))
	assert.Equal(t, false, CheckLinkHasSuffix(strArray[0], ".pdf"))

	assert.Equal(t, true, CheckLinkHasSuffix(strArray[1], "/"))
	assert.Equal(t, false, CheckLinkHasSuffix(strArray[1], ".pdf"))

	assert.Equal(t, false, CheckLinkHasSuffix(strArray[2], "/"))
	assert.Equal(t, true, CheckLinkHasSuffix(strArray[2], ".pdf"))
}
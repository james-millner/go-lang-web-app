package web

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
)

func TestLinkHasSuffix(t *testing.T) {
	strArray := []string{
		"https://thenextweb.com/contributors/2018/05/18/15-online-trends-to-watch-for-in-2018-and-beyond/",
		"https://thenextweb.com/section/insights/",
		"https://thenextweb.com/section/insights.pdf",
		"www.google.co.uk"}

	assert.Equal(t, true, strings.HasSuffix(strArray[0], "/"))
	assert.Equal(t, false, strings.HasSuffix(strArray[0], ".pdf"))

	assert.Equal(t, true, strings.HasSuffix(strArray[1], "/"))
	assert.Equal(t, false, strings.HasSuffix(strArray[1], ".pdf"))

	assert.Equal(t, false, strings.HasSuffix(strArray[2], "/"))
	assert.Equal(t, true, strings.HasSuffix(strArray[2], ".pdf"))

	assert.Equal(t, true, strings.HasSuffix(strArray[3], ".co.uk"))
	assert.Equal(t, true, strings.HasSuffix(strArray[3], "uk"))
	assert.Equal(t, true, strings.HasSuffix(strArray[3], "google.co.uk"))
	assert.Equal(t, false, strings.HasSuffix(strArray[3], "google"))
	assert.Equal(t, false, strings.HasSuffix(strArray[3], ".com"))
}

func TestGetPageLinks(t *testing.T) {
	b, err := ioutil.ReadFile("../../assets/test_data/thenextweb-homepage.html")
	assert.NoError(t, err)

	s := string(b[:])

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(s))

	if err != nil {
		assert.Fail(t, "Couldn't read document.")
	}

	links := GetLinks(doc)

	for l := range links {
		fmt.Println(links[l])
	}

	assert.Equal(t, 4, len(links))
}

func TestIsValidCaseStudyURL(t *testing.T) {
	assert.Equal(t, false, isValidCaseStudyLink("https://www.iqblade.com", false))
	assert.Equal(t, true, isValidCaseStudyLink("https://www.iqblade.com/case-studies", false))
	assert.Equal(t, true, isValidCaseStudyLink("https://www.iqblade.com/customers", false))
	assert.Equal(t, true, isValidCaseStudyLink("https://www.iqblade.com/customers", false))
	assert.Equal(t, true, isValidCaseStudyLink("https://uk.cdw.com/files/9115/0832/7959/CDW_-_Cloud_Spence_Case_Study.pdf", true))
}

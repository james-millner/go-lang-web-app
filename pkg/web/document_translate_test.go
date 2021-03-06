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
		"https://thenextcom/contributors/2018/05/18/15-online-trends-to-watch-for-in-2018-and-beyond/",
		"https://thenextcom/section/insights/",
		"https://thenextcom/section/insights.pdf",
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
	b, err := ioutil.ReadFile("../../assets/test_data/dummy-web.html")
	assert.NoError(t, err)

	s := string(b[:])

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(s))

	if err != nil {
		assert.Fail(t, "Couldn't read document.")
	}

	links := RetreiveLinksFromDocument("https://test.co.uk", doc)

	for l := range links {
		fmt.Println(links[l])
	}

	assert.Equal(t, 5, len(links))
}

func TestIsPossibleCaseStudyURL(t *testing.T) {
	assert.Equal(t, false, IsPossibleCaseStudyLink("https://www.iqblade.com"))
	assert.Equal(t, true, IsPossibleCaseStudyLink("https://www.iqblade.com/case-studies"))
	assert.Equal(t, true, IsPossibleCaseStudyLink("https://www.iqblade.com/customers"))
	assert.Equal(t, true, IsPossibleCaseStudyLink("https://www.iqblade.com/customers"))
	assert.Equal(t, true, IsPossibleCaseStudyLink("https://www.iqblade.com/wp-content/upload"))
	assert.Equal(t, true, IsPossibleCaseStudyLink("https://uk.cdw.com/files/9115/0832/7959/CDW_-_Cloud_Spence_Case_Study.pdf"))
}

func TestIsPDFDocument(t *testing.T) {
	assert.Equal(t, false, IsPDFDocument("https://www.iqblade.com"))
	assert.Equal(t, true, IsPDFDocument("https://uk.cdw.com/files/9115/0832/7959/CDW_-_Cloud_Spence_Case_Study.pdf"))
	assert.Equal(t, false, IsPDFDocument("https://uk.cdw.com/files/9115/0832/7959/CDW_-_Cloud_Spence_Case_Study"))
	assert.Equal(t, true, IsPDFDocument("https://uk.cdw.com/files/9115/.pdf/7959/CDW_-_Cloud_Spence_Case_Study"))
	assert.Equal(t, true, IsPDFDocument("https://media.cobcom/site-library/docs/default-source/case-studies-azure/leadent-solutions.pdf?sfvrsn=da1f30ab_6"))
}

func TestIsProbableLink(t *testing.T) {
	assert.Equal(t, true, IsProbableLink("https://www.iqblade.com"))
	assert.Equal(t, false, IsProbableLink("https://www.twitter.com/iqblade"))
	assert.Equal(t, false, IsProbableLink("https://www.linkedin.com/iqblade"))
	assert.Equal(t, true, IsProbableLink("https://media.cobcom/site-library/docs/default-source/case-studies-azure/leadent-solutions.pdf?sfvrsn=da1f30ab_6"))
}

func TestGetFileName(t *testing.T) {
	assert.Equal(t, "Risk Business", GetFileName("https://media.cobweb.com/site-library/docs/default-source/case-studies-azure/risk-business.pdf?sfvrsn=608c30ab_4"))
	assert.Equal(t, "Mayday IT Support Managed Services Case Study © Urban Network ", GetFileName("https://www.urbannetwork.co.uk/wp-content/uploads/2017/08/Mayday-IT-Support-Managed-Services-Case-Study-%C2%A9-Urban-Network-.pdf?x86826"))
	assert.Equal(t, "Splash", GetFileName("https://media.cobweb.com/site-library/docs/default-source/case-studies-azure/splash.pdf?sfvrsn=459730ab_2"))
	assert.Equal(t, "Lauras International", GetFileName("https://media.cobweb.com/site-library/docs/default-source/case-studies-azure/lauras-international.pdf?sfvrsn=de1f30ab_6"))
}

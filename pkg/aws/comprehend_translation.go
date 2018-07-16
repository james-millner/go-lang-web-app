package aws

import (
	"log"

	"github.com/aws/aws-sdk-go-v2/service/comprehend"
	"github.com/james-millner/go-lang-web-app/pkg/web"
)

//ComprehendDTO
type ComprehendDTO struct {
	Text string
	Tag  string
}

//DetermineOrganisationTag method
func DetermineOrganisationTag(output *comprehend.BatchDetectEntitiesOutput) []string {

	results := output.ResultList

	returnedResults := []string{}

	for _, obj := range results {
		log.Println(obj.GoString())
		for _, o := range obj.Entities {

			tag, _ := o.Type.MarshalValue()

			if tag == "ORGANIZATION" && !web.ArrayContains(returnedResults, *o.Text) {
				returnedResults = append(returnedResults, *o.Text)
			}
		}
	}

	return returnedResults

}

package aws

import (
	"log"
	"reflect"

	"github.com/aws/aws-sdk-go-v2/service/comprehend"
)

//ComprehendDTO
type ComprehendDTO struct {
	Text string
	Tag  string
}

//DetermineOrganisationTag method
func DetermineOrganisationTag(output *comprehend.BatchDetectEntitiesOutput) []ComprehendDTO {

	results := output.ResultList

	returnedResults := []ComprehendDTO{}

	for _, obj := range results {
		log.Println(obj.GoString())
		for _, o := range obj.Entities {

			tag, _ := o.Type.MarshalValue()

			dto := ComprehendDTO{Text: *o.Text, Tag: tag}

			if !contains(returnedResults, dto) {
				returnedResults = append(returnedResults, dto)
			}
		}
	}

	return returnedResults

}

//ArrayContains method for determining if a String exists within a string array.
func contains(arr []ComprehendDTO, obj ComprehendDTO) bool {
	for _, a := range arr {
		if reflect.DeepEqual(a, obj) {
			return true
		}
	}
	return false
}

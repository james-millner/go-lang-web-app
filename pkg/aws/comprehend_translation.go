package aws

import (
<<<<<<< HEAD
	"github.com/james-millner/go-lang-web-app/pkg/web"

=======
>>>>>>> master
	"github.com/aws/aws-sdk-go-v2/service/comprehend"
	"github.com/james-millner/go-lang-web-app/pkg/web"
)

// ComprehendDTO
type ComprehendDTO struct {
	Text string
	Tag  string
}

// DetermineOrganisationTag method
func DetermineOrganisationTag(output *comprehend.BatchDetectEntitiesOutput) ([]string, []string, []string) {

	results := output.ResultList

	organisations := []string{}
	people := []string{}
	locations := []string{}

	for _, obj := range results {
		for _, o := range obj.Entities {

			if o.Type == comprehend.EntityTypeOrganization && !web.ArrayContains(organisations, *o.Text) {
				organisations = append(organisations, *o.Text)
			}

			if o.Type == comprehend.EntityTypePerson && !web.ArrayContains(people, *o.Text) {
				people = append(people, *o.Text)
			}

			if o.Type == comprehend.EntityTypeLocation && !web.ArrayContains(locations, *o.Text) {
				locations = append(locations, *o.Text)
			}
		}
	}

	return organisations, people, locations

}

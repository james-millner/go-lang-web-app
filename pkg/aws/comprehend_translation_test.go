package aws

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/comprehend"
	"github.com/stretchr/testify/assert"
)

func TestTranslate(t *testing.T) {

	entityOneText := "IQBlade"
	entityTwoText := "Elder"
	entityThreeText := "Elder"
	entityFourText := "Elder"

	entities := []comprehend.Entity{
		comprehend.Entity{Text: &entityOneText, Type: comprehend.EntityTypeOrganization},
		comprehend.Entity{Text: &entityTwoText, Type: comprehend.EntityTypeOrganization},
		comprehend.Entity{Text: &entityThreeText, Type: comprehend.EntityTypeOrganization},
		comprehend.Entity{Text: &entityFourText, Type: comprehend.EntityTypePerson},
	}

	itemResults := &comprehend.BatchDetectEntitiesItemResult{Entities: entities}

	arr := []comprehend.BatchDetectEntitiesItemResult{*itemResults}

	obj := &comprehend.BatchDetectEntitiesOutput{ResultList: arr}

	organisations, people, locations := DetermineOrganisationTag(obj)

	assert.Len(t, organisations, 2)
	assert.Equal(t, organisations[0], "IQBlade")
	assert.Equal(t, organisations[1], "Elder")
	assert.Equal(t, len(people), 1)
	assert.Equal(t, len(locations), 0)
}

package aws

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/comprehend"
	"github.com/stretchr/testify/assert"
)

func TestTranslate(t *testing.T) {

	entityOneText := "IQBlade"
	entityTwoText := "Elder"

	entities := []comprehend.Entity{
		comprehend.Entity{Text: &entityOneText, Type: comprehend.EntityTypeOrganization},
		comprehend.Entity{Text: &entityTwoText, Type: comprehend.EntityTypeOrganization},
	}

	itemResults := &comprehend.BatchDetectEntitiesItemResult{Entities: entities}

	arr := []comprehend.BatchDetectEntitiesItemResult{*itemResults}

	obj := &comprehend.BatchDetectEntitiesOutput{ResultList: arr}

	results := DetermineOrganisationTag(obj)

	assert.Len(t, results, 2)
	assert.Equal(t, results[0].Text, "IQBlade")
	assert.Equal(t, results[0].Tag, "ORGANIZATION")
	assert.Equal(t, results[1].Text, "Elder")
	assert.Equal(t, results[1].Tag, "ORGANIZATION")

}

package aws

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/comprehend"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/comprehendiface"
	"github.com/stretchr/testify/assert"
)

// Define a mock struct to be used in your unit tests of myFunc.
type mockComprehendClient struct {
	comprehendiface.ComprehendAPI
}

func (m *mockComprehendClient) BatchDetectEntitiesRequest(input *comprehend.BatchDetectEntitiesInput) comprehend.BatchDetectEntitiesRequest {
	lang := "en"

	input.TextList = []string{"IQBlade"}
	input.LanguageCode = &lang

	obj := &comprehend.BatchDetectEntitiesRequest{Input: input}

	return *obj
}

func TestComprehend(t *testing.T) {

	mockSvc := &mockComprehendClient{}

	str := []string{}

	res := DetectEntities(mockSvc, str)

	fmt.Println(res)

	assert.NotNil(t, res)
	assert.Len(t, res.Input.TextList, 1)
}

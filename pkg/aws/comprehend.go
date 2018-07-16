package aws

import (
	"github.com/aws/aws-sdk-go-v2/aws/endpoints"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/comprehend"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/comprehendiface"
)

//RunComprehend method
func RunComprehend(body []string) (*comprehend.BatchDetectEntitiesOutput, error) {

	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}
	cfg.Region = endpoints.EuWest1RegionID

	// Create a Comprehend client from just a session.
	svc := comprehend.New(cfg)

	request := DetectEntities(svc, body)

	return request.Send()
}

//DetectEntities method
func DetectEntities(svc comprehendiface.ComprehendAPI, text []string) comprehend.BatchDetectEntitiesRequest {
	lang := "en"

	input := &comprehend.BatchDetectEntitiesInput{LanguageCode: &lang, TextList: text}

	request := svc.BatchDetectEntitiesRequest(input)
	return request
}

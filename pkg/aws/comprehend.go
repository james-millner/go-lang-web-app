package aws

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws/endpoints"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/comprehend"
)

//RunComprehend
func RunComprehend(body []string) *comprehend.BatchDetectEntitiesOutput {

	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}
	cfg.Region = endpoints.EuWest1RegionID

	fmt.Println("Config made")

	// Create a Comprehend client from just a session.
	svc := comprehend.New(cfg)

	fmt.Println("Created client.")

	lang := "en"

	input := &comprehend.BatchDetectEntitiesInput{LanguageCode: &lang, TextList: body}

	fmt.Println(input.GoString())

	request := svc.BatchDetectEntitiesRequest(input)

	fmt.Println("Request created.")

	resp, err := request.Send()

	if err == nil {
		fmt.Println(resp)
		return resp
	}

	fmt.Println("Error ")
	return nil
}

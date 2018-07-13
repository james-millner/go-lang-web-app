package aws

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComprehend(t *testing.T) {

	testSet := []string{"Two former distribution executives have revealed ambitious growth plans for IQBlade, a market intelligence platform that aims to use machine-learning techniques to help vendors select partners.", "Young said the concept of IQBlade was born three years ago when he and Abraham were working with vendors through their consulting business, Demuto.", "IQBlade currently has 6,000 fully fleshed-out UK profiles. Although the firm currently focuses on the UK, Young said the platform's capabilities can be replicated in France and the Nordics."}

	results := RunComprehend(testSet)

	if results != nil {
		entities := results.ResultList

		for _, i := range entities {
			fmt.Println(i.GoString())

			for _, o := range i.Entities {
				fmt.Println(o.Text)
				str, err := o.Type.MarshalValue()

				if err == nil {
					fmt.Println(str)
				}

			}
		}
	} else {
		fmt.Println("Errorerererer!")
	}

	assert.Equal(t, true, true)
}

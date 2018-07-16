package es

import (
	"context"
	"fmt"
	"log"

	"github.com/james-millner/go-lang-web-app/pkg/model"
	"github.com/james-millner/go-lang-web-app/pkg/web"
	"github.com/olivere/elastic"
)

var caseStudyMapping = `{
	"settings":{
		"number_of_shards":1,
		"number_of_replicas":0
	},
	"mappings":{
		"casestudy":{
			"properties":{
				"id":{
					"type":"string",
					"index": "not_analyzed"
				},
				"sourceUrl": {
					"type":"string",
					"index": "not_analyzed"
				},
				"companyNumber": {
					"type":"string",
					"index": "not_analyzed"
				},
				"caseStudyText":{
					"type":"string",
				},
				"organisations": {
					"type": "string",
					"index": "not_analyzed"
				},
				"createdAt" {
					"type": "date"
				},
				"updatedAt" {
					"type": "date"
				},
				"identifiedOn" {
					"type": "date"
				}
			}
		}
	}
}`

// Elasticsearch provides the interface of an Elasticsearch client
type Elasticsearch interface {
	PutRecord(ctx context.Context, study model.CaseStudy) error
}

// Elastic handles the communications to and from the Elasticsearch services used by this application
type Elastic struct {
	client *elastic.Client
}

// New creates a new ES struct for communicating with the ES Client
func New(client *elastic.Client) *Elastic {
	return &Elastic{
		client: client,
	}
}

// PutRecord receives a byte slice for putting a record into Elasticsearch
func (e *Elastic) PutRecord(ctx context.Context, study model.CaseStudy) error {
	index := "organisation-case-studies"

	err := e.ensureIndexExists(ctx, index)
	if err != nil {
		log.Fatalf("Error: %v", err)
		return err
	}

	dto := web.TranslateToElastic(study)

	e.client.Index().
		Index(index).
		Type("casestudy").
		Id(dto.ID).
		BodyJson(dto).
		Do(context.Background())

	log.Println("PUT!")

	return nil
}

func (e *Elastic) ensureIndexExists(ctx context.Context, i string) error {
	exists, err := e.client.IndexExists(i).Do(ctx)
	if err != nil {
		return fmt.Errorf("failed to retrieve existing index: %v", err)
	}
	if !exists {
		r, err := e.client.CreateIndex(i).Do(ctx)
		if err != nil {
			return fmt.Errorf("failed to create new index %s: %v", i, err)
		}

		if !r.Acknowledged {
			return fmt.Errorf("elasticsearch did not acknowledge creation of index %s: %v", i, err)
		}
	}

	return nil
}

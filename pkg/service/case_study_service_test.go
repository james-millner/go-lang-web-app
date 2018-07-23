package service

import (
	"testing"

	"github.com/google/go-tika/tika"
	"github.com/stretchr/testify/assert"

	"github.com/james-millner/go-lang-web-app/pkg/db"
)

func TestNewService(t *testing.T) {

	db := &db.DB{}
	tc := &tika.Client{}

	es := New(db, tc)

	assert.NotNil(t, es.DB)
	assert.NotNil(t, es.TikaClient)

}
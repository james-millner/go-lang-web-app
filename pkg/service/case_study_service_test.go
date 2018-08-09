package service

import (
	"testing"

	"github.com/google/go-tika/tika"
	"github.com/iqblade/casestudies/pkg/db"
	"github.com/stretchr/testify/assert"
)

func TestNewService(t *testing.T) {

	db := &db.DB{}
	tc := &tika.Client{}

	es := New(db, tc)

	assert.NotNil(t, es.DB)
	assert.NotNil(t, es.TikaClient)

}

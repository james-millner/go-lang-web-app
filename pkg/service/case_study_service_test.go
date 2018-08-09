package service

import (
	"testing"

	"github.com/google/go-tika/tika"
<<<<<<< HEAD
	"github.com/stretchr/testify/assert"

	"github.com/james-millner/go-lang-web-app/pkg/db"
=======
	"github.com/iqblade/casestudies/pkg/db"
	"github.com/stretchr/testify/assert"
>>>>>>> master
)

func TestNewService(t *testing.T) {

	db := &db.DB{}
	tc := &tika.Client{}

	es := New(db, tc)

	assert.NotNil(t, es.DB)
	assert.NotNil(t, es.TikaClient)

}

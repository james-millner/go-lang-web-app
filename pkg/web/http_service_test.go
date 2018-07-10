package web

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetGoqueryDocument_Success(t *testing.T) {
	b, err := ioutil.ReadFile("../../assets/test_data/dummy-web.html")
	assert.NoError(t, err)

	r := bytes.NewReader(b)

	d, err := GetGoqueryDocument(r)

	assert.NotNil(t, d)
	assert.Nil(t, err)
}

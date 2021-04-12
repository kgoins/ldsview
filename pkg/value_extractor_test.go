package ldsview

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValueExtractor_GetValues(t *testing.T) {
	a := assert.New(t)

	parser := NewLdifParser(TESTFILE)
	a.NotNil(parser)

	vals, err := GetValues(&parser, "cn")
	a.NoError(err)
	a.Equal(NUMENTITIES, len(vals))
}

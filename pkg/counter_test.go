package ldsview

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLdifParser_CountEntities(t *testing.T) {
	a := assert.New(t)

	parser := NewLdifParser(TESTFILE)
	a.NotNil(parser)

	count, err := parser.CountEntities()
	a.NoError(err)
	a.Equal(count, NUMENTITIES)
}

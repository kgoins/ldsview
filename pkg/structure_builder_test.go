package ldsview

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStructureBuilder_GetStructure(t *testing.T) {
	parser := NewLdifParser(TESTFILE)

	t.Run("parses the ldif objects correctly", func(t *testing.T) {
		structure, err := GetStructure(&parser)
		assert.Nil(t, err)
		assert.Greater(t, len(structure), 0)
	})
}

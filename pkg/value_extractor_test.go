package ldsview

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValueExtractor_GetValues(t *testing.T) {
	parser := NewLdifParser(TESTFILE)

	t.Run("parses the ldif objects correctly", func(t *testing.T) {
		structure, err := GetValues(&parser, "cn")
		assert.Nil(t, err)
		assert.Greater(t, len(structure), 0)
	})
}

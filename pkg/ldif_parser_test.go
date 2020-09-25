package ldsview

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var parser = NewLdifParser(TESTFILE)

func TestLdifParser(t *testing.T) {

	t.Run("counts the ldif objects correctly", func(t *testing.T) {
		want := 2
		got, err := parser.CountEntities()
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	})

	t.Run("counts the filtered ldif objects correctly", func(t *testing.T) {
		entityFilterStr := []string{"objectClass:=computer"}
		entityFilter, _ := BuildEntityFilter(entityFilterStr)
		parser.SetEntityFilter(entityFilter)

		want := 1
		got, err := parser.CountEntities()
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	})

}

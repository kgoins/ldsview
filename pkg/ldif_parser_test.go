package ldsview

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildEntities(t *testing.T) {
	var parser = NewLdifParser(TESTFILE)
	t.Run("parses the ldif objects correctly", func(t *testing.T) {
		entities := make(chan Entity)
		done := make(chan bool)

		filter, _ := BuildEntityFilter([]string{"cn:=*"})
		parser.SetEntityFilter(filter)

		type count struct{ c int }
		counter := &count{0}

		go func(ents chan Entity, c *count) {
			for _ = range ents {
				c.c = c.c + 1
			}
			done <- true
		}(entities, counter)

		err := parser.BuildEntities(entities, done)
		assert.Nil(t, err)
		assert.Equal(t, 3, counter.c)
	})
}

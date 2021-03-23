package ldsview

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
			for range ents {
				c.c = c.c + 1
			}
			done <- true
		}(entities, counter)

		err := parser.BuildEntities(entities, done)
		assert.Nil(t, err)
		assert.Equal(t, 3, counter.c)
	})
}

func TestBuildEntityWithEscapeChar(t *testing.T) {
	rq := require.New(t)
	var parser = NewLdifParser("../testdata/test_user_escape_char.ldif")

	user, err := parser.BuildEntity("sAMAccountName", "myuser")
	rq.NoError(err)
	rq.False(user.IsEmpty())

	_, worked := user.GetDN()
	rq.True(worked)
}

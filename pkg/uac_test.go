package ldsview

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUACPrint(t *testing.T) {
	t.Run("prints correct sorted output", func(t *testing.T) {
		want := "b5106b8639687bb965a84af85e69113a"
		buffer := &bytes.Buffer{}
		UACPrint(buffer)

		got := fmt.Sprintf("%x", md5.Sum(buffer.Bytes()))
		assert.Equal(t, want, got)
	})
}

func TestUACParse(t *testing.T) {
	t.Run("parses and calculates the correct values", func(t *testing.T) {

		want := []string{"NORMAL_ACCOUNT", "SCRIPT"}
		got, err := UACParse("513")
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	})

	t.Run("returns an error when invalid input is passed", func(t *testing.T) {
		got, err := UACParse("I AM NOT A NUMBER!")
		assert.Nil(t, got)
		assert.NotNil(t, err)
	})
}

func TestUACSearch(t *testing.T) {
	t.Run("returns an error when passed an invalid <file> param", func(t *testing.T) {
		got, err := UACSearch("noop", 512)
		assert.Nil(t, got)
		assert.NotNil(t, err)
	})

	t.Run("returns an error when passed an invalid <uacProp> param", func(t *testing.T) {
		//TODO:needs in/valid ldif test fixture
	})

	t.Run("returns a correct entity when there is a match preset", func(t *testing.T) {
		//TODO:needs in/valid ldif test fixture
		// got, err := UACSearch("test.ldif", 512)
	})

	t.Run("returns no entity when no match is preset", func(t *testing.T) {
		//TODO:needs in/valid ldif test fixture
		// got, err := UACSearch("test.ldif", 1)
	})
}

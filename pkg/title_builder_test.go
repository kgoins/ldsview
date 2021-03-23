package ldsview

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuildTitles(t *testing.T) {
	rq := require.New(t)
	var parser = NewLdifParser("../testdata/test_user_escape_char.ldif")

	user, err := parser.BuildEntity("sAMAccountName", "myuser")
	rq.NoError(err)

	title, err := BuildTitleLine(user)
	rq.NoError(err)
	rq.NotEmpty(title)
}

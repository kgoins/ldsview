package ldsview_test

import (
	"os"
	"testing"

	ldsview "github.com/kgoins/ldsview/pkg"
	"github.com/kgoins/ldsview/pkg/searcher"
	"github.com/stretchr/testify/require"
)

func TestGetUniqueVals(t *testing.T) {
	r := require.New(t)

	testFile, err := os.Open(TESTFILE)
	r.NoError(err)
	defer testFile.Close()

	searcher := searcher.NewLdifSearcher(testFile)
	r.NotNil(searcher)

	vals, err := ldsview.GetUniqueValues(searcher, "sAMAccountType")
	r.NoError(err)
	r.Equal(1, len(vals))
}

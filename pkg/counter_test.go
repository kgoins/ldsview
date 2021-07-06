package ldsview_test

import (
	"os"
	"testing"

	ldsview "github.com/kgoins/ldsview/pkg"
	"github.com/kgoins/ldsview/pkg/searcher"
	"github.com/stretchr/testify/require"
)

func TestLdifParser_CountEntities(t *testing.T) {
	r := require.New(t)

	testFile, err := os.Open(TESTFILE)
	r.NoError(err)
	defer testFile.Close()

	searcher := searcher.NewLdifSearcher(testFile)
	r.NotNil(searcher)

	count, err := ldsview.CountEntities(searcher)
	r.NoError(err)
	r.Equal(count, NUMENTITIES)
}

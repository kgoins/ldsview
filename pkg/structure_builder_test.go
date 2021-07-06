package ldsview_test

import (
	"os"
	"testing"

	ldsview "github.com/kgoins/ldsview/pkg"
	"github.com/kgoins/ldsview/pkg/searcher"
	"github.com/stretchr/testify/require"
)

func TestStructureBuilder_GetStructure(t *testing.T) {
	r := require.New(t)

	testFile, err := os.Open(TESTFILE)
	r.NoError(err)
	defer testFile.Close()

	searcher := searcher.NewLdifSearcher(testFile)
	r.NotNil(searcher)

	structure, err := ldsview.GetStructure(searcher)
	r.NoError(err)
	r.NotEmpty(structure)
}

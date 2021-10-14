package searcher_test

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	filter "github.com/kgoins/entityfilter/entityfilter"
	"github.com/kgoins/ldapentity/entity"
	"github.com/kgoins/ldsview/pkg/searcher"
	"github.com/stretchr/testify/require"
)

func getTestDataDir() (string, error) {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		return "", errors.New("reflection failed")
	}

	parent := filepath.Dir(file)
	pkgRoot := filepath.Dir(parent)
	projRoot := filepath.Dir(pkgRoot)

	return filepath.Join(projRoot, "testdata"), nil
}

func getTestUsersFile(testdataDir string) string {
	return filepath.Join(testdataDir, "test_users.ldif")
}

func TestLdifSearcher_SearchEntities(t *testing.T) {
	r := require.New(t)

	testDataDir, err := getTestDataDir()
	r.NoError(err)
	testUsersFile := getTestUsersFile(testDataDir)
	r.NoError(err)

	testFile, err := os.Open(testUsersFile)
	r.NoError(err)
	defer testFile.Close()

	searcher := searcher.NewLdifSearcher(testFile)
	r.NotNil(searcher)

	entityFilter, err := filter.ParseFilterStr("objectClass:=computer,cn:=MYPC")
	r.NoError(err)

	done := make(chan bool)
	resultStream := searcher.SearchEntities(done, nil, entityFilter)

	entities := []entity.Entity{}
	for result := range resultStream {
		r.NoError(result.Error)
		entities = append(entities, result.Entity)
	}

	r.Len(entities, 1)
	result := entities[0]
	cn, found := result.GetSingleValuedAttribute("cn")
	r.True(found)
	r.Equal("MYPC", cn)
}

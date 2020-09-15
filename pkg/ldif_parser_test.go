package ldsview_test

import (
	"testing"

	ldsview "github.com/kgoins/ldsview/pkg"
)

func TestLdifParser_CountEntities(t *testing.T) {
	parser := ldsview.NewLdifParser(TESTFILE)

	count, err := parser.CountEntities()
	if err != nil {
		t.Fatalf("unable to parse test file")
	}

	if count != 5 {
		t.Fatalf("Failed to count entities")
	}
}

func TestLdifParser_CountEntitiesWithFilter(t *testing.T) {
	parser := ldsview.NewLdifParser(TESTFILE)

	entityFilterStr := []string{"objectClass:=computer"}
	entityFilter, _ := ldsview.BuildEntityFilter(entityFilterStr)
	parser.SetEntityFilter(entityFilter)

	count, err := parser.CountEntities()
	if err != nil {
		t.Fatalf("unable to parse test file")
	}

	if count != 1 {
		t.Fatalf("Failed to count filtered entities")
	}
}

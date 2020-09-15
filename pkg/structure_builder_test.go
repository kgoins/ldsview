package ldsview_test

import (
	"testing"

	ldsview "github.com/kgoins/ldsview/pkg"
)

func TestStructureBuilder_GetStructure(t *testing.T) {
	parser := ldsview.NewLdifParser(TESTFILE)

	structure, err := ldsview.GetStructure(&parser)
	if err != nil {
		t.Fatalf("failed to parse ldif for domain structure")
	}

	if len(structure) == 0 {
		t.Fatalf("failed to parse ldif for domain structure")
	}
}

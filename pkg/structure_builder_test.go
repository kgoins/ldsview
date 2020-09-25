package ldsview

import (
	"testing"
)

func TestStructureBuilder_GetStructure(t *testing.T) {
	parser := NewLdifParser(TESTFILE)

	structure, err := GetStructure(&parser)
	if err != nil {
		t.Fatalf("failed to parse ldif for domain structure")
	}

	if len(structure) == 0 {
		t.Fatalf("failed to parse ldif for domain structure")
	}
}

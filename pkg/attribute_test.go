package ldsview_test

import (
	"testing"

	ldsview "github.com/kgoins/ldsview/pkg"
)

func TestAttribute_BuildFromValidLine(t *testing.T) {
	attrLine := "userAccountControl: 66048"

	attr, err := ldsview.BuildAttributeFromLine(attrLine)
	if err != nil {
		t.Fatalf("Unable to build from valid attr line")
	}

	if attr.Name != "userAccountControl" {
		t.Fatalf("Failed to parse attr name")
	}

	if attr.Value.Size() != 1 {
		t.Fatalf("Failed to parse attr value")
	}

	if attr.Value.GetSingleValue() != "66048" {
		t.Fatalf("Failed to parse attr value")
	}
}

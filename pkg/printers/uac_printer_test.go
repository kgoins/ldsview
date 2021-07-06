package printers_test

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"testing"

	"github.com/kgoins/ldsview/pkg/printers"
	"github.com/stretchr/testify/require"
)

func TestUACPrint(t *testing.T) {
	r := require.New(t)

	want := "b5106b8639687bb965a84af85e69113a"
	buffer := &bytes.Buffer{}
	printers.PrintAllUACFlags(buffer)

	got := fmt.Sprintf("%x", md5.Sum(buffer.Bytes()))
	r.Equal(want, got)
}

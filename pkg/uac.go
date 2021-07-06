package ldsview

import (
	"fmt"
	"io"
	"sort"
	"strconv"
	"text/tabwriter"

	"github.com/audibleblink/msldapuac"
)

// GetFlagsFromUAC wraps the msldapuac dependency for easy re-definition
func GetFlagsFromUAC(uac int64) ([]string, error) {
	return msldapuac.ParseUAC(uac)
}

// UACPrint prints the available UAC options that are available for searching
func UACPrint(dest io.Writer) {
	w := new(tabwriter.Writer)
	w.Init(dest, 8, 8, 0, '\t', 0)
	defer w.Flush()

	template := "%s\t%d\n"
	var sorted []string
	for k, v := range msldapuac.PropertyMap {
		sorted = append(sorted, fmt.Sprintf(template, v, k))
	}

	sort.Strings(sorted)
	fmt.Fprintf(w, "Property\tValue\n")
	fmt.Fprintf(w, "---\t---\n")
	for _, line := range sorted {
		fmt.Fprintf(w, line)
	}
}

// UACParse take a UAC int and return the names of the flags it has set
func UACParse(uacValue string) (flagNames []string, err error) {
	uacInt, err := strconv.ParseInt(uacValue, 10, 64)
	if err != nil {
		return
	}

	flagNames, err = GetFlagsFromUAC(uacInt)
	if err != nil {
		return
	}
	return
}

package printers

import (
	"fmt"
	"io"
	"sort"
	"text/tabwriter"

	"github.com/audibleblink/msldapuac"
)

// PrintUAC prints the available UAC options that are available for searching
func PrintAllUACFlags(dest io.Writer) {
	w := new(tabwriter.Writer)
	w.Init(dest, 8, 8, 0, '\t', 0)
	defer w.Flush()

	template := "%s\t%d\n"
	var sorted []string
	for k, v := range msldapuac.PropertyMap {
		sorted = append(sorted, fmt.Sprintf(template, v, k))
	}

	sort.Strings(sorted)
	fmt.Fprint(w, "Property\tValue\n")
	fmt.Fprint(w, "---\t---\n")
	for _, line := range sorted {
		fmt.Fprint(w, line)
	}
}

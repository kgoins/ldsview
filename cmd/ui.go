package cmd

import (
	"fmt"
	"os"
	"sort"

	ldsview "github.com/kgoins/ldsview/pkg"
)

func PrintAttribute(attr ldsview.EntityAttribute) {
	for _, line := range attr.Stringify() {
		fmt.Println(line)
	}
}

func PrintEntity(entity ldsview.Entity) {
	titleLine, err := BuildTitleLine(entity)
	if err != nil {
		os.Stderr.WriteString("Skipping output of malformed object\n")
	}

	fmt.Println(titleLine)

	attrNames := entity.GetAllAttributeNames()

	sort.Strings(attrNames)

	for _, name := range attrNames {
		attr, _ := entity.GetAttribute(name)
		PrintAttribute(attr)
	}

	fmt.Println()
}

package cmd

import (
	"fmt"
	"os"
	"sort"

	"github.com/kgoins/ldapentity/entity"
	"github.com/kgoins/ldifparser"
	"github.com/kgoins/ldifparser/syntax"
	"github.com/kgoins/ldsview/pkg/searcher"
	"github.com/spf13/cobra"
)

func PrintAttribute(attr entity.Attribute) {
	for _, line := range ldifparser.StringifyAttribute(attr) {
		fmt.Println(line)
	}
}

func PrintEntity(entity entity.Entity, decodeTS bool) {
	titleLine, err := syntax.BuildTitleLine(entity)
	if err != nil {
		os.Stderr.WriteString("Skipping output of malformed object\n")
		return
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

func ChannelPrinter(entities <-chan searcher.EntityResult, cmd *cobra.Command) (err error) {
	tdc, _ := cmd.Flags().GetBool("tdc")

	printLimit, intParseErr := cmd.Flags().GetInt("first")
	if intParseErr != nil {
		err = fmt.Errorf("unable to parse value: %s", intParseErr.Error())
		return
	}

	entCount := 0
	for entity := range entities {
		if entity.Error != nil {
			return entity.Error
		}

		entCount = entCount + 1
		PrintEntity(entity.Entity, tdc)

		if entCount == printLimit {
			break
		}
	}

	return
}

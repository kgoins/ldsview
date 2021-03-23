package cmd

import (
	"fmt"
	"os"
	"sort"

	ldsview "github.com/kgoins/ldsview/pkg"
	"github.com/spf13/cobra"
)

func PrintAttribute(attr ldsview.EntityAttribute) {
	for _, line := range attr.Stringify() {
		fmt.Println(line)
	}
}

func PrintEntity(entity ldsview.Entity, decodeTS bool) {
	titleLine, err := BuildTitleLine(entity)
	if err != nil {
		os.Stderr.WriteString("Skipping output of malformed object\n")
		return
	}

	fmt.Println(titleLine)

	attrNames := entity.GetAllAttributeNames()

	if decodeTS {
		entity.DeocdeTimestamps()
	}

	sort.Strings(attrNames)

	for _, name := range attrNames {
		attr, _ := entity.GetAttribute(name)
		PrintAttribute(attr)
	}

	fmt.Println()
}

// ChannelPrinter concurrently prints entity results and signals shared `done` channel
// when finished
func ChannelPrinter(entities chan ldsview.Entity, done chan bool, cmd *cobra.Command) {

	count, _ := cmd.Flags().GetBool("count")
	tdc, _ := cmd.Flags().GetBool("tdc")

	printLimit, intParseErr := cmd.Flags().GetInt("first")
	if intParseErr != nil {
		fmt.Printf("Unable to parse value: %s\n", intParseErr.Error())
		done <- true
		return
	}

	entCount := 0

	for entity := range entities {
		entCount = entCount + 1

		if !count {
			PrintEntity(entity, tdc)
		}

		if entCount == printLimit {
			break
		}
	}

	if count {
		fmt.Println("Entities: ", entCount)
	}
	done <- true
}

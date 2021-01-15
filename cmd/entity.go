package cmd

import (
	"bufio"
	"log"
	"os"

	"github.com/spf13/cobra"
	ldsview "github.com/kgoins/ldsview/pkg"
)

// entityCmd represents the entity command
var entityCmd = &cobra.Command{
	Use:   "entity keyValue",
	Short: "Extract an ldap object with a given attribute value",
	Run: func(cmd *cobra.Command, args []string) {
		dumpFile, _ := cmd.Flags().GetString("file")

		keyAttr, _ := cmd.Flags().GetString("key-attr")
		parseTdc, _ := cmd.Flags().GetBool("tdc")

		ldifParser := ldsview.NewLdifParser(dumpFile)

		filterParts, _ := cmd.Flags().GetStringSlice("include")
		filter := ldsview.BuildAttributeFilter(filterParts)
		ldifParser.SetAttributeFilter(filter)

		inputList, _ := cmd.Flags().GetString("input-file")
		if inputList == "" {
			if len(args) != 1 {
				log.Fatalf("Key value must be specified")
				return
			}

			keyValue := args[0]
			execEntityCmd(ldifParser, keyAttr, keyValue, parseTdc)

			return
		}

		entityKeys, err := getEntityKeysFromInputFile(inputList)
		if err != nil {
			log.Fatalf("unable to read input file")
		}

		for _, key := range entityKeys {
			execEntityCmd(ldifParser, keyAttr, key, parseTdc)
		}
	},
}

func getEntityKeysFromInputFile(inputListPath string) ([]string, error) {
	file, err := os.Open(inputListPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var entityKeys []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		entityKeys = append(entityKeys, scanner.Text())
	}

	return entityKeys, scanner.Err()
}

func execEntityCmd(ldifParser ldsview.LdifParser, keyAttr string, keyValue string, parseTdc bool) {
	entity, err := ldifParser.BuildEntity(keyAttr, keyValue)
	if err != nil {
		log.Printf("%s\n", err.Error())
		return
	}

	if entity.IsEmpty() {
		log.Printf("Entity %s not found\n", keyValue)
		return
	}

	PrintEntity(entity, parseTdc)
}

func init() {
	rootCmd.AddCommand(entityCmd)

	entityCmd.PersistentFlags().StringP(
		"key-attr",
		"k",
		"sAMAccountName",
		"The attribute used to identify the target entity",
	)

	entityCmd.PersistentFlags().StringSliceP(
		"include",
		"i",
		[]string{},
		"Select which attributes are displayed from the returned entities",
	)

	entityCmd.PersistentFlags().Bool(
		"tdc",
		false,
		"Decodes timestamps to a human readable format",
	)

	entityCmd.PersistentFlags().String(
		"input-file",
		"",
		"Specifies the path to a list of entity keys to be searched",
	)
}

package cmd

import (
	"bufio"
	"log"
	"os"

	"github.com/kgoins/ldsview/internal"
	"github.com/kgoins/ldsview/pkg/searcher"
	"github.com/spf13/cobra"
)

// entityCmd represents the entity command
var entityCmd = &cobra.Command{
	Use:   "entity keyValue",
	Short: "Extract an ldap object with a given attribute value",
	Run: func(cmd *cobra.Command, args []string) {
		svcs, err := internal.BulidContainerFromFlags(cmd)
		if err != nil {
			log.Fatal(err)
		}

		searcher := svcs.Get("ldapsearcher").(searcher.LdapSearcher)

		keyAttr, _ := cmd.Flags().GetString("key-attr")
		parseTdc, _ := cmd.Flags().GetBool("tdc")
		inputList, _ := cmd.Flags().GetString("input-file")

		// Get single entity
		if inputList == "" {
			if len(args) != 1 {
				log.Fatalf("key value must be specified")
				return
			}

			keyValue := args[0]
			execEntityCmd(searcher, keyAttr, keyValue, parseTdc)

			return
		}

		// Get multiple entities from list in file
		entityKeys, err := getEntityKeysFromInputFile(inputList)
		if err != nil {
			log.Fatalf("unable to read input file")
		}

		for _, key := range entityKeys {
			execEntityCmd(searcher, keyAttr, key, parseTdc)
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

func execEntityCmd(searcher searcher.LdapSearcher, keyAttr string, keyValue string, parseTdc bool) {
	entity, err := searcher.ReadEntity(keyAttr, keyValue)
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

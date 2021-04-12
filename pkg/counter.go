package ldsview

import (
	"errors"
	"os"
)

// CountEntities returns the number of entities in the input file
func (parser LdifParser) CountEntities() (count int, err error) {
	Logger.Info("Opening ldif file: " + parser.filename)
	dumpFile, err := os.Open(parser.filename)
	if err != nil {
		return
	}
	defer dumpFile.Close()

	Logger.Info("Finding first entity block")
	entityScanner := parser.findFirstEntityBlock(dumpFile)
	if entityScanner == nil {
		return count, errors.New("Unable to find first entity block")
	}

	for entityScanner.Scan() {
		titleLine := entityScanner.Text()
		if parser.isEntityTitle(titleLine) {
			count++
		}
	}

	return
}

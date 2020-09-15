package ldsview

import (
	"bufio"
	"errors"
	"os"
	"regexp"
	"strings"

	"github.com/icza/backscanner"
	"github.com/kgoins/ldsview/internal"
)

// LdifParser implements EntityBuilder and constructs LDAP Entities
// from an ldif file.
type LdifParser struct {
	filename string

	entityFilter    []IEntityFilter
	attributeFilter internal.HashSetStr

	titleLineRegex    *regexp.Regexp
	scannerBufferSize int
}

// NewLdifParser returns a constructed LdifEntityBuilder with null filters.
func NewLdifParser(filename string) LdifParser {
	regex, _ := regexp.Compile(`^# .*\.`)

	return LdifParser{
		filename:          filename,
		entityFilter:      []IEntityFilter{},
		attributeFilter:   internal.NewHashSetStr(),
		titleLineRegex:    regex,
		scannerBufferSize: 1024 * 1024 * 10,
	}
}

// SetEntityFilter modifies the parser to return only entities matching
// the attribute / value pairs in the filter.
func (parser *LdifParser) SetEntityFilter(filter []IEntityFilter) {
	parser.entityFilter = filter
}

// SetAttributeFilter modifies the parser to return only the ldap
// attributes present in the filter on entites that are parsed.
func (parser *LdifParser) SetAttributeFilter(filter internal.HashSetStr) {
	parser.attributeFilter = filter
}

func (parser LdifParser) getEntityFromBlock(entityBlock *bufio.Scanner) Entity {
	entityLines := []string{}

	for entityBlock.Scan() {
		if parser.isEntitySeparator(entityBlock.Text()) {
			break
		}

		entityLines = append(entityLines, entityBlock.Text())
	}

	entity := BuildEntityFromAttrList(entityLines, &parser.attributeFilter)
	return entity
}

func (parser LdifParser) findFirstEntityBlock(dumpFile *os.File) *bufio.Scanner {
	scanner := bufio.NewScanner(dumpFile)
	buf := make([]byte, 0, parser.scannerBufferSize)
	scanner.Buffer(buf, parser.scannerBufferSize)

	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == "" {
			return scanner
		}
	}

	return nil
}

func (parser LdifParser) isEntityTitle(line string) bool {
	return parser.titleLineRegex.MatchString(line)
}

func (parser LdifParser) isEntitySeparator(line string) bool {
	return strings.TrimSpace(line) == ""
}

// Iterates the input scanner until an entity block is found.
// Returns false if the end of the input stream was reached.
func (parser LdifParser) getNextEntityBlock(scanner *PositionedScanner) (*PositionedScanner, bool) {
	for scanner.Scan() {
		if parser.isEntityTitle(scanner.Text()) {
			return scanner, false
		}
	}

	return nil, true
}

// getKeyAddrOffset returns -1 if the entity is not found
func (parser LdifParser) getKeyAttrOffset(file *os.File, keyAttr EntityAttribute) (int64, error) {
	keyAttrStr := strings.ToLower(keyAttr.Stringify()[0])
	Logger.Info("searching with key: \"%s\"", keyAttrStr)

	scanner := NewPositionedScanner(file)

	for scanner.Scan() {
		attrLine := strings.ToLower(scanner.Text())
		if keyAttrStr == attrLine {
			return scanner.Position(), nil
		}
		Logger.Debug("attrLine \"%s\" does not match key", attrLine)
	}

	// Entity not found
	return -1, scanner.scanner.Err()
}

func (parser LdifParser) getPrevEntityOffset(file *os.File, lineOffset int64) (int, error) {
	scanner := backscanner.New(file, int(lineOffset))
	for {
		line, pos, err := scanner.Line()
		if err != nil {
			return pos, err
		}

		if parser.isEntityTitle(line) {
			return pos, nil
		}
	}
}

// BuildEntity returns an empty Entity object if the object is not found,
// other wise it returns the entity object or an error if one is encountered.
func (parser LdifParser) BuildEntity(keyAttrName string, keyAttrVal string) (Entity, error) {
	Logger.Info("Opening ldif file: " + parser.filename)
	dumpFile, err := os.Open(parser.filename)
	if err != nil {
		return Entity{}, err
	}
	defer dumpFile.Close()

	keyAttr := NewEntityAttribute(keyAttrName, keyAttrVal)

	keyAttrOffset, err := parser.getKeyAttrOffset(dumpFile, keyAttr)
	if err != nil {
		return Entity{}, err
	}
	Logger.Debug("Key found at position: %d", keyAttrOffset)

	// Object not found
	if keyAttrOffset == -1 {
		return Entity{}, nil
	}

	entityOffset, err := parser.getPrevEntityOffset(dumpFile, keyAttrOffset)
	if err != nil {
		return Entity{}, err
	}
	Logger.Info("Entity found at position: %d", entityOffset)

	_, err = dumpFile.Seek(int64(entityOffset), 0)
	if err != nil {
		return Entity{}, err
	}

	entityScanner := bufio.NewScanner(dumpFile)
	buf := make([]byte, parser.scannerBufferSize)
	entityScanner.Buffer(buf, parser.scannerBufferSize)

	Logger.Info("Parsing entity from block")
	entity := parser.getEntityFromBlock(entityScanner)
	return entity, nil
}

// BuildEntities constructs an ldap entity per entry in the input ldif file.
func (parser LdifParser) BuildEntities() ([]Entity, error) {
	entities := []Entity{}

	Logger.Info("Opening ldif file: " + parser.filename)
	dumpFile, err := os.Open(parser.filename)
	if err != nil {
		return nil, err
	}
	defer dumpFile.Close()

	Logger.Info("Finding first entity block")
	entityScanner := parser.findFirstEntityBlock(dumpFile)
	if entityScanner == nil {
		return entities, errors.New("Unable to find first entity block")
	}

	for entityScanner.Scan() {
		titleLine := entityScanner.Text()
		if !parser.isEntityTitle(titleLine) {
			continue
		}

		Logger.Info("Parsing entity")
		entity := parser.getEntityFromBlock(entityScanner)

		dn, dnFound := entity.GetDN()
		if !dnFound {
			Logger.Error("Unable to parse DN for entity: " + titleLine)
			continue
		}

		if parser.entityFilter != nil && len(parser.entityFilter) > 0 {
			Logger.Info("Applying entity filter to: " + dn.Value.GetSingleValue())
			if !MatchesFilter(entity, parser.entityFilter) {
				continue
			}
		}

		Logger.Info("Appending matched entity: " + dn.Value.GetSingleValue())
		entities = append(entities, entity)
	}

	return entities, nil
}

// CountEntities returns the number of ldap entites in the input ldif file.
func (parser LdifParser) CountEntities() (int, error) {
	filterParts := []string{"dn", "distinguishedName"}
	for _, filter := range parser.entityFilter {
		entityFilter := filter.(EntityFilter)
		filterParts = append(filterParts, entityFilter.AttributeName)
	}

	attrFilter := BuildAttributeFilter(filterParts)
	parser.SetAttributeFilter(attrFilter)

	entities, err := parser.BuildEntities()
	if err != nil {
		return 0, err
	}

	return len(entities), nil
}

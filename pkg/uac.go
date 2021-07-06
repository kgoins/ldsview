package ldsview

import (
	"fmt"
	"io"
	"sort"
	"strconv"
	"text/tabwriter"

	"github.com/audibleblink/msldapuac"
	"github.com/kgoins/ldapentity/entity"
)

type UACFilter struct {
	searchProp int
}

func NewUACFilter(searchTerm int) UACFilter {
	return UACFilter{searchProp: searchTerm}
}

func (filter UACFilter) Matches(entity entity.Entity) (isMatch bool) {
	uacAttr, found := entity.GetAttribute("userAccountControl")
	if found {
		uacString := uacAttr.Value.GetSingleValue()
		uac, _ := strconv.ParseInt(uacString, 10, 32)
		isMatch, err := msldapuac.IsSet(uac, filter.searchProp)
		if err != nil {
			fmt.Printf("Unable to verify UAC: %v\n", err.Error())
			return false
		}
		return isMatch
	}
	return
}

// GetFlagsFromUAC wraps the msldapuac dependency for easy re-definition
func GetFlagsFromUAC(uac int64) ([]string, error) {
	return msldapuac.ParseUAC(uac)
}

// UACSearch will seek for entities who have the given UAC property set
func UACSearch(entities *[]entity.Entity, uacProp int) (matches []entity.Entity) {
	for _, entity := range *entities {
		uacStr, found := entity.GetSingleValuedAttribute("userAccountControl")
		if !found {
			continue
		}

		i64, _ := strconv.ParseInt(uacStr, 10, 32)
		isMatch, err := msldapuac.IsSet(i64, uacProp)
		if err != nil {
			continue
		}
		if isMatch {
			matches = append(matches, entity)
		}
	}
	return
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

//UACParse take a UAC int and return
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

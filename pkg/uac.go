package ldsview

import (
	"strconv"

	"github.com/audibleblink/msldapuac"
)

// GetFlagsFromUAC wraps the msldapuac dependency for easy re-definition
func GetFlagsFromUAC(uac int64) ([]string, error) {
	return msldapuac.ParseUAC(uac)
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

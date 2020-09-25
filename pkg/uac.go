package ldsview

import (
	"github.com/audibleblink/msldapuac"
)

func GetFlagsFromUAC(uac int64) ([]string, error) {
	return msldapuac.ParseUAC(uac)
}

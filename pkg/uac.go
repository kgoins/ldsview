package ldsview

import (
	"fmt"
	"reflect"
	"strings"
)

type UACFlags struct {
	Enabled              bool
	LockedOut            bool
	NormalAccount        bool
	MustChangePassword   bool
	SmartcardRequired    bool
	PasswordNeverExpires bool
	TrustedForDelegation bool
	PasswordNotRequired  bool
	PreAuthNotRequired   bool
}

func GetFlagsFromUAC(uac int64) UACFlags {
	return UACFlags{
		Enabled:              (uac & 2) == 0,
		LockedOut:            (uac & 16) != 0,
		NormalAccount:        (uac & 512) != 0,
		SmartcardRequired:    (uac & 262144) != 0,
		PasswordNeverExpires: (uac & 65536) != 0,
		TrustedForDelegation: (uac & 16777216) != 0,
		PasswordNotRequired:  (uac & 32) != 0,
		PreAuthNotRequired:   (uac & 4194304) != 0,
	}
}

func (uac UACFlags) String() string {
	reflectedUAC := reflect.ValueOf(&uac).Elem()
	typeofUAC := reflectedUAC.Type()

	builder := strings.Builder{}
	for i := 0; i < reflectedUAC.NumField(); i++ {
		field := reflectedUAC.Field(i)
		if field.Bool() {
			fmt.Fprintln(&builder, typeofUAC.Field(i).Name)
		}
	}

	return builder.String()
}

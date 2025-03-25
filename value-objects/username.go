package valueObjects

import (
	"strings"
)

type Username struct {
	value string
	valid bool
}

func NewUsername(value string) *Username {
	formattedValue := strings.ToLower(strings.Trim(value, " "))
	return &Username{
		value: formattedValue,
		valid: strings.HasPrefix(formattedValue, "@"),
	}
}

func (u *Username) String() string {
	return u.value
}

func (u *Username) Equals(other Username) bool {
	return u.value == other.value
}

func (u *Username) IsValid() bool {
	return u.valid
}

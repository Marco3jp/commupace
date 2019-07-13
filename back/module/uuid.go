package module

import (
	"strings"
	"github.com/satori/go.uuid"
)

func CreateUUIDWithoutHyphen() (_uuid string, err error) {
	rawToken, err := uuid.NewV4()
	if err != nil {
		return "", &UUIDCreateError{err.Error()}
	}

	return strings.Replace(rawToken.String(), "-", "", -1), nil
}
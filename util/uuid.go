package utils

import (
	"github.com/google/uuid"
	"strings"
)

func GetId() string{
	return strings.Replace(uuid.New().String(), "-", "", -1)
}

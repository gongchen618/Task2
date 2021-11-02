package util

import (
	"github.com/google/uuid"
)

func GetRandom() (string) {
	return uuid.New().String()
}
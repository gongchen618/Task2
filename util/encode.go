package util

import (
	"crypto/sha256"
	"fmt"
)

func Encode (str string) (string) {
	return fmt.Sprint("%x", sha256.Sum256([]byte(str)))
}


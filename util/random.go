package util

import (
	"math/rand"
	"time"
)

func init () {
	rand.Seed(time.Now().Unix())
}

func GetRandom() (int) {
	return rand.Intn(100)
}
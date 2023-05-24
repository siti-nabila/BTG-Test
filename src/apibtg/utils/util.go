package utils

import (
	"math/rand"
	"time"
)

func GenerateID() int {
	min := 1
	max := 999
	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(max-min) + min

	return id
}

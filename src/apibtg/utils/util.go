package utils

import (
	"math/rand"
	"time"
)

func GenerateID() int {
	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(5)
	return id
}

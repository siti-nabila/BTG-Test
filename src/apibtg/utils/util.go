package utils

import (
	"math/rand"
	"strings"
	"time"
)

func GenerateID() int {
	min := 1
	max := 999
	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(max-min) + min

	return id
}

func SplitDateTime(data string) (string, string) {
	date := strings.Split(data, "T")
	time := strings.Split(date[1], "Z")
	return date[0], time[0]
}

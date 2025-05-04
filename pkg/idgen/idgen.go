package idgen

import (
	"math/rand"
	"strconv"
	"time"
)

func GenerateID() string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(100000 + rand.Intn(900000))
}

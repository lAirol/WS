package random

import (
	"math/rand"
	"time"
)

func GenerateUniqUserId() int {
	baseID := int(time.Now().Unix())
	randomSuffix := rand.Intn(100) + 1
	id := baseID*100 + randomSuffix
	return id
}

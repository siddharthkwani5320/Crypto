package random

import (
	"math/rand"
	"time"
)

//Generate random id no.
func RandomId() string {
	rand.Seed(time.Now().UnixNano())
	return randomIntegerString(4)
}

//Generate random Account no.
func RandomAcc() string {
	rand.Seed(time.Now().UnixNano())
	return randomIntegerString(12)
}

func randomIntegerString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(48, 57))
	}
	return string(bytes)
}

func randInt(min int, max int) int {

	return min + rand.Intn(max-min)
}

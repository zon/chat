package core

import (
	"crypto/rand"
	"encoding/hex"
	"math"
)

func RandomString(l int) string {
	buff := make([]byte, int(math.Ceil(float64(l)/2)))
	rand.Read(buff)
	str := hex.EncodeToString(buff)
	return str[:l]
}

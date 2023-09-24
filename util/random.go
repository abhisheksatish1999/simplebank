package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabets = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(l int) string {
	k := len(alphabets)
	var sb strings.Builder
	for i := 0; i < l; i++ {
		c := alphabets[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	curr := []string{"USD", "CAD", "EUR"}
	l := len(curr)
	return curr[rand.Intn(l)]
}

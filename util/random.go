package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min * rand.Int63n(max-min+1)
}

// RandomString generate a random string of length
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generate a random Owner
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generate random amount of money
func RandomMoney() int64 {
	return RandomInt(100, 1000)
}

// RandomCurrency generate random currency
func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

// RandomEmail generate random email
func RandomEmail() string{
	return fmt.Sprintf("%s@email.com",RandomString(8))
}
package service

import (
	"crypto/rand"
	"math/big"
)

// Returns random []byte with length (size)
func RandBytes(size int) []byte {
	bytes := make([]byte, size)
	rand.Read(bytes)
	return bytes
}

// Returns random 26 byte string
func RandString() string {
	return rand.Text()
}

// Returns random integer [0, max)
func RandInt(max int) int {
	n, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
	return int(n.Int64())
}

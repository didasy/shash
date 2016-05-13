// Package shash will generate a random, base-62 string with user defined length without nonsense.
// Usage Example:
//  func main() {
// 	 h := shash.New(10)
// 	 gen, _ := h.Generate()
// 	 fmt.Println(gen)
//  }
package shash

import (
	"crypto/rand"
	"math/big"
)

// The base62 charset as const CHARSET.
const (
	CHARSET = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	VERSION = "0.0.2"
)

var MAX_LEN = int64(len(CHARSET) - 1)

// Hash only has one property: Length, with default value of 5 on creation.
type Hash struct {
	Length int
}

// Create a new Hash with default length 5.
func New(n int) *Hash {
	if n < 5 {
		n = 5
	}
	return &Hash{
		Length: n,
	}
}

// Generate() will return a random string in base62. Error will return if there is an error with the crypto/rand package, nil if there isn't any error.
func (h *Hash) Generate() (string, error) {
	var gen = make([]rune, h.Length)
	l := big.NewInt(MAX_LEN)
	for i := 0; i < h.Length; i++ {
		r, err := rand.Int(rand.Reader, l)
		if err != nil {
			return "", err
		}
		gen[i] = rune(CHARSET[r.Uint64()])
	}
	return string(gen), nil
}

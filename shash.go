// Package shash will generate a random, base-62 string with user defined length without nonsense.
// Usage Example:
//  func main() {
// 	 h := shash.New()
// 	 h.SetLength(5)
// 	 gen, _ := h.Generate()
// 	 fmt.Println(gen)
//  }
package shash

import (
	"crypto/rand"
	"errors"
	"math/big"
	"strings"
)

// The base62 charset as const CHARSET.
const (
	CHARSET = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

// Hash only has one property: Length, with default value of 5 on creation.
type Hash struct {
	Length int
}

// Create a new Hash with default length 5.
func New() *Hash {
	return &Hash{5}
}

// Set Hash length and return error if length is less than 5, nil if there isn't any error.
func (h *Hash) SetLength(i int) error {
	if i < 5 {
		return errors.New("Length must be 5 or more")
	}
	h.Length = i
	return nil
}

// Generate() will return a random string in base62. Error will return if there is an error with the crypto/rand package, nil if there isn't any error.
func (h *Hash) Generate() (string, error) {
	gen := make([]string, h.Length)
	l := len(CHARSET) - 1
	for i := 0; i < h.Length; i++ {
		r, err := rand.Int(rand.Reader, big.NewInt(int64(l)))
		if err != nil {
			return "", err
		}
		x := string(CHARSET[r.Uint64()])
		gen = append(gen, x)
	}
	return strings.Join(gen, ""), nil
}
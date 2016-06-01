// Package shash will generate a random, base-62, base-36, and base-10 string with user defined length without nonsense.
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
	CHARSET         = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	CHARSET_BASE_36 = "abcdefghijklmnopqrstuvwxyz0123456789"
	CHARSET_BASE_10 = "0123456789"
	VERSION         = "0.1.0"
)

var (
	MAX_LEN         = int64(len(CHARSET) - 1)
	MAX_LEN_BASE_36 = int64(len(CHARSET_BASE_36) - 1)
	MAX_LEN_BASE_10 = int64(len(CHARSET_BASE_10) - 1)
)

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
	return generate(CHARSET, MAX_LEN, h.Length)
}

func (h *Hash) GenerateBase10() (string, error) {
	return generate(CHARSET_BASE_10, MAX_LEN_BASE_10, h.Length)
}

func (h *Hash) GenerateBase36() (string, error) {
	return generate(CHARSET_BASE_36, MAX_LEN_BASE_36, h.Length)
}

func generate(charset string, maxLen int64, hasherLength int) (string, error) {
	var gen = make([]rune, hasherLength)
	l := big.NewInt(maxLen)
	for i := 0; i < hasherLength; i++ {
		r, err := rand.Int(rand.Reader, l)
		if err != nil {
			return "", err
		}
		gen[i] = rune(charset[r.Uint64()])
	}
	return string(gen), nil
}

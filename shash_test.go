package shash_test

import (
	. "github.com/JesusIslam/shash"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

var _ = Describe("Shash", func() {

	var hash *Hash
	var str string
	var err error
	var inCharset bool

	Describe("Testing hash", func() {
		Context("Create new randomizer with provided length", func() {
			It("With given length of < 5, length should be 5", func() {
				hash = New(0)
				Expect(hash.Length).To(Equal(5))
			})
			It("With given length of 10, length must be 10", func() {
				hash = New(10)
				Expect(hash.Length).To(Equal(10))
			})
		})

		Context("Generate base62 hash", func() {
			It("Should return a string with length of 10 without returning an error", func() {
				str, err = hash.Generate()
				Expect(err).To(BeNil())
				Expect(str).NotTo(BeEmpty())
				Expect(str).To(HaveLen(10))
			})

			It("Should return only characters in CHARSET", func() {
				str, _ = hash.Generate()
				inCharset = isInCharset(str, CHARSET)
				Expect(inCharset).To(BeTrue())
			})
		})

		Context("Generate base36 hash", func() {
			It("Should return a string with length of 10 without returning an error", func() {
				str, err = hash.GenerateBase36()
				Expect(err).To(BeNil())
				Expect(str).NotTo(BeEmpty())
				Expect(str).To(HaveLen(10))
			})

			It("Should return only characters in CHARSET_BASE_36", func() {
				str, _ = hash.GenerateBase36()
				inCharset = isInCharset(str, CHARSET_BASE_36)
				Expect(inCharset).To(BeTrue())
			})
		})

		Context("Generate base10 hash", func() {
			It("Should return a string with length of 10 without returning an error", func() {
				str, err = hash.GenerateBase10()
				Expect(err).To(BeNil())
				Expect(str).NotTo(BeEmpty())
				Expect(str).To(HaveLen(10))
			})

			It("Should return only characters in CHARSET_BASE_10", func() {
				str, _ = hash.GenerateBase10()
				inCharset = isInCharset(str, CHARSET_BASE_10)
				Expect(inCharset).To(BeTrue())
			})
		})
	})
})

func isInCharset(str string, charset string) bool {
	for _, s := range str {
		found := false
		for _, c := range charset {
			if s == c {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	return true
}

func BenchmarkGenerate5Base62(b *testing.B) {
	h := New(5)
	for n := 0; n < b.N; n++ {
		h.Generate()
	}
}

func BenchmarkGenerate16Base62(b *testing.B) {
	h := New(16)
	for n := 0; n < b.N; n++ {
		h.Generate()
	}
}

func BenchmarkGenerate32Base62(b *testing.B) {
	h := New(32)
	for n := 0; n < b.N; n++ {
		h.Generate()
	}
}

func BenchmarkGenerate5Base36(b *testing.B) {
	h := New(5)
	for n := 0; n < b.N; n++ {
		h.GenerateBase36()
	}
}

func BenchmarkGenerate16Base36(b *testing.B) {
	h := New(16)
	for n := 0; n < b.N; n++ {
		h.GenerateBase36()
	}
}

func BenchmarkGenerate32Base36(b *testing.B) {
	h := New(32)
	for n := 0; n < b.N; n++ {
		h.GenerateBase36()
	}
}

func BenchmarkGenerate5Base10(b *testing.B) {
	h := New(5)
	for n := 0; n < b.N; n++ {
		h.GenerateBase10()
	}
}

func BenchmarkGenerate16Base10(b *testing.B) {
	h := New(16)
	for n := 0; n < b.N; n++ {
		h.GenerateBase10()
	}
}

func BenchmarkGenerate32Base10(b *testing.B) {
	h := New(32)
	for n := 0; n < b.N; n++ {
		h.GenerateBase10()
	}
}

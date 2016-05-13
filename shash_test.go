package shash_test

import (
	. "github.com/JesusIslam/shash"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

var _ = Describe("Shash", func() {

	var hash *Hash

	Describe("Testing hash", func() {
		Context("Create new with length length", func() {
			It("With hash length of < 5, length should be 5", func() {
				hash = New(0)
				Expect(hash.Length).To(Equal(5))
			})
			It("With hash length of 10, length must be 10", func() {
				hash = New(10)
				Expect(hash.Length).To(Equal(10))
			})
		})

		Context("Generate hash", func() {
			It("Should return a string with length of 10 without returning an error", func() {
				str, err := hash.Generate()
				Expect(err).To(BeNil())
				Expect(str).NotTo(BeEmpty())
				Expect(str).To(HaveLen(10))
			})
		})
	})
})

func BenchmarkGenerate(b *testing.B) {
	h := New(5)
	for n := 0; n < b.N; n++ {
		h.Generate()
	}
}
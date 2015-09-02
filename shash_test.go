package shash_test

import (
	. "github.com/JesusIslam/shash"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Shash", func() {

	var hash *Hash

	BeforeEach(func() {
		hash = New()
	})

	Describe("Testing hash", func() {
		Context("Set hash length", func() {
			It("With hash length of < 5, should return an error", func() {
				err := hash.SetLength(0)
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).NotTo(BeEmpty())
			})
			It("With hash length of > 5, should not return an error", func() {
				err := hash.SetLength(10)
				Expect(err).To(BeNil())
			})
		})

		Context("Generate hash", func() {
			It("Should return a string with length of 5 without returning an error", func() {
				str, err := hash.Generate()
				Expect(err).To(BeNil())
				Expect(str).NotTo(BeEmpty())
				Expect(str).To(HaveLen(5))
			})
		})
	})
})

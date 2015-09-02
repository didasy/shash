package shash_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestShash(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Shash Suite")
}

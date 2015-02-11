package main 

import (
	"testing"
	"github.com/JesusIslam/shash"
)

var result string

func TestGenerate(t *testing.T) {
	h := shash.New()
	r, err := h.Generate()
	if err != nil {
		t.Fatal(err)
	}
	if len(r) != 5 {
		t.Fatal("Generated string length is not 5")
	}
}

func BenchmarkGenerate(b *testing.B) {
	h := shash.New()
	var r string
	for n := 0; n < b.N; n++ {
		r, _ = h.Generate()
	}
	result = r
}
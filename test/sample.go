package main

import (
	"fmt"
	"github.com/JesusIslam/shash"
)

func main() {
	h := shash.New()
	r, _ := h.Generate()
	fmt.Println(r)
}
![Build Status](https://travis-ci.org/JesusIslam/shash.svg)
[![GoDoc](https://godoc.org/github.com/JesusIslam/shash?status.svg)](https://godoc.org/github.com/JesusIslam/shash)

# Shash
### A no-nonsense golang package for creating random string

This package is very simple, just do

```
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
```

The example will generate a base62 random string with length of 5 (default).

For the documentation see [here](http://godoc.org/github.com/JesusIslam/shash).

To test, just run `go test`, but you need to have [gomega](http://github.com/onsi/gomega) and [ginkgo](http://github.com/onsi/ginkgo) installed.

To install `go get github.com/JesusIslam/shash`.

### Benchmark
On i3-3217U @1.8GHz with `go test -bench . -benchtime=5s -cpu 4`

`BenchmarkGenerate-4       300000             29516 ns/op`

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

To test, just run `go test` in the test directory.

To install `go get github.com/JesusIslam/shash`.
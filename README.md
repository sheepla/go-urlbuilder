<div align="center">

# 🔗 urlbuilder

</div>

**urlbuilder** urlbuilder is a Go module based on `net/url` standard module, aimed at safely constructing URL strings with a concise syntax.

## Why?

It is a good idea to use the `net/url` standard module when safely constructing URL strings. 
However, if you use `net/url` as is, you often need to prepare temporary variables, and have to write non-declaretive
which I felt was a bit cumbersome when building complex URLs over and over again.

This module was created as a concise and easy way to construct URL strings based on `net/url`.

## Usage

```go
package main

import (
	"fmt"

	"github.com/sheepla/urlbuilder"
)

var sourceURL = "https://localhost:8080/path/to/resource#helloworld?key1=value1&key2=value2"

func main() {
	ExampleURLBuilder()
}

func ExampleURLBuilder() {
	u, err := urlbuilder.Parse(sourceURL)
	if err != nil {
		panic(err)
	}

	u.SetScheme("ftp").
		SetHost("another.example.com:12345").
		SetFragument("anotherFragument").
		SetPath("/", "日本語", "with space").
		AddQuery("key3", "key3").
		RemoveQuery("key2")

	// Will output:
	// ftp://another.example.com:12345/%25E6%2597%25A5%25E6%259C%25AC%25E8%25AA%259E/with%2520space?key3=key3#anotherFragument
	fmt.Println(u.MustString())
}
```

## License

MIT

## Author

[sheepla](https://github.com/sheepla)



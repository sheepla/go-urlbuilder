<div align="center">

# 🔗 urlbuilder

</div>

**urlbuilder** is a Go module to build or edit URL string based on `net/url` standard module.

## Usage

```go
package main

import (
	"fmt"

	"github.com/sheepla/go-urlbuilder"
)

var exampleURL = "https://localhost:8080/path/to/resource#helloworld?key1=value1&key2=value2"

func main() {
    u, err := urlbuilder.Parse(exampleURL)
    if err != nil {
        panic(err)
    }

	u.SetScheme("ftp").
		SetHost("another.example.com:12345").
		SetFragument("anotherFragument").
		EditPath(func(elements []string) []string {
			sl := []string{"日本語", "with space"}
			elements = append(elements, sl...)

			return elements
		}).
		AddQuery("key3", "key3").
		RemoveQuery("key2")

    // Will output:
    // ftp://another.example.com:12345/path/to/resource/%25E6%2597%25A5%25E6%259C%25AC%25E8%25AA%259E/with%2520space?key3=key3#anotherFragument
    fmt.Println(u.MustString())
}
```

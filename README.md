<div align="center">

# 🔗 go-urlbuilder

</div>

**go-urlbuilder** is a Go module based on `net/url` standard module, aimed at safely constructing URL strings with a concise syntax.

## Why?

It is a good idea to use the `net/url` standard module when safely constructing URL strings. 
However, if you use `net/url` as is, you often need to prepare temporary variables, and have to write non-declaretive code which I felt was a bit cumbersome when building complex URLs over and over again.

This module was created as a concise and easy way to construct URL strings based on `net/url`.

## Usage

```go
package main

import (
	"fmt"
	"net/url"

	"github.com/sheepla/go-urlbuilder"
)

var sourceURL = "https://localhost:8080/path/to/resource#helloworld?key1=value1&key2=value2"

func main() {
	u, err := urlbuilder.Parse(sourceURL)
	if err != nil {
		panic(err)
	}

	u.SetScheme("http").
		SetHost("example.com:12345").
		SetFragument("anotherFragument").
		EditPath(func(elements []string) []string {
			return append(elements, "Go言語")
		}).
		EditQuery(func(q url.Values) url.Values {
			q.Set("key1", "key1-edited")
			q.Del("key2")
			q.Add("key3", "value3")

			return q
		})

		// => http://example.com:12345/path/to/resource/Go%25E8%25A8%2580%25E8%25AA%259E?key1=key1-edited&key3=value3#anotherFragument
	fmt.Println(u.MustString())
}
```

## License

MIT

## Author

[sheepla](https://github.com/sheepla)



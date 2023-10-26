package main

import (
	"fmt"
	"net/url"

	"github.com/sheepla/go-urlbuilder"
)

var sourceURL = "https://localhost:8080/path/to/resource#helloworld?key1=value1&key2=value2"

func main() {
	u := urlbuilder.MustParse(sourceURL)

	u.SetScheme("http").
		SetHost("example.com:12345").
		SetFragment("anotherFragment").
		EditPath(func(elements []string) []string {
			return append(elements, "Go言語")
		}).
		EditQuery(func(q url.Values) url.Values {
			q.Set("key1", "key1-edited")
			q.Del("key2")
			q.Add("key3", "value3")

			return q
		})

	// => http://example.com:12345/path/to/resource/Go%25E8%25A8%2580%25E8%25AA%259E?key1=key1-edited&key3=value3#anotherFragment
	fmt.Println(u.MustString())
}

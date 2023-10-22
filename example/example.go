package main

import (
	"fmt"

	"github.com/sheepla/go-urlbuilder"
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

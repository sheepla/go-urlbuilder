package urlbuilder_test

import "testing"
import "github.com/sheepla/urlbuilder"

var exampleURL = "https://localhost:8080/path/to/resource#helloworld?key1=value1&key2=value2"

func Test(t *testing.T) {
    u, err := urlbuilder.Parse(exampleURL)
    if err != nil {
        t.Fatal(err)
    }

    u.SetScheme("ftp").
        SetHost("another.example.com:12345").
        SetFragument("anotherFragument").
        EditPath(func(elements []string)[]string {
            sl := []string{"日本語", "with space"}
            elements = append(elements, sl...)

            return elements
        }).
        AddQuery("key3", "key3").
        RemoveQuery("key2")

    have, err := u.String()
    if err != nil {
        t.Fatal(err)
    }

    want := "ftp://another.example.com:12345/path/to/resource/%25E6%2597%25A5%25E6%259C%25AC%25E8%25AA%259E/with%2520space?key3=key3#anotherFragument"
    if have != want {
        t.Fatalf("have=%s\nwant=%s\n", have, want)
    }
}

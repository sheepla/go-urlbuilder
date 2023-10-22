package urlbuilder_test

import (
	"testing"

	"net/url"

	"github.com/sheepla/go-urlbuilder"
)

func TestBasic(t *testing.T) {
    sourceURL := "https://localhost:8080/path/to/resource?key1=value1&key2=value2#helloWorld"
	u, err := urlbuilder.Parse(sourceURL)
	if err != nil {
		t.Fatal(err)
	}

	u.SetScheme("http").
		SetHost("another.example.com:12345").
		SetFragument("anotherFragument")

	have, err := u.String()
	if err != nil {
		t.Fatalf("an error occurred on constructing URL: %s\n", err)
	}

	want := "http://another.example.com:12345/path/to/resource?key1=value1&key2=value2#anotherFragument"
	if have != want {
		t.Fatalf("have=%s\nwant=%s\n", have, want)
	}
}

func TestPathEditing(t *testing.T) {

}

func TestQueryEditing(t *testing.T) {
    sourceURL := "https://localhost:8080/path/to/resource?key1=value1&key2=value2"
	u, err := urlbuilder.Parse(sourceURL)
	if err != nil {
		t.Fatal(err)
	}

    u.EditQuery(func(q url.Values) url.Values {
        t.Logf("current query: %s", q.Encode())

        q.Set("key1", "value1-edited")
        q.Del("key2")
        q.Add("key3", "value3")

        t.Logf("key1: %s", q.Get("key1"))
        t.Logf("key2: %s", q.Get("key2"))
        t.Logf("key3: %s", q.Get("key3"))
        t.Logf("edited query: %s", q.Encode())

        return q
    })

	have, err := u.String()
	if err != nil {
		t.Fatalf("an error occurred on constructing URL: %s\n", err)
	}

	want := "https://localhost:8080/path/to/resource?key1=value1-edited&key3=value3"

	if have != want {
		t.Fatalf("have=%s\nwant=%s\n", have, want)
	}
}

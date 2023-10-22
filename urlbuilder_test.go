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
		SetFragument("anotherFragument").
        SetUserWithPassword("u$er1", "P@$$w0rd!")

	have, err := u.String()
	if err != nil {
		t.Fatalf("an error occurred on constructing URL: %s\n", err)
	}

	want := "http://u$er1:P%40$$w0rd%21@another.example.com:12345/path/to/resource?key1=value1&key2=value2#anotherFragument"
	if have != want {
		t.Fatalf("have=%s\nwant=%s\n", have, want)
	}
}


func TestPathEditing(t *testing.T) {
    sourceURL := "https://localhost:8080/あ/progr@mm!ng"
	u, err := urlbuilder.Parse(sourceURL)
	if err != nil {
		t.Fatal(err)
	}

    u.EditPath(func(elements []string) []string {
        t.Log("current elements: ", elements)
        elements = append(elements, "Go言語")
        t.Log("edited elements: ", elements)

        return elements
    })

    have, err := u.String()
    if err != nil {
        t.Fatal(err)
    }

    want := "https://localhost:8080/%25E3%2581%2582/progr@mm%2521ng/Go%25E8%25A8%2580%25E8%25AA%259E"
    if have != want {
        t.Fatalf("have=%s\nwant=%s\n", have, want)
    }
    
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

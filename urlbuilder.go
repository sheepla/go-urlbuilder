package urlbuilder

import (
	"net/url"
	"strings"
)

type URL struct {
	internal *url.URL
	err      error
}

func (u *URL) Error() string {
	return u.err.Error()
}

func Parse(s string) (*URL, error) {
	netURL, err := url.Parse(s)
	if err != nil {
		return &URL{
			internal: netURL,
			err:      err,
		}, err
	}

	return &URL{
		internal: netURL,
		err:      nil,
	}, nil
}

func (u *URL) SetPath(base string, elements ...string) *URL {
	path, err := url.JoinPath(base, elements...)
	u.err = err
	u.internal.Path = path

	return u
}

func (u *URL) EditPath(editFunc func([]string) []string) *URL {
	elements := strings.Split(u.internal.Path, "/")
	elements = editFunc(elements)
	path, err := url.JoinPath("/", elements...)
	u.err = err
	u.internal.Path = path

	return u
}

func (u *URL) SetScheme(scheme string) *URL {
	u.internal.Scheme = scheme

	return u
}

func (u *URL) SetHost(host string) *URL {
	u.internal.Host = host

	return u
}

func (u *URL) SetFragument(fragument string) *URL {
	u.internal.Fragment = fragument

	return u
}

func (u *URL) SetQuery(key, value string) *URL {
	q, err := url.ParseQuery(u.internal.RawQuery)
	u.err = err
	q.Set(key, value)
	u.internal.RawQuery = q.Encode()

	return u
}

func (u *URL) AddQuery(key, value string) *URL {
	q, err := url.ParseQuery(u.internal.RawQuery)
	u.err = err
	q.Add(key, value)
	u.internal.RawQuery = q.Encode()

	return u
}

func (u *URL) RemoveQuery(key string) *URL {
	q, err := url.ParseQuery(u.internal.RawQuery)
	u.err = err
	q.Del(key)
	u.internal.RawQuery = q.Encode()

	return u
}

func (u *URL) String() (string, error) {
	return u.internal.String(), u.err
}

func (u *URL) MustString() string {
	s, err := u.String()
	if err != nil {
		panic(err)
	}

	return s
}

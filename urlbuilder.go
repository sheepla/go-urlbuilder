package urlbuilder

import (
	"net/url"
	"strings"
)

type URL struct {
	internal *url.URL
	Err      error
}

func (u *URL) Error() string {
	return u.Err.Error()
}

func Parse(s string) *URL {
	netURL, err := url.Parse(s)
	return &URL{
		internal: netURL,
		Err:      err,
	}
}

func (u *URL) SetPath(base string, elements ...string) *URL {
	path, err := url.JoinPath(base, elements...)
	u.Err = err
	u.internal.Path = path

	return u
}

func (u *URL) EditPath(editFunc func([]string) []string) *URL {
	elements := strings.Split(u.internal.Path, "/")

	// To prevent double escaping,
	// each path element is unescaped before being passed to the editing function.
	for i := 0; i < len(elements); i++ {
		if escaped, err := url.PathUnescape(elements[i]); err == nil {
			elements[i] = escaped
		}
	}

	elements = editFunc(elements)

	path, err := url.JoinPath("/", elements...)
	u.Err = err
	u.internal.Path = path

	return u
}

func (u *URL) AppendPath(elements ...string) *URL {
	u.EditPath(func(current []string) []string {
		return append(current, elements...)
	})

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

func (u *URL) SetUser(userName string) *URL {
	u.internal.User = url.User(userName)

	return u
}

func (u *URL) SetUserWithPassword(userName, password string) *URL {
	u.internal.User = url.UserPassword(userName, password)

	return u
}

func (u *URL) SetFragument(fragument string) *URL {
	u.internal.Fragment = fragument

	return u
}

func (u *URL) EditQuery(editFunc func(url.Values) url.Values) *URL {
	edited := editFunc(u.internal.Query())
	u.internal.RawQuery = edited.Encode()

	return u
}

func (u *URL) SetQuery(key, value string) *URL {
	u.EditQuery(func(q url.Values) url.Values {
		q.Set(key, value)
		return q
	})

	return u
}

func (u *URL) AddQuery(key, value string) *URL {
	u.EditQuery(func(q url.Values) url.Values {
		q.Add(key, value)
		return q
	})

	return u
}

func (u *URL) RemoveQuery(key string) *URL {
	u.EditQuery(func(q url.Values) url.Values {
		q.Del(key)
		return q
	})

	return u
}

func (u *URL) String() (string, error) {
	return u.internal.String(), u.Err
}

func (u *URL) MustString() string {
	s, err := u.String()
	if err != nil {
		panic(err)
	}

	return s
}

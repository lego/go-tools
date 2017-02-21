package pkg

import "net/url"

func fn() {
	_, _ = url.Parse("foobar")
	_, _ = url.Parse(":") // MATCH /is not a valid URL/
	_, _ = url.Parse("https://golang.org")
}

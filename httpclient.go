package config1

import (
	"github.com/sgtest/config0/http"
)

func Foo(url string) {
	http.RespondsWithHTTPOK("https://sourcegraph.com")
	http.HTTPGet("https://sourcegraph.com")
}

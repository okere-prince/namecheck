package stub

import (
	"io"
	"net/http"
	"strings"

	"github.com/jub0bs/namecheck"
)

type clientFunc func(url string) (*http.Response, error)

func (f clientFunc) Get(url string) (*http.Response, error) {
	return f(url)
}

func ClientWithError(err error) namecheck.Client {
	get := func(_ string) (*http.Response, error) {
		return nil, err
	}
	return clientFunc(get)
}

func ClientWithStatusCodeAndBody(sc int, body string) namecheck.Client {
	get := func(_ string) (*http.Response, error) {
		res := http.Response{
			StatusCode: sc,
			Body:       io.NopCloser(strings.NewReader(body)),
		}
		return &res, nil
	}
	return clientFunc(get)
}

func ClientWithStatusCode(sc int) namecheck.Client {
	return ClientWithStatusCodeAndBody(sc, "")
}

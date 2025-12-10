package v1

import (
	"io"
	"net/http"
	"strings"
)

const (
	invalidJSONBody = `{
			"result": [
				invalid
			]
		}`
	httpErrorBody    = "Not Found"
	httpErrorMessage = "got the 404 status code from the server: Not Found"
)

// newFakeClient creates a new ServiceClient with the given endpoint and transport.
func newFakeClient(endpoint string, transport http.RoundTripper) *ServiceClient {
	return &ServiceClient{
		HTTPClient: &http.Client{Transport: transport},
		APIUrl:     endpoint,
	}
}

// RoundTripFunc lets us use a function as an http.RoundTripper.
type RoundTripFunc func(req *http.Request) (*http.Response, error)

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

// NewFakeResponse creates a fake *http.Response with the provided status and body.
func NewFakeResponse(status int, body string) *http.Response {
	return &http.Response{
		StatusCode: status,
		Body:       io.NopCloser(strings.NewReader(body)),
	}
}

// NewFakeTransport returns a fake transport with the given response and error.
func NewFakeTransport(resp *http.Response, err error) RoundTripFunc {
	return RoundTripFunc(func(_ *http.Request) (*http.Response, error) {
		return resp, err
	})
}

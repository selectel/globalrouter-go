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
func newFakeClient(endpoint string, transport *NewFakeTransport) *ServiceClient {
	return &ServiceClient{
		HTTPClient: &http.Client{Transport: transport},
		APIUrl:     endpoint,
	}
}

// fakeTransport lets us check body after decoding json and also process response and error.
type NewFakeTransport struct {
	resp *http.Response
	err  error
	body []byte
}

// // RoundTripFunc lets us use a function as an http.RoundTripper.
// type RoundTripFunc func(req *http.Request) (*http.Response, error)

func (f *NewFakeTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.Body != nil {
		b, _ := io.ReadAll(req.Body)
		f.body = b
	}

	return f.resp, f.err
}

// NewFakeResponse creates a fake *http.Response with the provided status and body.
func NewFakeResponse(status int, body string) *http.Response {
	return &http.Response{
		StatusCode: status,
		Body:       io.NopCloser(strings.NewReader(body)),
	}
}

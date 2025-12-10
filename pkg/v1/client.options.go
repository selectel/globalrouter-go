package v1

import "net/http"

type ServiceClientOption func(*ServiceClient)

// WithAuthOpts is a functional parameter for Client, used to set on of implementations of AuthType.
// required argument for using client.
func WithTokenID(tokenID string) ServiceClientOption {
	return func(client *ServiceClient) {
		client.TokenID = tokenID
	}
}

// WithAPIUrl is a functional parameter for Client, used to set IAM API URL.
func WithAPIUrl(url string) ServiceClientOption {
	return func(client *ServiceClient) {
		client.APIUrl = url
	}
}

// WithCustomHTTPClient is a functional parameter for Client, used to set a custom HTTP client.
func WithCustomHTTPClient(httpClient *http.Client) ServiceClientOption {
	return func(client *ServiceClient) {
		client.HTTPClient = httpClient
	}
}

// WithClientUserAgent is a functional parameter for Client, used to set a custom User-Agent prefix.
//
// It is highly recommended to use this option!
func WithClientUserAgent(userAgent string) ServiceClientOption {
	return func(client *ServiceClient) {
		client.ClientUserAgent = userAgent
	}
}

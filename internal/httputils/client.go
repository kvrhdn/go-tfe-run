package httputils

import (
	"context"
	"net/http"
	"strings"

	"github.com/hashicorp/go-retryablehttp"
)

// RetryOnTLSHandhsakeTimeoutClient provides a http.Client that retries
// exclusively on "TLS handshake timeout" errors. It uses the default
// configuration provided by retryablehttp.NewClient().
func RetryOnTLSHandhsakeTimeoutClient() *http.Client {
	retryableHttpClient := retryablehttp.NewClient()
	retryableHttpClient.CheckRetry = retryOnTLSHandshakeTimeout
	return retryableHttpClient.StandardClient()
}

func retryOnTLSHandshakeTimeout(ctx context.Context, resp *http.Response, err error) (bool, error) {
	return strings.Contains(err.Error(), "TLS handshake timeout"), nil
}

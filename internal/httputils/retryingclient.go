package httputils

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/hashicorp/go-retryablehttp"
)

// RetryOnTLSHandhsakeTimeoutClient provides a http.Client that retries
// exclusively on "TLS handshake timeout" errors. It uses the default
// configuration provided by retryablehttp.NewClient().
func RetryOnTLSHandhsakeTimeoutClient() *http.Client {
	retryableHttpClient := retryablehttp.NewClient()
	retryableHttpClient.Logger = discardingLogger()
	retryableHttpClient.CheckRetry = retryOnTLSHandshakeTimeout
	return retryableHttpClient.StandardClient()
}

func discardingLogger() *log.Logger {
	var logger log.Logger
	logger.SetOutput(ioutil.Discard)
	return &logger
}

func retryOnTLSHandshakeTimeout(ctx context.Context, resp *http.Response, err error) (bool, error) {
	// do not retry on context.Canceled or context.DeadlineExceeded
	if ctx.Err() != nil {
		return false, ctx.Err()
	}

	if err != nil {
		return strings.Contains(err.Error(), "TLS handshake timeout"), nil
	}

	return false, nil
}

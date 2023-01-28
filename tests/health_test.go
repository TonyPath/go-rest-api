// go:build e2e

package tests

import (
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/require"
)

func TestHealthChechEndpoint(t *testing.T) {
	client := resty.New()

	resp, err := client.R().Get("http://localhost:8080/alive")
	require.NoError(t, err)

	require.Equal(t, 200, resp.StatusCode())
}

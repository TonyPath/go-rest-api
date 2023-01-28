// go:build e2e

package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/require"
)

func TestPostComment(t *testing.T) {
	t.Run("can post comment", func(t *testing.T) {
		client := resty.New()

		resp, err := client.R().
			SetHeader("Authorization", "bearer "+createToken()).
			SetBody(`{"slug":"/", "auth":"tony", "body":"go rest api"}`).
			Post("http://localhost:8080/api/v1/comments")

		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode())
	})

	t.Run("cannot post comment without JWT", func(t *testing.T) {
		client := resty.New()

		resp, err := client.R().
			SetBody(`{"slug":"/", "auth":"tony", "body":"go rest api"}`).
			Post("http://localhost:8080/api/v1/comments")

		require.NoError(t, err)
		require.Equal(t, http.StatusUnauthorized, resp.StatusCode())
	})
}

func createToken() string {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenStr, err := token.SignedString([]byte("secret"))
	if err != nil {
		fmt.Println(err)
	}

	return tokenStr
}

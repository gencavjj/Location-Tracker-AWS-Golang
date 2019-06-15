package pkg

import (
	"net/http"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	httpMethod := http.MethodGet
	path := "/healthcheck"

	c := New()

	t.Run("Not found", func(t *testing.T) {
		resp, err := c.Handler(events.APIGatewayProxyRequest{
			HTTPMethod: httpMethod,
			Path:       "/favicon.ico",
		})
		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
		assert.Nil(t, err)
	})

	t.Run("Successful Request", func(t *testing.T) {
		resp, err := c.Handler(events.APIGatewayProxyRequest{
			HTTPMethod: httpMethod,
			Path:       path,
		})

		assert.Equal(t, "ok", resp.Body)
		assert.Nil(t, err)
	})
}

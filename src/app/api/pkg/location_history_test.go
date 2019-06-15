package pkg

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestGetLocationHistory(t *testing.T) {
	httpMethod := http.MethodGet
	path := "/api/location-history"

	c := New()

	r := new(Location)
	b, err := json.Marshal(r)
	assert.NoError(t, err)

	t.Run("Successful Request", func(t *testing.T) {
		resp, err := c.Handler(events.APIGatewayProxyRequest{
			HTTPMethod: httpMethod,
			Path:       path,
			Body:       string(b),
		})

		s := new(Location)

		assert.Equal(t, "[{\"locationId\":1,\"latitude\":36.3217,\"longitude\":-118.13116,\"date\":\"2019-06-14T17:30:12Z\",\"time\":\"230\"}]", resp.Body)
		assert.Equal(t, r, s)
		assert.Nil(t, err)
	})
}

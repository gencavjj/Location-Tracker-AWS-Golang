package pkg

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

// Core is the component core.
type Core struct{}

// New returns a new component core.
func New() *Core {
	return &Core{}
}

// Handler receives all the incoming requests.
func (c *Core) Handler(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if r.HTTPMethod == "OPTIONS" {
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("not found"),
			StatusCode: http.StatusOK,
			Headers:    headers(),
		}, nil
	}

	switch route := fmt.Sprintf("%v %v", r.HTTPMethod, r.Path); route {
	case "GET /healthcheck":
		return c.healthCheck(r)
	case "GET /api/location-history":
		return c.GetLocationHistory(r)
	}

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("not found"),
		StatusCode: http.StatusNotFound,
		Headers:    headers(),
	}, nil
}

// headers will return headers to allow CORS.
func headers() map[string]string {
	h := make(map[string]string)
	h["Access-Control-Allow-Origin"] = "*"
	h["Access-Control-Allow-Methods"] = "POST, GET, OPTIONS, PUT, DELETE"
	h["Access-Control-Allow-Headers"] = "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"
	return h
}

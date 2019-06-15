package pkg

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func (c *Core) healthCheck(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("ok"),
		StatusCode: http.StatusOK,
	}, nil
}

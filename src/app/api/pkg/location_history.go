package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
)

// Location is a point on the map.
type Location struct {
	LocationID int       `json:"locationId"`
	Latitude   float64   `json:"latitude"`
	Longitude  float64   `json:"longitude"`
	Date       time.Time `json:"date"`
	Time       string    `json:"time"`
}

// GetLocationHistory outputs the location history in an array.
func (c *Core) GetLocationHistory(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Add location to the array.
	arr := make([]Location, 0)
	l := Location{}
	l.LocationID = 1
	l.Latitude = 36.3217
	l.Longitude = -118.13116
	t, err := time.Parse("01/02/2006 15:04:05 MST", "06/14/2019 17:30:12 MDT")
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("error converting to time: %v", err),
			StatusCode: http.StatusInternalServerError,
			Headers:    headers(),
		}, err
	}
	l.Date = t
	l.Time = "230"

	arr = append(arr, l)

	// Convert the struct to json bytes.
	b, err := json.Marshal(arr)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("error converting to bytes: %v", err),
			StatusCode: http.StatusInternalServerError,
			Headers:    headers(),
		}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       string(b),
		StatusCode: http.StatusOK,
		Headers:    headers(),
	}, nil
}

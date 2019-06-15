package main

import (
	"app/api/pkg"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	c := pkg.New()
	lambda.Start(c.Handler)
}

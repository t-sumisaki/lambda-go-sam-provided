package handler

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
)

type Response events.APIGatewayProxyResponse

func HandleRequest(ctx context.Context) (Response, error) {

	return Response{
		StatusCode: 200,
		Body:       "hello, from SAM",
	}, nil
}

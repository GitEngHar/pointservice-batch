package lambda_handler

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
)

type HealthCheckHandler struct{}

func NewHealthCheckHandler() *HealthCheckHandler {
	return &HealthCheckHandler{}
}

func (h *HealthCheckHandler) Handle(ctx context.Context) (events.APIGatewayProxyResponse, error) {
	// レスポンス本文
	body := fmt.Sprintf("ok")
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    map[string]string{"Content-Type": "text/plain"},
		Body:       body,
	}, nil
}

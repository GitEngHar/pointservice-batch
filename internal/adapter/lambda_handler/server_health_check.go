package lambda_handler

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"pointservice-batch/internal/adapter/client_handler"
	"pointservice-batch/internal/infra/client"
)

type ServerHealthCheckHandler struct{}

func NewServerHealthCheckHandler() *ServerHealthCheckHandler {
	return &ServerHealthCheckHandler{}
}

func (h *ServerHealthCheckHandler) Handle(ctx context.Context) (events.APIGatewayProxyResponse, error) {
	config := client.NewConfig()
	handler := client_handler.NewHttpHealthCheckHandler()
	response, err := client.NewHttpClient(config, handler).Client().Start()
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Headers:    map[string]string{"Content-Type": "text/plain"},
			Body:       err.Error(),
		}, nil
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    map[string]string{"Content-Type": "text/plain"},
		Body:       response,
	}, nil
}

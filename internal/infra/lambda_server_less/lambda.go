package lambda_server_less

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Config struct {
	// 抽象化する
	handler                  Handler
	apiGatewayRequestHandler ApiGatewayRequestHandler
}

type Handler interface {
	Handle(ctx context.Context) (events.APIGatewayProxyResponse, error)
}

type ApiGatewayRequestHandler interface {
	Handle(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Handler(handler Handler) *Config {
	c.handler = handler
	return c
}

func (c *Config) ApiGatewayRequestHandler(apiGatewayRequestHandler ApiGatewayRequestHandler) *Config {
	c.apiGatewayRequestHandler = apiGatewayRequestHandler
	return c
}

func (c *Config) Run() {
	lambda.Start(c.handler.Handle)
}

func (c *Config) ApiGatewayRequestRun() {
	lambda.Start(c.apiGatewayRequestHandler.Handle)
}

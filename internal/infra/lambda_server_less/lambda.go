package lambda_server_less

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Config struct {
	// 抽象化する
	handler Handler
}

type Handler interface {
	Handle(ctx context.Context) (events.APIGatewayProxyResponse, error)
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Handler(handler Handler) *Config {
	c.handler = handler
	return c
}

func (c *Config) Run() {
	lambda.Start(c.handler.Handle)
}

package main

import (
	"log"
	"os"
	"pointservice-batch/internal/adapter/lambda_handler"
	"pointservice-batch/internal/infra/lambda_server_less"
)

func main() {
	mode := os.Getenv("APP_ENV")
	switch mode {
	case "aws":
		handler := lambda_handler.NewVerifyApiGateway()
		lambda_server_less.NewConfig().
			ApiGatewayRequestHandler(handler).
			ApiGatewayRequestRun()
	case "local":
		log.Fatalf("unsupported APP_ENV: %s", mode)
	default:
		log.Fatalf("unsupported APP_ENV: %s", mode)
	}

}

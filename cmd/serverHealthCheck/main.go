package main

import (
	"log"
	"os"
	"pointservice-batch/internal/adapter/echo_handler"
	"pointservice-batch/internal/adapter/lambda_handler"
	"pointservice-batch/internal/infra/echo_server"
	"pointservice-batch/internal/infra/lambda_server_less"
)

func main() {
	mode := os.Getenv("APP_ENV")
	switch mode {
	case "aws":
		handler := lambda_handler.NewServerHealthCheckHandler()
		lambda_server_less.NewConfig().
			Handler(handler).
			Run()
	case "local":
		handler := echo_handler.NewServerHealthCheckHandler()
		echo_server.NewConfig().
			Handler(handler).
			Run()
	default:
		log.Fatalf("unsupported APP_ENV: %s", mode)
	}

}

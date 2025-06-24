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
		handler := lambda_handler.NewHealthCheckHandler()
		lambda_server_less.NewConfig().
			Handler(handler).
			Start()
	case "local":
		handler := echo_handler.NewHealthCheckHandler()
		var app = echo_server.NewConfig().WebServer(handler)
		app.Start()
	default:
		log.Fatalf("unsupported APP_ENV: %s", mode)
	}

}

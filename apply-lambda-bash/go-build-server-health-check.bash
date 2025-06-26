GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
  go build -o bootstrap ~/Documents/pointservice-batch/cmd/serverHealthCheck/main.go && zip function.zip bootstrap

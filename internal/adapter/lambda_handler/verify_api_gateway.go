package lambda_handler

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"log"
)

type VerifyApiGateway struct{}

type Payload struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func NewVerifyApiGateway() *VerifyApiGateway {
	return &VerifyApiGateway{}
}

func (v *VerifyApiGateway) Handle(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// PathParameterの取得
	proxyPath := req.PathParameters["proxy"]

	// json objectを想定
	// bodyの取得
	var p Payload
	if err := json.Unmarshal([]byte(req.Body), &p); err != nil {
		log.Printf("Failed to unmarshal body: %v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       `{"error":"invalid request body"}`,
		}, nil
	}
	log.Printf("Request body: %+v", p)
	resBody, _ := json.Marshal(map[string]interface{}{
		"path":    proxyPath,
		"payload": p,
	})
	return events.APIGatewayProxyResponse{
		StatusCode:      200,
		Headers:         map[string]string{"Content-Type": "application/json"},
		Body:            string(resBody),
		IsBase64Encoded: false,
	}, nil
}

package handlers

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func ApiResponse(status int, body interface{}) (*events.APIGatewayProxyResponse, error) {
	resp := events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Type":                 "application/json",
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Methods": "DELETE, GET, HEAD, OPTIONS, PATCH, POST, PUT",
			"Access-Control-Allow-Headers": "X-Amz-Date,X-Api-Key,X-Amz-Security-Token,X-Requested-With,X-Auth-Token,Referer,User-Agent,Origin,Content-Type,Authorization,Accept,Access-Control-Allow-Methods,Access-Control-Allow-Origin,Access-Control-Allow-Headers",
		}}
	resp.StatusCode = status
	stringBody, _ := json.Marshal(body)
	resp.Body = string(stringBody)
	return &resp, nil
}

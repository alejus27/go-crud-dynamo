package handlers

import (
	"net/http"

	"go-aws/pkg/program"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var ErrorMethodNotAllowed = "method not allowed"

type ErrorBody struct {
	ErrorMsg *string `json:"error,omitempty"`
}

func GetProgram(req *events.APIGatewayProxyRequest, tbName string, client dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {
	id := req.QueryStringParameters["id"]
	if len(id) > 0 {
		result, err := program.FetchProgram(id, tbName, client)
		if err != nil {
			return ApiResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
		}
		return ApiResponse(http.StatusOK, result)
	}

	result, err := program.FetchAllProgram(tbName, client)
	if err != nil {
		return ApiResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}
	return ApiResponse(http.StatusOK, result)
}

func CreateProgram(req *events.APIGatewayProxyRequest, tbName string, client dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {
	result, err := program.CreateProgram(req, tbName, client)
	if err != nil {
		return ApiResponse(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}

	return ApiResponse(http.StatusCreated, result)
}

func UpdateProgram(req *events.APIGatewayProxyRequest, tbName string, client dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {
	result, err := program.UpdateProgram(req, tbName, client)
	if err != nil {
		return ApiResponse(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}

	return ApiResponse(http.StatusOK, result)
}

func DeleteProgram(req *events.APIGatewayProxyRequest, tbName string, client dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {
	err := program.DeleteProgram(req, tbName, client)
	if err != nil {
		return ApiResponse(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}

	return ApiResponse(http.StatusOK, nil)
}

func UnhandledMethod() (*events.APIGatewayProxyResponse, error) {
	return ApiResponse(http.StatusMethodNotAllowed, ErrorMethodNotAllowed)
}

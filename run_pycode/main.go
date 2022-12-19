package main

import (
	"encoding/json"
	"errors"
	"go-aws/pkg/handlers"
	"go-aws/pkg/program"
	"go-aws/run_pycode/runner"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(req *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var codePython runner.CodePython
	if err := json.Unmarshal([]byte(req.Body), &codePython); err != nil {
		return nil, errors.New(program.ErrorInvalidProgramData)
	}
	result := runner.RunPyCode(codePython.Code)

	return handlers.ApiResponse(http.StatusOK, result)
}

func main() {
	lambda.Start(handler)
}

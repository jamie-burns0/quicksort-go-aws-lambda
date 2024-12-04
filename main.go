package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jamie-burns0/quicksort-go/quicksort"
)

type RequestBody struct {
	UnsortedData []int `json:"data"`
}

type ResponseBody struct {
	SortedData []int `json:"data"`
}

func main() {
	lambda.Start(handleRequest)
}

func handleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var requestBody RequestBody
	err := json.Unmarshal([]byte(request.Body), &requestBody)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	sortedData := quicksort.Sort(requestBody.UnsortedData)

	responseBody := ResponseBody{
		SortedData: sortedData,
	}

	jsonBytes, err := json.Marshal(responseBody)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body: string(jsonBytes),
	}

	return response, nil
}
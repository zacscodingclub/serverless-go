package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func findOne(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	movies := []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}{
		{
			ID:   1,
			Name: "Avengers",
		},
		{
			ID:   2,
			Name: "Ant-Man",
		},
		{
			ID:   3,
			Name: "Thor",
		},
		{
			ID:   4,
			Name: "Hulk",
		}, {
			ID:   5,
			Name: "Doctor Strange",
		},
	}
	id, err := strconv.Atoi(req.PathParameters["id"])

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "ID must be a number",
		}, nil
	}

	response, err := json.Marshal(movies[id-1])
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(response),
	}, nil
}

func main() {
	lambda.Start(findOne)
}

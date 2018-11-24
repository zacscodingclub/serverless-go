package main

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Movie struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func insert(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var movie Movie
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
	err := json.Unmarshal([]byte(req.Body), &movie)

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "Invalid movie payload",
		}, nil
	}
	movies = append(movies, movie)

	response, err := json.Marshal(movies)
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
	lambda.Start(insert)
}

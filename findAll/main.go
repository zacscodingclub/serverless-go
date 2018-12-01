package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Movie struct {
	ID   string
	Name string
}

func findAll(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	size, err := strconv.Atoi(request.Headers["Count"])
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "Count Header should be a number",
		}, nil
	}

	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Error while retrieving AWS credentials",
		}, nil
	}

	svc := dynamodb.New(cfg)
	req := svc.ScanRequest(&dynamodb.ScanInput{
		TableName: aws.String(os.Getenv("TABLE")),
		Limit:     aws.Int64(int64(size)),
	})
	res, err := req.Send()
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Error while scanning DynamoDB",
		}, nil
	}

	movies := make([]Movie, 0)
	for _, item := range res.Items {
		movies = append(movies, Movie{
			ID:   *item["ID"].S,
			Name: *item["Name"].S,
		})
	}
	data, err := json.Marshal(movies)

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(data),
	}, nil
}

func main() {
	lambda.Start(findAll)
}

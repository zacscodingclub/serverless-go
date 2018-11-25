package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/dynamodbattribute"

	"github.com/aws/aws-sdk-go-v2/aws/external"
)

type Movie struct {
	ID   string
	Name string
}

func main() {
	cfg, err := external.LoadDefaultAWSConfig()

	movies, err := readMovies("movies.json")
	if err != nil {
		log.Fatal(err)
	}

	for _, movie := range movies {
		log.Printf("Inserting: %s", movie.Name)
		err = insertMovie(cfg, movie)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func readMovies(fn string) ([]Movie, error) {
	movies := make([]Movie, 0)

	data, err := ioutil.ReadFile(fn)
	if err != nil {
		return movies, err
	}

	err = json.Unmarshal(data, &movies)
	if err != nil {
		return movies, err
	}

	return movies, nil
}

func insertMovie(cfg aws.Config, movie Movie) error {
	item, err := dynamodbattribute.MarshalMap(movie)
	if err != nil {
		return err
	}

	svc := dynamodb.New(cfg)
	req := svc.PutItemRequest(&dynamodb.PutItemInput{
		TableName: aws.String("movies_00"),
		Item:      item,
	})
	_, err = req.Send()
	if err != nil {
		return err
	}
	return nil
}

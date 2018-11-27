#!/bin/bash

echo "Deploying Lambda to AWS"

aws lambda update-function-code --function-name FindAllMovies \
    --zip-file fileb://./deployment.zip \
    --region us-west-2
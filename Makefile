build:
	./findAll/build.sh

deploy:
	aws lambda update-function-code --function-name HelloServerless \
		--zip-file fileb://./deployment.zip \
		--region us-west-2
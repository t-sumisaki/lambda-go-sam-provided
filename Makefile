.PHONY: clean

clean:
	rm -rf ./bin

build-MyLambdaFunction:
	GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -ldflags="-w -s" -o ./bin/bootstrap ./
	cp ./bin/bootstrap $(ARTIFACTS_DIR)
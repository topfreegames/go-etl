.PHONY: plugins

plugins:
	@find plugins -type d -path 'plugins/*' -maxdepth 1 -exec bash -c "ls {}/*.go | xargs go build -buildmode=plugin -o {}/main.so" \;

start:
	@go run main.go start

setup:
	@dep ensure

TAG := latest
image:
	@docker build -t tfgco/go-etl:$(TAG) .

push-image:
	@docker push tfgco/go-etl:$(TAG)

plugins-linux:
	@find plugins -type d -path 'plugins/*' -maxdepth 1 -exec sh -c "ls {}/*.go | xargs go build -buildmode=plugin -o {}/main.so" \;

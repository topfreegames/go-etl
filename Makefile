.PHONY: plugins

plugins:
	@find plugins -type d -path 'plugins/*' -maxdepth 1 -exec bash -c "ls {}/*.go | xargs go build -buildmode=plugin -o {}/main.so" \;

start:
	@go run main.go start

setup:
	@dep ensure

image:
	@docker build -t local/go-etl .

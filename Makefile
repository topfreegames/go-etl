.PHONY: plugins

plugins:
	@find plugins -type d -path 'plugins/*' -maxdepth 1 -exec go build -buildmode=plugin -o {}/main.so {}/main.go \;

start:
	@go run main.go start

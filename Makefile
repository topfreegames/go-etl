.PHONY: plugins

plugins:
	@find plugins -type d -path 'plugins/*' -maxdepth 1 -exec go build -buildmode=plugin -o {}/main.so {}/main.go {}/extractor.go {}/transformer.go {}/loader.go \;

start:
	@go run main.go start

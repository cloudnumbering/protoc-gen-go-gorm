build-options:
	buf generate --template proto/options/buf.gen.yaml --path proto/options
build-example:
	go install
	go install github.com/favadi/protoc-go-inject-tag@latest
	go install github.com/mitchellh/protoc-gen-go-json@latest
	buf generate --template example/mysql/buf.gen.yaml --path example/mysql
	protoc-go-inject-tag -input example/mysql/*.*.*.go
clean:
	rm -f example/mysql/*.go
	rm -f options/*.go
generate: clean build-options build-example
test: generate
	go test -v ./test
.PHONY: test

test: clean
	./run-tests.sh

lint:
	golangci-lint run ./...

proto-compile:
	[ -d /usr/local/lib/protobuf/include/hedwig ] || (echo "Ensure github.com/cloudchacho/hedwig is cloned at /usr/local/lib/protobuf/include/hedwig/"; exit 2)
	cd protobuf && protoc -I/usr/local/lib/protobuf/include -I. --go_out=. --go_opt=module=github.com/cloudchacho/hedwig-go/protobuf hedwig/protobuf/container.proto hedwig/protobuf/options.proto
	cd protobuf/internal && protoc -I/usr/local/lib/protobuf/include -I. --go_out=. --go_opt=module=github.com/cloudchacho/hedwig-go/protobuf/internal protobuf.proto protobuf_alternate.proto protobuf_bad.proto
	cd examples && protoc -I/usr/local/lib/protobuf/include -I. --go_out=. --go_opt=module=github.com/cloudchacho/hedwig-go/examples schema.proto

mod-tidy:
	go mod tidy
	cd examples && go mod tidy

clean:
	find . -name coverage.txt -delete

all: generate check test build

generate:
	go-bindata -nomemcopy -nometadata ./data
	gofmt -s -w bindata.go

check:
	gofmt -s -l -d .
	goimports -l -d .
	golangci-lint run \
	--enable-all \
	--disable gochecknoinits \
	--disable gochecknoglobals \
	.

test:
	go test -v .

CGO_ENABLED:=0
build:
	go build \
	-tags 'static' \
	-ldflags '-w -extldflags "-static"' \
	.

clean:
	rm project

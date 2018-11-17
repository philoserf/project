all: generate test

generate:
	go-bindata -nomemcopy -nometadata ./data
	gofmt -s -w bindata.go

test:
	drone exec

clean:
	rm project

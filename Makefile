VERSION     := 0.1.1
GO_BUILDOPT := -ldflags '-s -w -X main.version $(VERSION)'

run:
	go run main.go ${ARGS}

fmt:
	go fmt ./...

build: fmt
	go build $(GO_BUILDOPT) -o bin/kangen main.go

release: fmt
	GOOS=linux GOARCH=amd64 go build $(GO_BUILDOPT) -o bin/kangen$(VERSION).linux-amd64 main.go
	GOOS=linux GOARCH=386 go build $(GO_BUILDOPT) -o bin/kangen$(VERSION).linux-386 main.go

clean:
	rm -f bin/kangen*

install: build
	cp bin/kangen /usr/local/bin/

uninstall: clean
	rm -f /usr/local/bin/kangen

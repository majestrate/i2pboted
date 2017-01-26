BIN = i2pboted

$(BIN): clean
	GOPATH=$(PWD) go clean -v
	GOPATH=$(PWD) go build -v -ldflags "-X i2pbote.Version=git-$(shell git rev-parse HEAD)" -o $(BIN)

clean:
	rm -f $(BIN)

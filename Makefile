all:
	go clean -v
	go build -v -ldflags "-X i2pbote.Version=git-$(shell git rev-parse HEAD)"

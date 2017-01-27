BOTECTL = botectl
BOTED = i2pbote

all: build

build:
	GOPATH=$(PWD) go build -v -ldflags "-X i2pbote.Version=git-$(shell git rev-parse HEAD)" -o $(BOTED)
	GOPATH=$(PWD) go build -v -ldflags "-X i2pbote.Version=git-$(shell git rev-parse HEAD)" -o $(BOTECTL) i2pbote/tools/botectl

test:
	GOPATH=$(PWD) go test -v i2pbote/...

clean:
	rm -f $(BOTED) $(BOTECTL)

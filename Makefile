REPO = $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
BOTECTL = bote-cli
BOTED = boted

all: clean build

build: $(BOTED) $(BOTECTL)

$(BOTED):
	GOPATH=$(REPO) go build -v -ldflags "-X i2pbote.version.Git=-$(shell git rev-parse --short HEAD)" -o $(BOTED)
$(BOTECTL):
	GOPATH=$(REPO) go build -v -ldflags "-X i2pbote.version.Git=-$(shell git rev-parse --short HEAD)" -o $(BOTECTL) i2pbote/tools/botectl

test:
	GOPATH=$(REPO) go test -v i2pbote/...

clean:
	rm -f $(BOTED) $(BOTECTL)

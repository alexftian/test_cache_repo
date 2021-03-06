all: build

TAG?=dev
FLAGS=
ENVVAR=
# 操作系统
GOOS?=linux
COMPONENT=lru
# 下载路径
DOWNLOADPATH=download

build: clean
	$(ENVVAR) GOOS=$(GOOS) go build -o ${COMPONENT}

test: clean build
	$(ENVVAR) go test --test.short -race ./... $(FLAGS)

run: build
	./${COMPONENT}
	make clean

clean:
	rm -rf ${COMPONENT} ${DOWNLOADPATH}
format:
	test -z "$$(find . -path ./vendor -prune -type f -o -name '*.go' -exec gofmt -s -d {} + | tee /dev/stderr)" || \
	test -z "$$(find . -path ./vendor -prune -type f -o -name '*.go' -exec gofmt -s -w {} + | tee /dev/stderr)"

.PHONY: all  build test clean format

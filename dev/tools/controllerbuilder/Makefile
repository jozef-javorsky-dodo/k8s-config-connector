.PHONY: all
all: build

.PHONY: build
build: codebot controllerbuilder benchmark ctf

.PHONY: controllerbuilder
controllerbuilder:
	go build -o ${GOPATH}/bin/controllerbuilder main.go

.PHONY: codebot
codebot:
	go build -o ${GOPATH}/bin/codebot cmd/codebot/main.go

.PHONY: ctf
ctf:
	go build -o ${GOPATH}/bin/ctf cmd/ctf/main.go

.PHONY: benchmark
benchmark:
	go build -o ${GOPATH}/bin/benchmark cmd/benchmark/type/main.go

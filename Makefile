GOPATH:=$(shell pwd)
GO:=go
GOFLAGS:=-v -p 1
EXEC:=$(shell basename `pwd`)
TODAY := $(shell date +%s%N | cut -b1-13)
default: build #bin/${EXEC}

build:
	@echo "========== Building $@ =========="
	sh -c 'export GOPATH=${GOPATH}; $(GO) build $(GOFLAGS) -o ${GOPATH}/bin/${EXEC} fernandomitre7.com/${EXEC}'
install: 
	@echo "========== Compiling $@ =========="
	sh -c 'export GOPATH=${GOPATH}; $(GO) install $(GOFLAGS) fernandomitre7.com/${EXEC}'
clean:
	@echo "Deleting binary files ..."; sh -c 'if [ -f bin/${EXEC} ]; then rm -f bin/${EXEC} && echo ../bin/${EXEC} ;fi;'
	@echo "Moving log files "; sh -c 'if [ -f logs/${EXEC}.log ]; then mv logs/${EXEC}.log logs/${EXEC}.${TODAY}.log; fi;'
.PHONY: all

all: deps.get thrift

thrift:
	./tools/thrift -r --out .  --gen go vr.thrift

deps.get:
	go mod tidy

test:
	go test -v ./tests/

clean:
	@rm -rf vr/GoUnusedProtection__.go
	@rm -rf vr/vr-consts.go
	@rm -rf vr/vr.go

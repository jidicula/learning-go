.DEFAULT_GOAL := build

.PHONY: build run-ok run-cancel clean

build:
	go build

run-ok: build
	./context_cancel false

run-cancel: build
	./context_cancel true

clean:
	rm context_cancel

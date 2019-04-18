SHELL  := /bin/bash
GOPATH :=

all: ragel test

ragel:
	for f in $$(find . -type f -regex '.*\.rl'); do \
		ragel -s -Z -G1 -o "$$(echo $$f | sed 's/\.rl/.go/g')" $$f; \
	done

test:
	go test -cover -coverprofile=cover.out -v -timeout=15m ./... \
	&& go tool cover -html=cover.out -o cover.html

bench:
	go test -v -benchtime 2s -test.bench=. -run='^Bench' ./...

clean:
	rm -f cover.html cover.out

.EXPORT_ALL_VARIABLES:

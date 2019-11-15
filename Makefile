

build: hello

hello:
	GO111MODULE=off go  build -o $@ ./cmd

clean:
	rm -f hello

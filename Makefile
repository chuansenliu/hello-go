

build: hello

hello:
	go build -o $@ ./cmd

clean:
	rm -f hello

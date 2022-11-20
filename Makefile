BIN		=	main

build:
	go build -o "$(BIN)"

clean:
	go clean

run: build
	./"$(BIN)"

.PHONY: build clean run
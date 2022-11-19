BIN		=	main

build:
	go build -o "$(BIN)"

clean:
	go clean

run:
	./"$(BIN)"

.PHONY: build clean run
NAME		=	crud_user
BIN	=	main

build:
	go build -o "$(BIN)"

clean:
	go clean

run:
	./"$(BIN)"

image:
	docker build --tag $(NAME) .

container: image
	docker run  -p 8080:8080 $(NAME)

stop:
	docker stop -t 60 mysql-db

.PHONY: build clean run
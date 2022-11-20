NAME		=	crud_user
BIN	=	main

build:
	go build -o "$(BIN)"

clean:
	go clean

run: build
	./"$(BIN)"

image:
	docker build --tag $(NAME) .

container: image start
	docker run  --name $(NAME) -p 8080:8080 $(NAME)

start:
	docker start mysql-db
	docker start crud_user

stop:
	docker stop -t 60 mysql-db
	docker stop -t 60 crud_user

.PHONY: build clean run
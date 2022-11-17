FROM golang:1.19-alpine 

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY ./services/*.go ./services/
COPY ./router/*.go ./router/
COPY ./database/*.go ./database/
COPY local.env ./
COPY *.go ./
RUN go build -o /crud-user

EXPOSE 3306
#EXPOSE 8080

CMD [ "/crud-user" ]

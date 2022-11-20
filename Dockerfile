FROM golang:alpine as builder

WORKDIR /app
COPY . .
RUN go mod download 
RUN go build -o main .

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /app/main .
COPY --from=builder /app/.env .       

# Expose port 8080 to the outside world
EXPOSE 8080
EXPOSE 3306

#Command to run the executable
CMD ["./main"]
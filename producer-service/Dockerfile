# Use an official Go runtime as a parent image
FROM golang:1.20-bookworm AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go project files into the container
COPY . .

RUN apt-get update && \
    apt-get install -y git gcc libc6 libc6-dev libc6-dbg libbsd-dev && \
    rm -rf /var/lib/apt/lists/*

ENV CGO_ENABLED 1
ENV GOFLAGS -mod=vendor
ENV GOOS=linux
ENV GOARCH=amd64

# Build the Go application
RUN go build -tags musl -tags no_librdkafka_vendor cmd/producer-service/main.go

# Start a new, lightweight imageFROM debian:bullseye-slim
FROM debian:bookworm

# Set the working directory for the final image
WORKDIR /app

# Copy the compiled Go binary from the build image
COPY --from=build /app/main .

# Expose the port your server will listen on
EXPOSE 8080

# Define the command to run your server
ENTRYPOINT ["./main"]

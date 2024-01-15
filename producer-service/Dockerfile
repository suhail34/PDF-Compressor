# Use an official Go runtime as a parent image
FROM golang:1.20 AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go project files into the container
COPY . .

# Build the Go application
RUN go mod download && CGO_ENABLED=0 GOOS=linux go build cmd/producer-service/main.go

# Start a new, lightweight image
FROM gcr.io/distroless/base-debian11

# Set the working directory for the final image
WORKDIR /app

# Copy the compiled Go binary from the build image
COPY --from=build /app/main .

# Expose the port your server will listen on
EXPOSE 8080

# Define the command to run your server
CMD ["./main"]

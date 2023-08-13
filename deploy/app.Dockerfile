# Use an official Golang runtime as a parent image
FROM golang:1.20.7-alpine

# Set the working directory to /go/src/app
WORKDIR /go/src/app

# Copy the source code from the host to the container
COPY . .

# Build the Go application
RUN go build -o /go/bin/app ./cmd/server

# Use an official Alpine Linux runtime as a parent image
FROM alpine:3.14

# Set the working directory to /app
WORKDIR /app

# Copy the binary from the previous stage
COPY --from=0 /go/bin/app .

# Expose port 8080 for the application
EXPOSE 5000

# Run the application
CMD ["./app"]

##########################################
# append these lines below to your docker-compose.yaml
##########################################
#   ent-user-app:
#     build:
#       context: .
#       dockerfile: app.Dockerfile
#     ports:
#       - "8080:5000"
##########################################

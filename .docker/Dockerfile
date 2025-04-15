# Start with a base image containing golang runtime
FROM golang:1.20

# Set the working directory in the container
WORKDIR /go/src/app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Copy the current directory contents into the container at /go/src/app
COPY ../ .

EXPOSE 8080
EXPOSE 8000
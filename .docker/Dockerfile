# Start with a base image containing golang runtime
FROM golang:1.21

# Set the working directory in the container
WORKDIR /go/src/app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Install gcloud CLI
RUN apt-get update && apt-get install -y \
    curl \
    apt-transport-https \
    ca-certificates && \
    echo "deb [signed-by=/usr/share/keyrings/cloud.google.gpg] https://packages.cloud.google.com/apt cloud-sdk main" | tee -a /etc/apt/sources.list.d/google-cloud-sdk.list && \
    curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key --keyring /usr/share/keyrings/cloud.google.gpg add - && \
    apt-get update && apt-get install -y google-cloud-sdk google-cloud-cli-app-engine-go && \
    rm -rf /var/lib/apt/lists/*

# Copy the current directory contents into the container at /go/src/app
COPY ../ .

EXPOSE 8080
EXPOSE 8000
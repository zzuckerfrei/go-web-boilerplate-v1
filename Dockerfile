# Start from the official golang image
FROM golang:1.22

# Install necessary libraries
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        libc6-dev \
        && \
    rm -rf /var/lib/apt/lists/*

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod  ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 80 to the outside world
EXPOSE 80

# Command to run the executable
CMD ["./main"]
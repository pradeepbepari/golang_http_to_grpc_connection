    FROM --platform=linux/amd64 golang:1.23 as builder

    # Set the working directory inside the container
    WORKDIR /app

    # Copy the go.mod and go.sum files to the working directory
    COPY grpc-portal/go.mod grpc-portal/go.sum ./

    # Download the Go module dependencies
    RUN go mod download

    # Copy the rest of the application code to the working directory
    COPY . .

    # Build the Go application
    RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o grpc ./grpc-portal/main.go

    # Use Alpine Linux as the base image for the final build
    FROM  --platform=linux/amd64 alpine:latest

    # Install GLIBC on Alpine
    RUN apk add --no-cache \
        ca-certificates \
        tzdata \
        libc6-compat

    # Set the working directory inside the container
    WORKDIR /app

    # Copy the binary from the builder stage
    COPY --from=builder /app/grpc .

    # Expose the port that the application will run on
    EXPOSE 8000

    # Command to run the application
    CMD ["./grpc"]

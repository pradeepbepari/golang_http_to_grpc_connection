# ---- Base Stage ----
    FROM --platform=linux/amd64 golang:1.23 as builder

    # Set the working directory inside the container
    WORKDIR /app

    # Copy the go.mod and go.sum files to the working directory
    COPY http_server/go.mod http_server/go.sum ./

    # Download the Go module dependencies
    RUN go mod download
    RUN go install github.com/swaggo/swag/cmd/swag@latest
    # Copy the rest of the application code to the working directory
    COPY . .
    # COPY ./grpc-portal/cmd/proto/ ./grpc-portal/cmd/proto/
    # ---- Final Dev Stage (with hot reload) ----
    RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o main ./http_server/main.go

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
    COPY --from=builder /app/main .
    
        # Expose the port that the application will run on
    EXPOSE 8000
    
    # Command to run the application
    CMD ["./main"]
    
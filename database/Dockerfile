# ---- Base Stage ----
    FROM golang:1.23 as builder

    # Set the working directory inside the container
    WORKDIR /app

    # Copy the go.mod and go.sum files to the working directory
    COPY database/go.mod database/go.sum ./

    # Download the Go module dependencies
    RUN go mod download
    COPY . .

    COPY database/migrations /app/migrations

    RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o main ./database/main.go

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
    COPY --from=builder /app/migrations /app/migrations


    
        # Expose the port that the application will run on
    EXPOSE 8000
    
    # Command to run the application
    CMD ["./main"]
    
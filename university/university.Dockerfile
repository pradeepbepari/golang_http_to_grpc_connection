FROM --platform=linux/amd64 golang:1.23-alpine as builder
# Set the working directory inside the container
WORKDIR /app
# Copy the go.mod and go.sum files to the working directory
COPY university/go.mod university/go.sum ./
# Download the Go module dependencies
RUN go mod download
# Copy the rest of the application code to the working directory

COPY . .


# COPY ./grpc-portal/cmd/proto/ ./grpc-portal/cmd/proto/
RUN CGO_ENABLED=0 GOARCH=arm64 GOOS=linux go build -o main ./university/main.go

# Use Alpine Linux as the base image for the final build
FROM  --platform=linux/amd64 alpine:latest
# Install GLIBC on Alpine
RUN apk add --no-cache \
 ca-certificates \
 tzdata \
 libstdc++

# Set the working directory inside the container
WORKDIR /app
# Copy the binary from the builder stage
COPY --from=builder /app/main .
# Expose the port that the application will run on
EXPOSE 8000
# Command to run the application
CMD ["./main"]
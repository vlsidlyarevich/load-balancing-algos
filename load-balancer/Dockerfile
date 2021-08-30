FROM golang:1.15-alpine

# Set necessary environment variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

# Copy the sources into the container
COPY src .

# Install dependencies
RUN go mod download
RUN go mod verify

# Build the application
RUN go build -o main ./cmd/load-balancer

## Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/main .

# Command to run when starting the container
CMD ["/dist/main"]
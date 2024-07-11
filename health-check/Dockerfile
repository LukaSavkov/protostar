# Start from the latest golang base image
FROM golang:latest as builder

# Add Maintainer Info
LABEL maintainer="Janko Rakonjac"

# Set the Current Working Directory inside the container
WORKDIR /app/app

# Copy go mod and sum files
COPY ./protostar/health-check/go.mod ./protostar/health-check/go.sum ./

COPY ./magnetar ../../magnetar
COPY ./oort ../../oort

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy everything from the current directory to the Working Directory inside the container
COPY ./protostar/health-check/ .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .



######## Start a new stage from scratch #######
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/app/main .
RUN mkdir data


# Command to run the executable
CMD ["./main"]

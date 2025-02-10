
# Start from the latest golang base image
FROM golang:1.23 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . . 

# Build the Go app
#RUN go build -o main .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -a -o main .


# Copy the Pre-built binary file from the previous stage
#COPY --from=builder /app/main .

#FROM scratch
#WORKDIR /app
#COPY --from=builder /app/main .

# Expose port 8080 to the outside world
EXPOSE 8080

ENTRYPOINT [ "./main" ]

# Command to run the executable
#CMD ["./main"]



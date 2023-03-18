FROM golang:1.19.0

WORKDIR /Users/macbook/Desktop/docker/xendit-1

RUN go install github.com/cosmtrek/air@latest

# Copy the source from the current directory to the working Directory inside the container
COPY . .
COPY .env .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# Build the Go app
# RUN go build -o /build

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the executable
# CMD [ "/build" ]

RUN go mod tidy
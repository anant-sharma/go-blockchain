FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/anant-sharma/go-blockchain

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go get -d -v ./...

EXPOSE 8080
CMD [ "go", "run", "main.go"]

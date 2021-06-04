#FROM golang:1.16
#
## Set the Current Working Directory inside the container
##WORKDIR $GOPATH/src/github.com/InternAtTouch/docker1
##WORKDIR $APP_HOME
###WORKDIR /app/go-sample-app
###RUN go get -u github.com/InternAtTouch/docker1
### Copy everything from the current directory to the PWD (Present Working Directory) inside the container
###COPY app.go .
##COPY go.mod .
##COPY go.sum .
##
##RUN go mod download
##COPY . .
### Install the package
###RUN go install -v ./...
##ENV GO111MODULE=on
###RUN go build -o ./app.go
### This container exposes port 8080 to the outside world
##EXPOSE 8081
##
### Run the executable
##CMD ["go-sample-app"]
#
#WORKDIR /go/src/app
#COPY . .
##RUN go get -u github.com/gnnchya/InternAtTouch/tree/main
#
#RUN go get -d -v ./...
#RUN go install -v ./...
#
#CMD ["app.go"]

FROM golang:1.16-alpine

# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

# Add Maintainer Info
LABEL maintainer="Rajeev Singh <rajeevhub@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["./main"]
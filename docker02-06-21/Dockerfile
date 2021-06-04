FROM golang:1.16

#RUN apk update && apk upgrade && \
#    apk add --no-cache bash git openssh

# Set the Current Working Directory inside the container
WORKDIR /app

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/main .

EXPOSE 8080
CMD ["./main"]
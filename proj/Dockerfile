FROM golang:1.16.4
WORKDIR /docktest
COPY go.mod .
COPY go.sum .
COPY main.go .
RUN go mod download
COPY . .
RUN go build -o testapp .
CMD ["main.go"]
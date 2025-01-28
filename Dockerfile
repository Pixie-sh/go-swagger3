FROM golang:alpine 
WORKDIR /go/src/main
RUN go install github.com/pixie-sh/go-swagger3@latest

ENTRYPOINT ["go-swagger3"]

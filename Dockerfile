FROM golang:latest

WORKDIR /go/

COPY . .

RUN go mod init main \
    && go mod tidy \
    && go build

CMD ["go", "run", "main.go"]

FROM golang:1.16-alpine

WORKDIR /go/src/app
COPY . .

RUN go mod tidy
RUN go build -o app .

EXPOSE 3000

CMD ["./app"]
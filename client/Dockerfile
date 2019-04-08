FROM golang:1.12.1 as builder
ENV GOBIN=/go/bin
RUN mkdir -p /go/src/github.com/ShwethaKumbla/webcrawler
WORKDIR /go/src/github.com/ShwethaKumbla/webcrawler
COPY . /go/src/github.com/ShwethaKumbla/webcrawler/
RUN go get ./...
RUN CGO_ENABLED=0 GOOS=linux go build -o crawler-client client.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /go/src/github.com/ShwethaKumbla/webcrawler/crawler-client ./

ENTRYPOINT ["./crawler-client"]

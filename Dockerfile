FROM golang:1.12.1
ENV GOBIN=/go/bin
RUN mkdir -p /go/src/github.com/webcrawler
WORKDIR /go/src/github.com/webcrawler
COPY . /go/src/github.com/webcrawler/
RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build -o webcrawler main.go

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/webcrawler/webcrawler .
ENTRYPOINT ["./webcrawler"]

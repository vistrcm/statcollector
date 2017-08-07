# build application phase
FROM golang:1.8.3
WORKDIR /go/src/github.com/vistrcm/statcollector/
#RUN go get -d -v golang.org/x/net/html
COPY main.go .
RUN go build -a -o statcollector .

# build image phase
FROM alpine:latest
#RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/vistrcm/statcollector/statcollector .
CMD ["./statcollector"]

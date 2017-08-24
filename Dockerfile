# build application phase
FROM golang:1.8.3 as builder
WORKDIR /go/src/github.com/vistrcm/statcollector/
COPY ./ .
# handle dependencies
RUN go get -u github.com/kardianos/govendor
RUN govendor sync
# build with specific params to avoid issues of running in alpine
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -o statcollector .

# build image phase
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go/src/github.com/vistrcm/statcollector/statcollector /usr/local/bin/
ENTRYPOINT statcollector

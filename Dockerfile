# build application phase
FROM golang:1.15 as builder

WORKDIR /go/src/github.com/vistrcm/statcollector/
COPY ./ .

# build with specific params to avoid issues of running in alpine
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -o statcollector .

# build image phase
FROM scratch
COPY --from=builder /go/src/github.com/vistrcm/statcollector/statcollector /statcollector
# array in etrypoint is a dirty hack to be able to pass parameters via CMD later
ENTRYPOINT ["/statcollector"]

# build application phase
FROM golang:1.10 as builder

## handle dependencies and linting
RUN echo "installing required packages" \
    && go get -v -u github.com/golang/dep/cmd/dep\
    && go get -v -u github.com/alecthomas/gometalinter \
    && gometalinter --install

WORKDIR /go/src/github.com/vistrcm/statcollector/
COPY ./ .

## handle dependencies and linting
RUN echo "installing deps" \
    && dep ensure -v \
    && echo "let'd do some linting" \
    && gometalinter --deadline=300s --vendor ./...
# build with specific params to avoid issues of running in alpine
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -o statcollector .

# build image phase
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go/src/github.com/vistrcm/statcollector/statcollector /usr/local/bin/
# array in etrypoint is a dirty hack to be able to pass parameters via CMD later
ENTRYPOINT ["statcollector"]

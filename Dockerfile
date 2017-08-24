# build application phase
FROM golang:1.8.3 as builder
WORKDIR /go/src/github.com/vistrcm/statcollector/
COPY ./ .
RUN go build -a -o statcollector .

# build image phase
FROM alpine:latest as runner
#RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/vistrcm/statcollector/statcollector .
CMD ["./statcollector"]

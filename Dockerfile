FROM golang:latest AS build

ADD . /go/src/github.com/nyushi/discord-roomkeeper
WORKDIR /go/src/github.com/nyushi/discord-roomkeeper
RUN go get -u -d && go build

FROM alpine:3.8
RUN apk add --no-cache libc6-compat ca-certificates && update-ca-certificates
COPY --from=build /go/src/github.com/nyushi/discord-roomkeeper /usr/local/bin/
CMD ["/usr/local/bin/discord-roomkeeper"]

FROM golang:1.20-alpine3.17 as builder
WORKDIR /code

COPY ./echo-server.go ./

RUN CGO_ENABLED=0 go build -o server ./echo-server.go

FROM alpine:3.17
WORKDIR /bin
COPY --from=builder /code/server /bin/server
CMD /bin/server

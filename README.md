[![Go Report Card](https://goreportcard.com/badge/github.com/neha-viswanathan/echo-client-server)](https://goreportcard.com/report/github.com/neha-viswanathan/echo-client-server)


# Echo Client-Server in Go

This is a simple echo client-server written in Go, where the server just echoes any message 
that the client sends.


### Getting started

Start the server on a terminal window:
```bash
go run ./server/echo-server.go
```

Start the client on another terminal window:
```bash
go run ./client/echo-client.go
```

The repo also contains a Dockerfile for the server if it needs to be run as a container.

```bash
docker run --rm nviswanathan/echo-server:latest
```

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	addr := "0.0.0.0:8100"
	server, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("error listening: ", err)
	}
	defer server.Close()

	log.Println("server is ready")

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal("error accepting conn: ", err)
		}

		go handleClients(conn)
	}
}

func handleClients(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Print(str)
	}

	log.Println("connection closed")
}

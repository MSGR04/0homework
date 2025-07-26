package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	port := ":8080"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error listening on port %s", port))
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(fmt.Sprintf("Error accepting connection on port %s", port))
			os.Exit(1)
		}
		go handleRequest(conn)
	}

}

func handleRequest(conn net.Conn) {
	defer conn.Close()

	if _, err := conn.Write([]byte("OK\n")); err != nil {
		fmt.Println("Error writing response:", err)
		return
	}

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Received message from client: %s\n", line)
		conn.Write([]byte(fmt.Sprintf("Received message from client: %s\n", line)))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading:", err)
	}
}

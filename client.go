package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	port := "localhost:8080"
	conn, err := net.Dial("tcp", port)
	if err != nil {
		fmt.Println("Error connecting:", err)
	}
	defer conn.Close()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter text to server: ")
	conn.Write([]byte("ping\n"))
	response, _ := bufio.NewReader(conn).ReadString('\n')
	if response == "OK\n" {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK:", response)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from stdin:", err.Error())
	}
}

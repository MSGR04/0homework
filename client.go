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
	for scanner.Scan() {
		text := scanner.Text()
		conn.Write([]byte(text + "\n"))

		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from server:", err)
		}
		if response == "OK\\n" {
			fmt.Println("Text says OK: ")
		} else {
			fmt.Println("Text says not OK, but: ", response)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from stdin:", err.Error())
	}
}

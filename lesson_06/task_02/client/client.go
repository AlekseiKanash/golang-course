package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		fmt.Print("Can't connect\n")
		os.Exit(1)
	}
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Print("buffio read error. Closig connection.\n")
			conn.Close()
			os.Exit(1)
		}

		if text == "exit\n" {
			conn.Close()
			os.Exit(0)
		}

		strtosend := text + "\n"
		i, err := fmt.Fprintf(conn, strtosend)
		if i < len(strtosend) || err != nil {
			fmt.Print("Error writing data to connection. Closig connection.\n")
			conn.Close()
			os.Exit(1)
		}

		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Print("buffio read error. Closig connection.\n")
			conn.Close()
			os.Exit(1)
		}
		fmt.Print("Message from server: " + message)
	}
}

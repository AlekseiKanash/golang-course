package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func TrimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}

func main() {

	fmt.Println("Launching server...")
	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Print("Can't start server\n")
		os.Exit(1)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Print("Can't accept connection\n")
			continue
		}
		go func(conn net.Conn) {
			for {
				message, err := bufio.NewReader(conn).ReadString('\n')
				if err != nil {
					fmt.Print("buffio read error. Closig connection.\n")
					conn.Close()
					return
				}
				fmt.Print("Message Received:", string(message))

				message = TrimSuffix(message, "\n")
				newmessage := ""
				i, err := strconv.ParseInt(message, 10, 64)
				if err == nil {
					newmessage = fmt.Sprintf("%d", 2*i)
				} else {
					newmessage = strings.ToUpper(message)
				}
				strtosend := newmessage + "\n"
				sent, err := conn.Write([]byte(strtosend))
				if int(sent) < len(strtosend) || err != nil {
					fmt.Print("Error writing data to connection. Closig connection.\n")
					conn.Close()
					os.Exit(1)
				}

			}
		}(conn)
	}
}

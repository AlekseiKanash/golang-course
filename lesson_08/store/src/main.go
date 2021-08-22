package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/AlekseiKanash/golang-course/lesson_08/store/src/grpcstore"
)

func getInput(input chan string) int {
	for {
		in := bufio.NewReader(os.Stdin)
		result, err := in.ReadString('\n')
		if err != nil {
			fmt.Print("buffio read error. Closig connection.\n")
			return 1
		}

		input <- result
	}
}

func handleServerLifetime(server *grpcstore.Server) int {
	server.Start()
	defer server.Stop()
	input := make(chan string, 1)
	go getInput(input)

	fmt.Println("Type exit to stop the server.")
	for {
		select {
		case inputStr := <-input:
			switch inputStr {
			case "stop\n":
				{
					server.Stop()
				}
			case "start\n":
				{
					fmt.Println("Start command!")
					server.Start()
					fmt.Println("Done!")
				}
			case "exit\n":
				{
					server.Stop()
					os.Exit(0)
				}
			}
		case <-time.After(1000 * time.Millisecond):
			{
				continue
			}
		}
	}
}

func main() {
	server := grpcstore.Server{Addr: "0.0.0.0:9000"}
	server.Init()
	res := handleServerLifetime(&server)
	os.Exit(res)
}

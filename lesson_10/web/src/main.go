package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/AlekseiKanash/golang-course/lesson_10/web/src/rest"
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

func handleServerLifetime(server *rest.Server) int {
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
			case "test\n":
				{

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
	server := rest.Server{Addr: "0.0.0.0:80"}
	server.Init()
	res := handleServerLifetime(&server)
	os.Exit(res)
}

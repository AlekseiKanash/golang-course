package main

import (
	"context"
	"log"

	pb "github.com/AlekseiKanash/golang-course/lesson_07/proto"
	"google.golang.org/grpc"
)

func register(c pb.EchoClient, user string) {
	reg_response, err := c.Register(context.Background(), &pb.RegisterRequest{Body: user})
	if err != nil {
		log.Fatalf("Can't register user: %s", err)
	}
	log.Printf("Response from server: %d", reg_response.Id)
}

func list(c pb.EchoClient) {
	reg_response, err := c.List(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatalf("Can't list users: %s", err)
	}
	log.Printf("Response from server: %v", reg_response.Records)
}

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := pb.NewEchoClient(conn)

	response, err := c.SayHello(context.Background(), &pb.Message{Body: "Hello From Client!"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)

	register(c, "Aleksei")
	register(c, "Nikolai")
	register(c, "Serhei")
	register(c, "Aleksei")

	list(c)
}

package store

import (
	"context"
	"fmt"
	"log"

	pb "github.com/AlekseiKanash/golang-course/lesson_08/proto"
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

func Report() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial("store:9000", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("did not connect: %s\n", err)
		return
	}
	defer conn.Close()

	c := pb.NewEchoClient(conn)

	response, err := c.SayHello(context.Background(), &pb.Message{Body: "Hello From Client!"})
	if err != nil {
		fmt.Printf("Error when calling SayHello: %s\n", err)
		return
	}
	log.Printf("Response from server: %s", response.Body)

	register(c, "Aleksei")
	register(c, "Nikolai")
	register(c, "Serhei")
	register(c, "Aleksei")

	list(c)
}

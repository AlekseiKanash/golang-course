package store

import (
	"context"
	"fmt"
	"log"

	pb "github.com/AlekseiKanash/golang-course/lesson_08/proto"
	"google.golang.org/grpc"
)

func Save(data *pb.SaveRequest) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("store:9000", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("did not connect: %s\n", err)
		return
	}
	defer conn.Close()

	client := pb.NewEchoClient(conn)

	response, err := client.Save(context.Background(), data)
	if err != nil {
		fmt.Printf("Error when calling Register: %s\n", err)
		return
	}
	log.Printf("Response from server: %s", response.Body)
}

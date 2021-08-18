package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/AlekseiKanash/golang-course/lesson_07/proto"
	"google.golang.org/grpc"
)

type ServerRegisterError struct {
	string Message
}

func (sr ServerRegisterError) String() string {
	return fmt.Sprintf("%s", sr.Message)
}

type Server struct {
	data map[uint32]string
}

func (s *Server) hasValue(value string) (uint32, bool) {
	for n, x := range s.data {
		if x == value {
			return n, true
		}
	}
	return 0, false
}

func (s *Server) SayHello(ctx context.Context, in *pb.Message) (*pb.Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &pb.Message{Body: "Hello From the Server!"}, nil
}

func (s *Server) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	if index, hasValue := s.hasValue(in.Body); hasValue {
		err := pb.RpcGeneralError{Body: "User Already Exists"}
		return &pb.RegisterResponse{Id: index, Error: &err}, &err
	}

	if s.data == nil {
		s.data = make(map[uint32]string)
	}
	newId := uint32(len(s.data))
	s.data[newId] = in.Body
	return &pb.RegisterResponse{Id: newId}, nil
}

func (s *Server) List(ctx context.Context, in *pb.Empty) (*pb.ListResponse, error) {
	return &pb.ListResponse{Records: s.data}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterEchoServer(grpcServer, &Server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

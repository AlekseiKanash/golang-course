package grpcstore

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	pb "github.com/AlekseiKanash/golang-course/lesson_08/proto"
	"google.golang.org/grpc"
)

type Server struct {
	IsRunning bool
	Addr      string
	server    *grpc.Server
	wg        *sync.WaitGroup
	isInit    bool
	data      map[uint32]string
}

type ServerRegisterError struct {
	Message string
}

func (sr ServerRegisterError) String() string {
	return fmt.Sprintf("%s", sr.Message)
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
		return &pb.RegisterResponse{Id: index, Error: &err}, nil
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

func (s *Server) Init() {

}

func (s *Server) Start() {
	if s.IsRunning {
		fmt.Println("Already running.")
		return
	}

	if nil == s.wg {
		s.wg = &sync.WaitGroup{}
	}

	s.wg.Add(1)

	fmt.Println("Starting GRPC server")

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s.server = grpc.NewServer()

	pb.RegisterEchoServer(s.server, &Server{})

	go func(wg *sync.WaitGroup) {
		defer fmt.Println("done")
		defer wg.Done() // let main know we are done cleaning up

		// always returns error. ErrServerClosed on graceful close
		s.IsRunning = true
		if err := s.server.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %s", err)
		}
		fmt.Printf("Server is stopped. %v\n", err)
		s.IsRunning = false

	}(s.wg)
}

func (s *Server) Stop() {
	if s.IsRunning {
		s.server.Stop()
		s.wg.Wait()
	}
}

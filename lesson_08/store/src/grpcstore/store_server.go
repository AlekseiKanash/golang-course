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
	data      []pb.SaveRequest
}

func (s *Server) Save(ctx context.Context, in *pb.SaveRequest) (*pb.SaveResponse, error) {

	s.data = append(s.data, *in)
	ret_str := fmt.Sprintf("Saved. Total: %d", len(s.data))
	return &pb.SaveResponse{Body: ret_str}, nil
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

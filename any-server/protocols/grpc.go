package protocols

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	// "google.golang.org/grpc/peer"
	"google.golang.org/grpc/reflection"

	ecpb "google.golang.org/grpc/examples/features/proto/echo"
	hwpb "google.golang.org/grpc/examples/helloworld/helloworld"
)

// hwServer is used to implement helloworld.GreeterServer.
type hwServer struct {
	hwpb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *hwServer) SayHello(ctx context.Context, in *hwpb.HelloRequest) (*hwpb.HelloReply, error) {
	// peer.FromContext(ctx)
	log.Printf("receiving call, %v", in)
	return &hwpb.HelloReply{Message: "Hello " + in.Name}, nil
}

type ecServer struct {
	ecpb.UnimplementedEchoServer
}

func (s *ecServer) UnaryEcho(ctx context.Context, req *ecpb.EchoRequest) (*ecpb.EchoResponse, error) {
	return &ecpb.EchoResponse{Message: req.Message}, nil
}

func StartGRPCServer(port string) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("grpc server listening at %v\n", lis.Addr())

	s := grpc.NewServer()

	// Register Greeter on the server.
	hwpb.RegisterGreeterServer(s, &hwServer{})

	// Register RouteGuide on the same server.
	ecpb.RegisterEchoServer(s, &ecServer{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

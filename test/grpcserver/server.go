package grpcserver

import (
	"context"
	"net"

	proto "github.com/shunta-furukawa/seq-gen/test/grpcserver/proto/grpcserver"
	"google.golang.org/grpc"
)

// SimpleServer is a test implementation of the gRPC server.
type SimpleServer struct {
	pb.UnimplementedTestServiceServer
}

func (s *SimpleServer) TestMethod(ctx context.Context, req *proto.TestRequest) (*proto.TestResponse, error) {
	return &pb.TestResponse{Message: "Hello " + req.Name}, nil
}

func StartGRPCServer() (*grpc.Server, net.Listener, error) {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		return nil, nil, err
	}

	grpcServer := grpc.NewServer()
	pb.RegisterTestServiceServer(grpcServer, &SimpleServer{})
	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			panic(err)
		}
	}()

	return grpcServer, listener, nil
}

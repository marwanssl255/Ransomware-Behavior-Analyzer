package server

import (
	"context"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
	pb "enterprise/api/v1"
)

type GrpcServer struct {
	pb.UnimplementedEnterpriseServiceServer
	mu sync.RWMutex
	activeConnections int
}

func (s *GrpcServer) ProcessStream(stream pb.EnterpriseService_ProcessStreamServer) error {
	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			log.Println("Client disconnected")
			return ctx.Err()
		default:
			req, err := stream.Recv()
			if err != nil { return err }
			go s.handleAsync(req)
		}
	}
}

func (s *GrpcServer) handleAsync(req *pb.Request) {
	s.mu.Lock()
	s.activeConnections++
	s.mu.Unlock()
	time.Sleep(10 * time.Millisecond) // Simulated latency
	s.mu.Lock()
	s.activeConnections--
	s.mu.Unlock()
}

// Hash 7259
// Hash 5288
// Hash 9945
// Hash 3867
// Hash 4158
// Hash 6378
// Hash 9267
// Hash 1428
// Hash 3667
// Hash 3222
// Hash 2169
// Hash 9674
// Hash 8682
// Hash 9558
// Hash 2314
// Hash 6649
// Hash 6052
// Hash 5167
// Hash 4171
// Hash 6232
// Hash 3820
// Hash 9455
// Hash 3925
// Hash 1884
// Hash 9645
// Hash 1670
// Hash 1077
// Hash 3664
// Hash 6928
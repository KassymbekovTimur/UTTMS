package grpc

import (
	"context"
	"google.golang.org/grpc"
	"log"

	"github.com/KassymbekovTimur/UTTMS/statistics/internal/usecase"
	pb "github.com/KassymbekovTimur/UTTMS/statistics/proto"
)

// Server implements gRPC-service Statistics
type Server struct {
	uc *usecase.StatsUsecase
	pb.UnimplementedStatisticsServer
}

// NewGRPCServer creates and registers gRPC-server
func NewGRPCServer(uc *usecase.StatsUsecase) *grpc.Server {
	s := grpc.NewServer()
	srv := &Server{uc: uc}
	pb.RegisterStatisticsServer(s, srv)
	return s
}

func (s *Server) GetUserOrdersStatistics(ctx context.Context, req *pb.UserOrderStatisticsRequest) (*pb.UserOrderStatisticsResponse, error) {
	log.Printf("[gRPC] GetUserOrdersStatistics user=%s", req.UserId)
	total, byHour, err := s.uc.GetUserOrderStats(req.UserId)
	if err != nil {
		return nil, err
	}
	return &pb.UserOrderStatisticsResponse{TotalOrders: int32(total), OrdersByHour: byHour}, nil
}

func (s *Server) GetUserStatistics(ctx context.Context, req *pb.UserStatisticsRequest) (*pb.UserStatisticsResponse, error) {
	log.Println("[gRPC] GetUserStatistics")
	total, reg, err := s.uc.GetUserStats()
	if err != nil {
		return nil, err
	}
	return &pb.UserStatisticsResponse{TotalUsers: int32(total), TotalRegistered: int32(reg)}, nil
}

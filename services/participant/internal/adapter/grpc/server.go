package grpcadapter

import (
	"context"
	"log"

	"github.com/KassymbekovTimur/UTTMS/participant/internal/usecase"
	pb "github.com/KassymbekovTimur/UTTMS/participant/proto"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedParticipantServiceServer
	uc *usecase.ParticipantUsecase
}

func NewGRPCServer(uc *usecase.ParticipantUsecase) *grpc.Server {
	s := grpc.NewServer()
	pb.RegisterParticipantServiceServer(s, &Server{uc: uc})
	return s
}

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	log.Printf("[GRPC] Register: %s", req.Email)
	p, err := s.uc.Register(req.Name, req.Email)
	if err != nil {
		return nil, err
	}
	return &pb.RegisterResponse{ParticipantId: p.ID}, nil
}

func (s *Server) ConfirmEmail(ctx context.Context, req *pb.ConfirmEmailRequest) (*pb.EmptyResponse, error) {
	log.Printf("[GRPC] ConfirmEmail: %s", req.Token)
	err := s.uc.ConfirmEmail(req.Token)
	return &pb.EmptyResponse{}, err
}

func (s *Server) JoinSchedule(ctx context.Context, req *pb.JoinRequest) (*pb.EmptyResponse, error) {
	log.Printf("[GRPC] JoinSchedule: %s to %s", req.ParticipantId, req.ScheduleId)
	err := s.uc.JoinSchedule(req.ParticipantId, req.ScheduleId)
	return &pb.EmptyResponse{}, err
}

func (s *Server) LeaveSchedule(ctx context.Context, req *pb.LeaveRequest) (*pb.EmptyResponse, error) {
	log.Printf("[GRPC] LeaveSchedule: %s from %s", req.ParticipantId, req.ScheduleId)
	err := s.uc.LeaveSchedule(req.ParticipantId, req.ScheduleId)
	return &pb.EmptyResponse{}, err
}

func (s *Server) GetParticipant(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	log.Printf("[GRPC] GetParticipant: %s", req.ParticipantId)
	p, err := s.uc.GetParticipant(req.ParticipantId)
	if err != nil {
		return nil, err
	}
	return &pb.GetResponse{Participant: &pb.Participant{
		Id:          p.ID,
		Name:        p.Name,
		Email:       p.Email,
		ScheduleIds: p.ScheduleIDs,
		Status:      p.Status,
	}}, nil
}

func (s *Server) ListParticipants(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
	log.Printf("[GRPC] ListParticipants for schedule %s", req.ScheduleId)
	list, err := s.uc.ListParticipants(req.ScheduleId)
	if err != nil {
		return nil, err
	}
	res := []*pb.Participant{}
	for _, p := range list {
		res = append(res, &pb.Participant{
			Id:          p.ID,
			Name:        p.Name,
			Email:       p.Email,
			ScheduleIds: p.ScheduleIDs,
			Status:      p.Status,
		})
	}
	return &pb.ListResponse{Participants: res}, nil
}

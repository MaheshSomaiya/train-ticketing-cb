package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	"train-ticket-grpc/proto/proto"

	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedTrainServiceServer
	mu             sync.Mutex
	tickets        map[string]*proto.TicketResponse
	users          map[string]*proto.User
	allocatedSeats map[string]string
}

func (s *server) PurchaseTicket(ctx context.Context, req *proto.PurchaseRequest) (*proto.TicketResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	seat := "A1"
	if len(s.allocatedSeats) > 0 {
		seat = "B1"
	}

	user := &proto.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
	}

	ticket := &proto.TicketResponse{
		From:  req.From,
		To:    req.To,
		User:  user,
		Seat:  seat,
		Price: 20,
	}

	s.tickets[req.Email] = ticket
	s.users[req.Email] = user
	s.allocatedSeats[req.Email] = seat

	return ticket, nil
}

func (s *server) ViewTicket(ctx context.Context, req *proto.ViewRequest) (*proto.TicketResponse, error) {
	ticket, exists := s.tickets[req.Email]
	if !exists {
		return nil, fmt.Errorf("ticket not found for email: %s", req.Email)
	}
	return ticket, nil
}

func (s *server) ViewUsers(ctx context.Context, req *proto.ViewRequest) (*proto.UsersResponse, error) {
	var userList []*proto.User
	for _, user := range s.users {
		userList = append(userList, user)
	}
	return &proto.UsersResponse{Users: userList}, nil
}

func (s *server) RemoveUser(ctx context.Context, req *proto.RemoveRequest) (*proto.RemoveResponse, error) {
	delete(s.tickets, req.Email)
	delete(s.users, req.Email)
	delete(s.allocatedSeats, req.Email)
	return &proto.RemoveResponse{Success: true}, nil
}

func (s *server) ModifySeat(ctx context.Context, req *proto.ModifySeatRequest) (*proto.TicketResponse, error) {
	ticket, exists := s.tickets[req.Email]
	if !exists {
		return nil, fmt.Errorf("ticket not found for email: %s", req.Email)
	}

	ticket.Seat = req.NewSeat
	return ticket, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterTrainServiceServer(grpcServer, &server{tickets: make(map[string]*proto.TicketResponse), users: make(map[string]*proto.User), allocatedSeats: make(map[string]string)})

	fmt.Println("Server listening at :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

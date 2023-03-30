package main

import (
	"context"

	pb "github.com/feexie/Project_Bullion"
)

const (
	port = ":50051"
)

type UserManagementServer struct {
	pb.UnimplementedUsermanagementServer
}

func (s *UserManagementServer) createNewUser(ctx context.Context, in *pb.NewUser) (*pb.User, error)

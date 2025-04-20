package interfaces

import pb "github.com/raihan-faza/rent-my-garden/proto"

type User interface {
	CreateUser(pb.RegisterRequest) (pb.TokenResponse, error)
	UpdateUser(pb.RegisterRequest) error
	DeleteUser(pb.RegisterRequest) error
}

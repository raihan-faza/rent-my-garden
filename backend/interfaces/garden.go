package interfaces

import pb "github.com/raihan-faza/rent-my-garden/proto"

type Garden interface {
	CreateGarden(request pb.GardenRequest) (pb.GardenResponse, error)
	UpdateGarden(request pb.GardenRequest) (pb.GardenResponse, error)
	DeleteGarden(request pb.GardenRequest) (pb.GardenResponse, error)
}

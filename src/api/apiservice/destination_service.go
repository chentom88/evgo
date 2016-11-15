package apiservice

import (
	"api/api_proto"
	"fmt"

	"golang.org/x/net/context"
)

type destinationService struct{}

func NewDestinationService() *destinationService {
	return &destinationService{}
}

func (d *destinationService) AddDestination(ctx context.Context, destination *api_proto.Destination) (*api_proto.ApiResponse, error) {
	fmt.Printf("Received destination %+v \n", destination)

	return &api_proto.ApiResponse{
		Code:    0,
		Message: "Successfully added destination.",
	}, nil
}

package service

import (
	"context"
	"errors"

	pb "project/api/PermAddress"
	"project/internal/data"
	"project/internal/validations"
)

type PermAddressServiceService struct {
	pb.UnimplementedPermAddressServiceServer
	repo *data.PermAddressRepository
}

func NewPermAddressServiceService(repo *data.PermAddressRepository) *PermAddressServiceService {
	return &PermAddressServiceService{repo: repo}
}

func (s *PermAddressServiceService) CreatePermAddress(ctx context.Context, req *pb.CreatePermAddressRequest) (*pb.CreatePermAddressResponse, error) {

	if validations.IsValidZipCode(req.Requestbody.Zipcode) {
		err := s.repo.CreatePermAddress(req.Requestbody.DoorNo, req.Requestbody.Street, req.Requestbody.City,
			req.Requestbody.Zipcode, req.Requestbody.State, req.Requestbody.Country)
		if err != nil {
			return nil, err
		}
		return &pb.CreatePermAddressResponse{PermAddId: 0000}, nil
	} else {
		return &pb.CreatePermAddressResponse{PermAddId: 999}, errors.New("invalid zipcode")
	}

}

func (s *PermAddressServiceService) GetPermAddress(ctx context.Context, req *pb.GetPermAddressRequest) (*pb.GetPermAddressResponse, error) {

	permaddress, err := s.repo.GetByIDPermAdress(req.PermAddId)
	if err != nil {
		return nil, err
	}
	return &pb.GetPermAddressResponse{
		PermAddId: permaddress.Perm_add_id,
		DoorNo:    permaddress.Door_no,
		Street:    permaddress.Street,
		City:      permaddress.City,
		Zipcode:   permaddress.Zipcode,
		State:     permaddress.State,
		Country:   permaddress.Country,
	}, nil
}

func (s *PermAddressServiceService) UpdatePermAddress(ctx context.Context, req *pb.UpdatePermAddressRequest) (*pb.UpdatePermAddressResponse, error) {

	if validations.IsValidZipCode(req.Requestbody.Zipcode) {
		err := s.repo.UpdatePermAddress(req.PermAddId, req.Requestbody.DoorNo, req.Requestbody.Street, req.Requestbody.City,
			req.Requestbody.Zipcode, req.Requestbody.State, req.Requestbody.Country)
		if err != nil {
			return nil, err
		}
		return &pb.UpdatePermAddressResponse{PermAddId: req.PermAddId}, nil
	} else {
		return &pb.UpdatePermAddressResponse{PermAddId: 999}, errors.New("invalid zipcode")
	}

}

func (s *PermAddressServiceService) DeletePermAddress(ctx context.Context, req *pb.DeletePermAddressRequest) (*pb.DeletePermAddressResponse, error) {

	err := s.repo.DeletePermAddress(req.PermAddId)
	if err != nil {
		return nil, err
	}
	return &pb.DeletePermAddressResponse{PermAddId: req.PermAddId}, nil
}

func (s *PermAddressServiceService) GetPermAddresses(ctx context.Context, req *pb.GetPermAddressesRequest) (*pb.GetPermAddressesResponse, error) {

	permaddresses, err := s.repo.GetAllPermAddress()
	if err != nil {
		return nil, err
	}
	var permadressResponses []*pb.PermAddress
	for _, c := range *permaddresses {
		permadressResponses = append(permadressResponses, &pb.PermAddress{
			PermAddId: c.Perm_add_id,
			DoorNo:    c.Door_no,
			Street:    c.Street,
			City:      c.City,
			Zipcode:   c.Zipcode,
			State:     c.State,
			Country:   c.Country,
		})
	}
	return &pb.GetPermAddressesResponse{Permadresses: permadressResponses}, nil
}

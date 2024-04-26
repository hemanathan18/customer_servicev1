package service

import (
	"context"
	"errors"

	pb "project/api/TempAddress"
	"project/internal/data"
	"project/internal/validations"
)

type TempAddressServiceService struct {
	pb.UnimplementedTempAddressServiceServer
	repo *data.TempAddressRepository
}

func NewTempAddressServiceService(repo *data.TempAddressRepository) *TempAddressServiceService {
	return &TempAddressServiceService{repo: repo}
}

func (s *TempAddressServiceService) CreateTempAddress(ctx context.Context, req *pb.CreateTempAddressRequest) (*pb.CreateTempAddressResponse, error) {

	if validations.IsValidZipCode(req.Requestbody.Zipcode) {
		err := s.repo.CreateTempAddress(req.Requestbody.DoorNo, req.Requestbody.Street, req.Requestbody.City,
			req.Requestbody.Zipcode, req.Requestbody.State, req.Requestbody.Country)
		if err != nil {
			return nil, err
		}
		return &pb.CreateTempAddressResponse{TempAddId: 7788}, nil
	} else {
		return &pb.CreateTempAddressResponse{TempAddId: 999}, errors.New("invalid zipcode")
	}

}

func (s *TempAddressServiceService) GetTempAddress(ctx context.Context, req *pb.GetTempAddressRequest) (*pb.GetTempAddressResponse, error) {

	tempaddress, err := s.repo.GetByIDTempAdress(req.TempAddId)
	if err != nil {
		return nil, err
	}

	return &pb.GetTempAddressResponse{
		TempAddId: tempaddress.Temp_add_id,
		DoorNo:    tempaddress.Door_no,
		Street:    tempaddress.Street,
		City:      tempaddress.City,
		Zipcode:   tempaddress.Zipcode,
		State:     tempaddress.State,
		Country:   tempaddress.Country,
	}, nil
}

func (s *TempAddressServiceService) UpdateTempAddress(ctx context.Context, req *pb.UpdateTempAddressRequest) (*pb.UpdateTempAddressResponse, error) {

	if validations.IsValidZipCode(req.Requestbody.Zipcode) {
		err := s.repo.UpdateTempAddress(req.TempAddId, req.Requestbody.DoorNo, req.Requestbody.Street, req.Requestbody.City,
			req.Requestbody.Zipcode, req.Requestbody.State, req.Requestbody.Country)
		if err != nil {
			return nil, err
		}
		return &pb.UpdateTempAddressResponse{TempAddId: req.TempAddId}, nil
	} else {
		return &pb.UpdateTempAddressResponse{TempAddId: 999}, errors.New("invalid zipcode")
	}

}

func (s *TempAddressServiceService) DeleteTempAddress(ctx context.Context, req *pb.DeleteTempAddressRequest) (*pb.DeleteTempAddressResponse, error) {

	err := s.repo.DeleteTempAddress(req.TempAddId)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteTempAddressResponse{TempAddId: req.TempAddId}, nil
}

func (s *TempAddressServiceService) GetTempAddresses(ctx context.Context, req *pb.GetTempAddressesRequest) (*pb.GetTempAddressesResponse, error) {

	tempaddresses, err := s.repo.GetAllTempAddress()
	if err != nil {
		return nil, err
	}
	var tempadressResponses []*pb.TempAddress
	for _, c := range *tempaddresses {
		tempadressResponses = append(tempadressResponses, &pb.TempAddress{
			TempAddId: c.Temp_add_id,
			DoorNo:    c.Door_no,
			Street:    c.Street,
			City:      c.City,
			Zipcode:   c.Zipcode,
			State:     c.State,
			Country:   c.Country,
		})
	}
	return &pb.GetTempAddressesResponse{Tempadresses: tempadressResponses}, nil
}

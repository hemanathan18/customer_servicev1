package service

import (
	"context"
	"errors"

	pb "project/api/OfficialNumber"
	"project/internal/data"
	"project/internal/validations"
)

type OffPhoneServiceService struct {
	pb.UnimplementedOffPhoneServiceServer
	repo *data.OfficialNumberRepository
}

func NewOffPhoneServiceService(repo *data.OfficialNumberRepository) *OffPhoneServiceService {
	return &OffPhoneServiceService{repo: repo}
}

func (s *OffPhoneServiceService) CreateOffPhoneNumber(ctx context.Context, req *pb.CreateOffPhoneNumberRequest) (*pb.CreateOffPhoneNumberResponse, error) {

	if validations.IsValidPhoneNumber(req.Requestbody.Number) {

		err := s.repo.CreateOffPhone(req.Requestbody.CountryCode, req.Requestbody.Number)
		if err != nil {
			return nil, err
		}
		return &pb.CreateOffPhoneNumberResponse{OffPhId: 7777}, nil
	} else {
		return &pb.CreateOffPhoneNumberResponse{OffPhId: 999}, errors.New("phone number not valid")
	}

}
func (s *OffPhoneServiceService) GetOffPhoneNumber(ctx context.Context, req *pb.GetOffPhoneNumberRequest) (*pb.GetOffPhoneNumberResponse, error) {

	offphone, err := s.repo.GetByIDOffPhone(req.OffPhId)
	if err != nil {
		return nil, err
	}

	return &pb.GetOffPhoneNumberResponse{
		OffPhId:     offphone.Off_ph_id,
		CountryCode: offphone.Country_Code,
		Number:      offphone.Number,
	}, nil
}

func (s *OffPhoneServiceService) UpdateOffPhoneNumber(ctx context.Context, req *pb.UpdateOffPhoneNumberRequest) (*pb.UpdateOffPhoneNumberResponse, error) {

	if validations.IsValidPhoneNumber(req.Requestbody.Number) {

		err := s.repo.UpdateOffPhone(req.OffPhId, req.Requestbody.CountryCode, req.Requestbody.Number)
		if err != nil {
			return nil, err
		}
		return &pb.UpdateOffPhoneNumberResponse{OffPhId: req.OffPhId}, nil
	} else {
		return &pb.UpdateOffPhoneNumberResponse{OffPhId: 999}, errors.New("phone number not valid")
	}

}

func (s *OffPhoneServiceService) DeleteOffPhoneNumber(ctx context.Context, req *pb.DeleteOffPhoneNumberRequest) (*pb.DeleteOffPhoneNumberResponse, error) {

	err := s.repo.DeleteOffPhone(req.OffPhId)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteOffPhoneNumberResponse{OffPhId: req.OffPhId}, nil
}

func (s *OffPhoneServiceService) GetOffPhoneNumbers(ctx context.Context, req *pb.GetOffPhoneNumbersRequest) (*pb.GetOffPhoneNumbersResponse, error) {

	offphones, err := s.repo.GetAllOffPhone()
	if err != nil {
		return nil, err
	}
	var offphoneResponses []*pb.OffPhoneNumber
	for _, c := range *offphones {
		offphoneResponses = append(offphoneResponses, &pb.OffPhoneNumber{
			OffPhId:     c.Off_ph_id,
			CountryCode: c.Country_Code,
			Number:      c.Number,
		})
	}

	return &pb.GetOffPhoneNumbersResponse{Offphonenumbers: offphoneResponses}, nil
}

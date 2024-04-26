package service

import (
	"context"
	"errors"

	pb "project/api/PersonalNumber"
	"project/internal/data"
	"project/internal/validations"
)

type PersPhoneServiceService struct {
	pb.UnimplementedPersPhoneServiceServer
	repo *data.PersonalNumberRepository
}

func NewPersPhoneServiceService(repo *data.PersonalNumberRepository) *PersPhoneServiceService {
	return &PersPhoneServiceService{repo: repo}
}

func (s *PersPhoneServiceService) CreatePersPhoneNumber(ctx context.Context, req *pb.CreatePersPhoneNumberRequest) (*pb.CreatePersPhoneNumberResponse, error) {

	if validations.IsValidPhoneNumber(req.Requestbody.Number) {

		err := s.repo.CreatePersPhone(req.Requestbody.CountryCode, req.Requestbody.Number)
		if err != nil {
			return nil, err
		}
		return &pb.CreatePersPhoneNumberResponse{PersPhId: 77777}, nil
	} else {
		return &pb.CreatePersPhoneNumberResponse{PersPhId: 999}, errors.New("phone number not valid")
	}

}

func (s *PersPhoneServiceService) GetPersPhoneNumber(ctx context.Context, req *pb.GetPersPhoneNumberRequest) (*pb.GetPersPhoneNumberResponse, error) {

	persphone, err := s.repo.GetByIDPersPhone(req.PersPhId)
	if err != nil {
		return nil, err
	}
	return &pb.GetPersPhoneNumberResponse{
		PersPhId:    persphone.Pers_ph_id,
		CountryCode: persphone.Country_Code,
		Number:      persphone.Number,
	}, nil
}

func (s *PersPhoneServiceService) UpdatePersPhoneNumber(ctx context.Context, req *pb.UpdatePersPhoneNumberRequest) (*pb.UpdatePersPhoneNumberResponse, error) {

	if validations.IsValidPhoneNumber(req.Requestbody.Number) {
		err := s.repo.UpdatePersPhone(req.PersPhId, req.Requestbody.CountryCode, req.Requestbody.Number)
		if err != nil {
			return nil, err
		}
		return &pb.UpdatePersPhoneNumberResponse{PersPhId: req.PersPhId}, nil
	} else {
		return &pb.UpdatePersPhoneNumberResponse{PersPhId: 999}, errors.New("phone number not valid")
	}

}

func (s *PersPhoneServiceService) DeletePersPhoneNumber(ctx context.Context, req *pb.DeletePersPhoneNumberRequest) (*pb.DeletePersPhoneNumberResponse, error) {

	err := s.repo.DeletePersPhone(req.PersPhId)
	if err != nil {
		return nil, err
	}
	return &pb.DeletePersPhoneNumberResponse{PersPhId: req.PersPhId}, nil
}

func (s *PersPhoneServiceService) GetPersPhoneNumbers(ctx context.Context, req *pb.GetPersPhoneNumbersRequest) (*pb.GetPersPhoneNumbersResponse, error) {

	persphones, err := s.repo.GetAllPersPhone()
	if err != nil {
		return nil, err
	}
	var persphoneResponses []*pb.PersPhoneNumber
	for _, c := range *persphones {
		persphoneResponses = append(persphoneResponses, &pb.PersPhoneNumber{
			PersPhId:    c.Pers_ph_id,
			CountryCode: c.Country_Code,
			Number:      c.Number,
		})
	}
	return &pb.GetPersPhoneNumbersResponse{Persphonenumbers: persphoneResponses}, nil
}

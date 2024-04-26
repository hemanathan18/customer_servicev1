package service

import (
	"context"
	"errors"

	pb "project/api/OfficialMail"
	"project/internal/data"
	"project/internal/validations"
)

type OffMailServiceService struct {
	pb.UnimplementedOffMailServiceServer
	repo *data.OfficialMailRepository
}

func NewOffMailServiceService(repo *data.OfficialMailRepository) *OffMailServiceService {
	return &OffMailServiceService{repo: repo}
}

func (s *OffMailServiceService) CreateOffMail(ctx context.Context, req *pb.CreateOffMailRequest) (*pb.CreateOffMailResponse, error) {

	if validations.IsValidEmail(req.Requestbody.Email) {
		err := s.repo.CreateOffMail(req.Requestbody.Email)
		if err != nil {
			return nil, err
		}
		return &pb.CreateOffMailResponse{OffEmId: 77}, nil
	} else {
		return &pb.CreateOffMailResponse{OffEmId: 999}, errors.New("email format wrong")
	}

}

func (s *OffMailServiceService) GetOffMail(ctx context.Context, req *pb.GetOffMailRequest) (*pb.GetOffMailResponse, error) {

	offmail, err := s.repo.GetByIDOffMail(req.OffEmId)
	if err != nil {
		return nil, err
	}

	return &pb.GetOffMailResponse{
		OffEmId: offmail.Off_em_id,
		Email:   offmail.Email,
	}, nil

}
func (s *OffMailServiceService) UpdateOffMail(ctx context.Context, req *pb.UpdateOffMailRequest) (*pb.UpdateOffMailResponse, error) {

	if validations.IsValidEmail(req.Requestbody.Email) {

		err := s.repo.UpdateOffMail(req.OffEmId, req.Requestbody.Email)
		if err != nil {
			return nil, err
		}
		return &pb.UpdateOffMailResponse{OffEmId: req.OffEmId}, nil
	} else {
		return &pb.UpdateOffMailResponse{OffEmId: 999}, errors.New("email format wrong")
	}
}

func (s *OffMailServiceService) DeleteOffMail(ctx context.Context, req *pb.DeleteOffMailRequest) (*pb.DeleteOffMailResponse, error) {

	err := s.repo.DeleteOffMail(req.OffEmId)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteOffMailResponse{OffEmId: req.OffEmId}, nil
}
func (s *OffMailServiceService) GetOffMails(ctx context.Context, req *pb.GetOffMailsRequest) (*pb.GetOffMailsResponse, error) {

	offmails, err := s.repo.GetAllOffMail()
	if err != nil {
		return nil, err
	}
	var offmailResponses []*pb.OffMail
	for _, c := range *offmails {
		offmailResponses = append(offmailResponses, &pb.OffMail{
			OffEmId: c.Off_em_id,
			Email:   c.Email,
		})
	}

	return &pb.GetOffMailsResponse{Offemails: offmailResponses}, nil
}

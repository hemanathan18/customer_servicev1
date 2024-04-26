package service

import (
	"context"
	"errors"

	pb "project/api/PersonalMail"
	"project/internal/data"
	"project/internal/validations"
)

type PersMailServiceService struct {
	pb.UnimplementedPersMailServiceServer
	repo *data.PersonalMailRepository
}

func NewPersMailServiceService(repo *data.PersonalMailRepository) *PersMailServiceService {
	return &PersMailServiceService{repo: repo}
}

func (s *PersMailServiceService) CreatePersMail(ctx context.Context, req *pb.CreatePersMailRequest) (*pb.CreatePersMailResponse, error) {

	if validations.IsValidEmail(req.Requestbody.Email) {
		err := s.repo.CreatePersMail(req.Requestbody.Email)
		if err != nil {
			return nil, err
		}
		return &pb.CreatePersMailResponse{PersEmId: 777}, nil
	} else {
		return &pb.CreatePersMailResponse{PersEmId: 999}, errors.New("email format wrong")
	}

}

func (s *PersMailServiceService) GetPersMail(ctx context.Context, req *pb.GetPersMailRequest) (*pb.GetPersMailResponse, error) {

	persmail, err := s.repo.GetByIDPersMail(req.PersEmId)
	if err != nil {
		return nil, err
	}
	return &pb.GetPersMailResponse{
		PersEmId: persmail.Pers_em_id,
		Email:    persmail.Email,
	}, nil
}

func (s *PersMailServiceService) UpdatePersMail(ctx context.Context, req *pb.UpdatePersMailRequest) (*pb.UpdatePersMailResponse, error) {

	if validations.IsValidEmail(req.Requestbody.Email) {
		err := s.repo.UpdatePersMail(req.PersEmId, req.Requestbody.Email)
		if err != nil {
			return nil, err
		}
		return &pb.UpdatePersMailResponse{PersEmId: req.PersEmId}, nil
	} else {
		return &pb.UpdatePersMailResponse{PersEmId: 999}, errors.New("email format wrong")
	}
}

func (s *PersMailServiceService) DeletePersMail(ctx context.Context, req *pb.DeletePersMailRequest) (*pb.DeletePersMailResponse, error) {

	err := s.repo.DeletePersMail(req.PersEmId)
	if err != nil {
		return nil, err
	}
	return &pb.DeletePersMailResponse{PersEmId: req.PersEmId}, nil
}

func (s *PersMailServiceService) GetPersMails(ctx context.Context, req *pb.GetPersMailsRequest) (*pb.GetPersMailsResponse, error) {

	persmails, err := s.repo.GetAllPersMail()
	if err != nil {
		return nil, err
	}
	var persmailResponses []*pb.PersMail
	for _, c := range *persmails {
		persmailResponses = append(persmailResponses, &pb.PersMail{
			PersEmId: c.Pers_em_id,
			Email:    c.Email,
		})
	}
	return &pb.GetPersMailsResponse{Persemails: persmailResponses}, nil
}

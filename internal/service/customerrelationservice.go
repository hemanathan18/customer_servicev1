package service

import (
	"context"
	"fmt"

	pb "project/api/CustomerRelation"
	"project/internal/data"
)

type CustomerRelationServiceService struct {
	pb.UnimplementedCustomerRelationServiceServer
	repo *data.CustomerRelationRepository
}

func NewCustomerRelationServiceService(repo *data.CustomerRelationRepository) *CustomerRelationServiceService {
	return &CustomerRelationServiceService{repo: repo}
}

func (s *CustomerRelationServiceService) CreateCustomerRelation(ctx context.Context, req *pb.CreateCustomerRelationRequest) (*pb.CreateCustomerRelationResponse, error) {

	err := s.repo.CreateRelation(req.Requestbody.CustomerId, req.Requestbody.OffEmId, req.Requestbody.PersEmId,
		req.Requestbody.OffPhId, req.Requestbody.PersPhId, req.Requestbody.TempAddId, req.Requestbody.PermAddId)
	if err != nil {
		return nil, err
	}

	return &pb.CreateCustomerRelationResponse{Message: "Customer relation created"}, nil
}
func (s *CustomerRelationServiceService) GetCustomerRelation(ctx context.Context, req *pb.GetCustomerRelationRequest) (*pb.GetCustomerRelationResponse, error) {
	customerrelation, err := s.repo.GetByIDRelation(req.CrId)
	if err != nil {
		return nil, err
	}
	return &pb.GetCustomerRelationResponse{
		CrId:       customerrelation.Cr_Id,
		CustomerId: customerrelation.Customer_ID,
		OffEmId:    customerrelation.Off_em_id,
		PersEmId:   customerrelation.Pers_em_id,
		OffPhId:    customerrelation.Off_ph_id,
		PersPhId:   customerrelation.Pers_ph_id,
		TempAddId:  customerrelation.Temp_add_id,
		PermAddId:  customerrelation.Perm_add_id,
	}, nil
}

func (s *CustomerRelationServiceService) UpdateCustomerRelation(ctx context.Context, req *pb.UpdateCustomerRelationRequest) (*pb.UpdateCustomerRelationResponse, error) {

	err := s.repo.UpdateRelation(req.CrId, req.Requestbody.CustomerId, req.Requestbody.OffEmId, req.Requestbody.PersEmId,
		req.Requestbody.OffPhId, req.Requestbody.PersPhId, req.Requestbody.TempAddId, req.Requestbody.PermAddId)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateCustomerRelationResponse{Message: "Updated customer relation id:" + fmt.Sprint(req.CrId)}, nil
}

func (s *CustomerRelationServiceService) DeleteCustomerRelation(ctx context.Context, req *pb.DeleteCustomerRelationRequest) (*pb.DeleteCustomerRelationResponse, error) {

	err := s.repo.DeleteRelation(req.CrId)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteCustomerRelationResponse{Message: "Deleted customer relation Id:" + fmt.Sprint(req.CrId)}, nil
}

func (s *CustomerRelationServiceService) GetCustomerRelations(ctx context.Context, req *pb.GetCustomerRelationsRequest) (*pb.GetCustomerRelationsResponse, error) {
	customerrelations, err := s.repo.GetAllRelation()
	if err != nil {
		return nil, err
	}
	var customerrelationResponses []*pb.CustomerRelation
	for _, c := range *customerrelations {
		customerrelationResponses = append(customerrelationResponses, &pb.CustomerRelation{
			CrId:       c.Cr_Id,
			CustomerId: c.Customer_ID,
			OffEmId:    c.Off_em_id,
			PersEmId:   c.Pers_em_id,
			OffPhId:    c.Off_ph_id,
			PersPhId:   c.Pers_ph_id,
			TempAddId:  c.Temp_add_id,
			PermAddId:  c.Perm_add_id,
		})
	}
	return &pb.GetCustomerRelationsResponse{Customerrelations: customerrelationResponses}, nil
}

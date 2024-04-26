package service

import (
	context "context"
	"fmt"
	pb "project/api/Customer"
	"project/internal/data"
	"time"
)

type CustomerService struct {
	pb.UnimplementedCustomerServiceServer
	repo *data.CustomerRepository
}

func NewCustomerService(repo *data.CustomerRepository) *CustomerService {
	return &CustomerService{
		repo: repo,
	}
}

func (s *CustomerService) CreateCustomer(ctx context.Context, req *pb.CreateCustomerRequest) (*pb.CreateCustomerResponse, error) {

	//validations pending

	dob, err1 := time.Parse(time.DateOnly, req.Requestbody.Dob)
	if err1 != nil {
		return nil, err1
	}

	err := s.repo.Create(req.Requestbody.Name, req.Requestbody.Gender, dob)
	if err != nil {
		return nil, err
	}
	return &pb.CreateCustomerResponse{Message: "Customer created: " + req.Requestbody.Name}, nil
}

func (s *CustomerService) GetCustomer(ctx context.Context, req *pb.GetCustomerRequest) (*pb.GetCustomerResponse, error) {
	customer, err := s.repo.GetByID(req.CustomerId)
	if err != nil {
		return nil, err
	}
	return &pb.GetCustomerResponse{
		CustomerId: customer.Customer_ID,
		Name:       customer.Name,
		Gender:     customer.Gender,
		Dob:        customer.DOB.String(),
	}, nil
}

func (s *CustomerService) UpdateCustomer(ctx context.Context, req *pb.UpdateCustomerRequest) (*pb.UpdateCustomerResponse, error) {

	//validations pending

	dob, err1 := time.Parse(time.DateOnly, req.Requestbody.Dob)
	if err1 != nil {
		return nil, err1
	}

	err := s.repo.Update(req.CustomerId, req.Requestbody.Name, req.Requestbody.Gender, dob)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateCustomerResponse{Message: "Updated customerID: " + fmt.Sprint(req.CustomerId)}, nil
}

func (s *CustomerService) DeleteCustomer(ctx context.Context, req *pb.DeleteCustomerRequest) (*pb.DeleteCustomerResponse, error) {
	err := s.repo.Delete(req.CustomerId)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteCustomerResponse{Success: true}, nil
}

func (s *CustomerService) GetCustomers(ctx context.Context, req *pb.GetCustomersRequest) (*pb.GetCustomersResponse, error) {
	customers, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	var customerResponses []*pb.Customer
	for _, c := range *customers {
		customerResponses = append(customerResponses, &pb.Customer{
			CustomerId: c.Customer_ID,
			Name:       c.Name,
			Gender:     c.Gender,
			Dob:        c.DOB.String(),
		})
	}
	return &pb.GetCustomersResponse{Customers: customerResponses}, nil
}

package server

import (
	c1 "project/api/Customer"
	cr "project/api/CustomerRelation"
	om "project/api/OfficialMail"
	on "project/api/OfficialNumber"
	pa "project/api/PermAddress"
	pm "project/api/PersonalMail"
	pn "project/api/PersonalNumber"
	ta "project/api/TempAddress"
	"project/internal/conf"
	"project/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, cus *service.CustomerService, offm *service.OffMailServiceService, offp *service.OffPhoneServiceService,
	padd *service.PermAddressServiceService, pmail *service.PersMailServiceService, pphone *service.PersPhoneServiceService,
	tadd *service.TempAddressServiceService, cusrel *service.CustomerRelationServiceService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			validate.Validator(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	c1.RegisterCustomerServiceServer(srv, cus)
	om.RegisterOffMailServiceServer(srv, offm)
	on.RegisterOffPhoneServiceServer(srv, offp)
	pa.RegisterPermAddressServiceServer(srv, padd)
	pm.RegisterPersMailServiceServer(srv, pmail)
	pn.RegisterPersPhoneServiceServer(srv, pphone)
	ta.RegisterTempAddressServiceServer(srv, tadd)
	cr.RegisterCustomerRelationServiceServer(srv, cusrel)
	return srv
}

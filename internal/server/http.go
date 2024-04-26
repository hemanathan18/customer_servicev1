package server

import (
	v1 "project/api/Customer"
	cr2 "project/api/CustomerRelation"
	om1 "project/api/OfficialMail"
	on1 "project/api/OfficialNumber"
	pa1 "project/api/PermAddress"
	pm1 "project/api/PersonalMail"
	pn1 "project/api/PersonalNumber"
	ta1 "project/api/TempAddress"
	"project/internal/conf"
	"project/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, cus *service.CustomerService, offm *service.OffMailServiceService, offp *service.OffPhoneServiceService,
	padd *service.PermAddressServiceService, pmail *service.PersMailServiceService, pphone *service.PersPhoneServiceService,
	tadd *service.TempAddressServiceService, cusrel *service.CustomerRelationServiceService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			validate.Validator(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterCustomerServiceHTTPServer(srv, cus)
	om1.RegisterOffMailServiceHTTPServer(srv, offm)
	on1.RegisterOffPhoneServiceHTTPServer(srv, offp)
	pa1.RegisterPermAddressServiceHTTPServer(srv, padd)
	pm1.RegisterPersMailServiceHTTPServer(srv, pmail)
	pn1.RegisterPersPhoneServiceHTTPServer(srv, pphone)
	ta1.RegisterTempAddressServiceHTTPServer(srv, tadd)
	cr2.RegisterCustomerRelationServiceHTTPServer(srv, cusrel)
	return srv
}

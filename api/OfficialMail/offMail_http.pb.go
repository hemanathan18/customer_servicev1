// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             v5.26.1
// source: offMail.proto

package OfficialMail

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationOffMailServiceCreateOffMail = "/OfficialMail.OffMailService/CreateOffMail"
const OperationOffMailServiceDeleteOffMail = "/OfficialMail.OffMailService/DeleteOffMail"
const OperationOffMailServiceGetOffMail = "/OfficialMail.OffMailService/GetOffMail"
const OperationOffMailServiceGetOffMails = "/OfficialMail.OffMailService/GetOffMails"
const OperationOffMailServiceUpdateOffMail = "/OfficialMail.OffMailService/UpdateOffMail"

type OffMailServiceHTTPServer interface {
	CreateOffMail(context.Context, *CreateOffMailRequest) (*CreateOffMailResponse, error)
	DeleteOffMail(context.Context, *DeleteOffMailRequest) (*DeleteOffMailResponse, error)
	GetOffMail(context.Context, *GetOffMailRequest) (*GetOffMailResponse, error)
	GetOffMails(context.Context, *GetOffMailsRequest) (*GetOffMailsResponse, error)
	UpdateOffMail(context.Context, *UpdateOffMailRequest) (*UpdateOffMailResponse, error)
}

func RegisterOffMailServiceHTTPServer(s *http.Server, srv OffMailServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/off_mails", _OffMailService_CreateOffMail0_HTTP_Handler(srv))
	r.GET("/off_mails/{off_em_id}", _OffMailService_GetOffMail0_HTTP_Handler(srv))
	r.PUT("/off_mails/{off_em_id}", _OffMailService_UpdateOffMail0_HTTP_Handler(srv))
	r.DELETE("/off_mails/{off_em_id}", _OffMailService_DeleteOffMail0_HTTP_Handler(srv))
	r.GET("/off_mails", _OffMailService_GetOffMails0_HTTP_Handler(srv))
}

func _OffMailService_CreateOffMail0_HTTP_Handler(srv OffMailServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateOffMailRequest
		if err := ctx.Bind(&in.Requestbody); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationOffMailServiceCreateOffMail)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateOffMail(ctx, req.(*CreateOffMailRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateOffMailResponse)
		return ctx.Result(200, reply)
	}
}

func _OffMailService_GetOffMail0_HTTP_Handler(srv OffMailServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetOffMailRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationOffMailServiceGetOffMail)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetOffMail(ctx, req.(*GetOffMailRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetOffMailResponse)
		return ctx.Result(200, reply)
	}
}

func _OffMailService_UpdateOffMail0_HTTP_Handler(srv OffMailServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateOffMailRequest
		if err := ctx.Bind(&in.Requestbody); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationOffMailServiceUpdateOffMail)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateOffMail(ctx, req.(*UpdateOffMailRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateOffMailResponse)
		return ctx.Result(200, reply)
	}
}

func _OffMailService_DeleteOffMail0_HTTP_Handler(srv OffMailServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteOffMailRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationOffMailServiceDeleteOffMail)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteOffMail(ctx, req.(*DeleteOffMailRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteOffMailResponse)
		return ctx.Result(200, reply)
	}
}

func _OffMailService_GetOffMails0_HTTP_Handler(srv OffMailServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetOffMailsRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationOffMailServiceGetOffMails)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetOffMails(ctx, req.(*GetOffMailsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetOffMailsResponse)
		return ctx.Result(200, reply)
	}
}

type OffMailServiceHTTPClient interface {
	CreateOffMail(ctx context.Context, req *CreateOffMailRequest, opts ...http.CallOption) (rsp *CreateOffMailResponse, err error)
	DeleteOffMail(ctx context.Context, req *DeleteOffMailRequest, opts ...http.CallOption) (rsp *DeleteOffMailResponse, err error)
	GetOffMail(ctx context.Context, req *GetOffMailRequest, opts ...http.CallOption) (rsp *GetOffMailResponse, err error)
	GetOffMails(ctx context.Context, req *GetOffMailsRequest, opts ...http.CallOption) (rsp *GetOffMailsResponse, err error)
	UpdateOffMail(ctx context.Context, req *UpdateOffMailRequest, opts ...http.CallOption) (rsp *UpdateOffMailResponse, err error)
}

type OffMailServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewOffMailServiceHTTPClient(client *http.Client) OffMailServiceHTTPClient {
	return &OffMailServiceHTTPClientImpl{client}
}

func (c *OffMailServiceHTTPClientImpl) CreateOffMail(ctx context.Context, in *CreateOffMailRequest, opts ...http.CallOption) (*CreateOffMailResponse, error) {
	var out CreateOffMailResponse
	pattern := "/off_mails"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationOffMailServiceCreateOffMail))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in.Requestbody, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *OffMailServiceHTTPClientImpl) DeleteOffMail(ctx context.Context, in *DeleteOffMailRequest, opts ...http.CallOption) (*DeleteOffMailResponse, error) {
	var out DeleteOffMailResponse
	pattern := "/off_mails/{off_em_id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationOffMailServiceDeleteOffMail))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *OffMailServiceHTTPClientImpl) GetOffMail(ctx context.Context, in *GetOffMailRequest, opts ...http.CallOption) (*GetOffMailResponse, error) {
	var out GetOffMailResponse
	pattern := "/off_mails/{off_em_id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationOffMailServiceGetOffMail))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *OffMailServiceHTTPClientImpl) GetOffMails(ctx context.Context, in *GetOffMailsRequest, opts ...http.CallOption) (*GetOffMailsResponse, error) {
	var out GetOffMailsResponse
	pattern := "/off_mails"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationOffMailServiceGetOffMails))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *OffMailServiceHTTPClientImpl) UpdateOffMail(ctx context.Context, in *UpdateOffMailRequest, opts ...http.CallOption) (*UpdateOffMailResponse, error) {
	var out UpdateOffMailResponse
	pattern := "/off_mails/{off_em_id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationOffMailServiceUpdateOffMail))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in.Requestbody, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             v5.26.1
// source: tempAddress.proto

package TempAddress

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

const OperationTempAddressServiceCreateTempAddress = "/TempAddress.TempAddressService/CreateTempAddress"
const OperationTempAddressServiceDeleteTempAddress = "/TempAddress.TempAddressService/DeleteTempAddress"
const OperationTempAddressServiceGetTempAddress = "/TempAddress.TempAddressService/GetTempAddress"
const OperationTempAddressServiceGetTempAddresses = "/TempAddress.TempAddressService/GetTempAddresses"
const OperationTempAddressServiceUpdateTempAddress = "/TempAddress.TempAddressService/UpdateTempAddress"

type TempAddressServiceHTTPServer interface {
	CreateTempAddress(context.Context, *CreateTempAddressRequest) (*CreateTempAddressResponse, error)
	DeleteTempAddress(context.Context, *DeleteTempAddressRequest) (*DeleteTempAddressResponse, error)
	GetTempAddress(context.Context, *GetTempAddressRequest) (*GetTempAddressResponse, error)
	GetTempAddresses(context.Context, *GetTempAddressesRequest) (*GetTempAddressesResponse, error)
	UpdateTempAddress(context.Context, *UpdateTempAddressRequest) (*UpdateTempAddressResponse, error)
}

func RegisterTempAddressServiceHTTPServer(s *http.Server, srv TempAddressServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/temp_addresses", _TempAddressService_CreateTempAddress0_HTTP_Handler(srv))
	r.GET("/temp_addresses/{temp_add_id}", _TempAddressService_GetTempAddress0_HTTP_Handler(srv))
	r.PUT("/temp_addresses/{temp_add_id}", _TempAddressService_UpdateTempAddress0_HTTP_Handler(srv))
	r.DELETE("/temp_addresses/{temp_add_id}", _TempAddressService_DeleteTempAddress0_HTTP_Handler(srv))
	r.GET("/temp_addresses", _TempAddressService_GetTempAddresses0_HTTP_Handler(srv))
}

func _TempAddressService_CreateTempAddress0_HTTP_Handler(srv TempAddressServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateTempAddressRequest
		if err := ctx.Bind(&in.Requestbody); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTempAddressServiceCreateTempAddress)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateTempAddress(ctx, req.(*CreateTempAddressRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateTempAddressResponse)
		return ctx.Result(200, reply)
	}
}

func _TempAddressService_GetTempAddress0_HTTP_Handler(srv TempAddressServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetTempAddressRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTempAddressServiceGetTempAddress)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetTempAddress(ctx, req.(*GetTempAddressRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetTempAddressResponse)
		return ctx.Result(200, reply)
	}
}

func _TempAddressService_UpdateTempAddress0_HTTP_Handler(srv TempAddressServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateTempAddressRequest
		if err := ctx.Bind(&in.Requestbody); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTempAddressServiceUpdateTempAddress)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateTempAddress(ctx, req.(*UpdateTempAddressRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateTempAddressResponse)
		return ctx.Result(200, reply)
	}
}

func _TempAddressService_DeleteTempAddress0_HTTP_Handler(srv TempAddressServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteTempAddressRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTempAddressServiceDeleteTempAddress)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteTempAddress(ctx, req.(*DeleteTempAddressRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteTempAddressResponse)
		return ctx.Result(200, reply)
	}
}

func _TempAddressService_GetTempAddresses0_HTTP_Handler(srv TempAddressServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetTempAddressesRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTempAddressServiceGetTempAddresses)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetTempAddresses(ctx, req.(*GetTempAddressesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetTempAddressesResponse)
		return ctx.Result(200, reply)
	}
}

type TempAddressServiceHTTPClient interface {
	CreateTempAddress(ctx context.Context, req *CreateTempAddressRequest, opts ...http.CallOption) (rsp *CreateTempAddressResponse, err error)
	DeleteTempAddress(ctx context.Context, req *DeleteTempAddressRequest, opts ...http.CallOption) (rsp *DeleteTempAddressResponse, err error)
	GetTempAddress(ctx context.Context, req *GetTempAddressRequest, opts ...http.CallOption) (rsp *GetTempAddressResponse, err error)
	GetTempAddresses(ctx context.Context, req *GetTempAddressesRequest, opts ...http.CallOption) (rsp *GetTempAddressesResponse, err error)
	UpdateTempAddress(ctx context.Context, req *UpdateTempAddressRequest, opts ...http.CallOption) (rsp *UpdateTempAddressResponse, err error)
}

type TempAddressServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewTempAddressServiceHTTPClient(client *http.Client) TempAddressServiceHTTPClient {
	return &TempAddressServiceHTTPClientImpl{client}
}

func (c *TempAddressServiceHTTPClientImpl) CreateTempAddress(ctx context.Context, in *CreateTempAddressRequest, opts ...http.CallOption) (*CreateTempAddressResponse, error) {
	var out CreateTempAddressResponse
	pattern := "/temp_addresses"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTempAddressServiceCreateTempAddress))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in.Requestbody, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *TempAddressServiceHTTPClientImpl) DeleteTempAddress(ctx context.Context, in *DeleteTempAddressRequest, opts ...http.CallOption) (*DeleteTempAddressResponse, error) {
	var out DeleteTempAddressResponse
	pattern := "/temp_addresses/{temp_add_id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTempAddressServiceDeleteTempAddress))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *TempAddressServiceHTTPClientImpl) GetTempAddress(ctx context.Context, in *GetTempAddressRequest, opts ...http.CallOption) (*GetTempAddressResponse, error) {
	var out GetTempAddressResponse
	pattern := "/temp_addresses/{temp_add_id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTempAddressServiceGetTempAddress))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *TempAddressServiceHTTPClientImpl) GetTempAddresses(ctx context.Context, in *GetTempAddressesRequest, opts ...http.CallOption) (*GetTempAddressesResponse, error) {
	var out GetTempAddressesResponse
	pattern := "/temp_addresses"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTempAddressServiceGetTempAddresses))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *TempAddressServiceHTTPClientImpl) UpdateTempAddress(ctx context.Context, in *UpdateTempAddressRequest, opts ...http.CallOption) (*UpdateTempAddressResponse, error) {
	var out UpdateTempAddressResponse
	pattern := "/temp_addresses/{temp_add_id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTempAddressServiceUpdateTempAddress))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in.Requestbody, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
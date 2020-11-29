// Code generated by goa v3.2.5, DO NOT EDIT.
//
// members gRPC client encoders and decoders
//
// Command:
// $ goa gen members/design

package client

import (
	"context"
	memberspb "members/gen/grpc/members/pb"
	members "members/gen/members"

	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// BuildAddFunc builds the remote method to invoke for "members" service "add"
// endpoint.
func BuildAddFunc(grpccli memberspb.MembersClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Add(ctx, reqpb.(*memberspb.AddRequest), opts...)
		}
		return grpccli.Add(ctx, &memberspb.AddRequest{}, opts...)
	}
}

// EncodeAddRequest encodes requests sent to members add endpoint.
func EncodeAddRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*members.AddPayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("members", "add", "*members.AddPayload", v)
	}
	return NewAddRequest(payload), nil
}

// DecodeAddResponse decodes responses from the members add endpoint.
func DecodeAddResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	message, ok := v.(*memberspb.AddResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("members", "add", "*memberspb.AddResponse", v)
	}
	res := NewAddResult(message)
	return res, nil
}
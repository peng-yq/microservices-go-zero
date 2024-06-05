// Code generated by goctl. DO NOT EDIT.
// Source: payment.proto

package payment

import (
	"context"

	"microservices-go-zero/app/payment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreatePaymentReq                     = pb.CreatePaymentReq
	CreatePaymentResp                    = pb.CreatePaymentResp
	GetPaymentBySnReq                    = pb.GetPaymentBySnReq
	GetPaymentBySnResp                   = pb.GetPaymentBySnResp
	GetPaymentSuccessRefundByOrderSnReq  = pb.GetPaymentSuccessRefundByOrderSnReq
	GetPaymentSuccessRefundByOrderSnResp = pb.GetPaymentSuccessRefundByOrderSnResp
	PaymentDetail                        = pb.PaymentDetail
	UpdateTradeStateReq                  = pb.UpdateTradeStateReq
	UpdateTradeStateResp                 = pb.UpdateTradeStateResp

	Payment interface {
		// create a WeChat payment pre-processing order
		CreatePayment(ctx context.Context, in *CreatePaymentReq, opts ...grpc.CallOption) (*CreatePaymentResp, error)
		// query transaction records based on sn
		GetPaymentBySn(ctx context.Context, in *GetPaymentBySnReq, opts ...grpc.CallOption) (*GetPaymentBySnResp, error)
		// Query transaction records based on order sn
		GetPaymentSuccessRefundByOrderSn(ctx context.Context, in *GetPaymentSuccessRefundByOrderSnReq, opts ...grpc.CallOption) (*GetPaymentSuccessRefundByOrderSnResp, error)
		// update trade state
		UpdateTradeState(ctx context.Context, in *UpdateTradeStateReq, opts ...grpc.CallOption) (*UpdateTradeStateResp, error)
	}

	defaultPayment struct {
		cli zrpc.Client
	}
)

func NewPayment(cli zrpc.Client) Payment {
	return &defaultPayment{
		cli: cli,
	}
}

// create a WeChat payment pre-processing order
func (m *defaultPayment) CreatePayment(ctx context.Context, in *CreatePaymentReq, opts ...grpc.CallOption) (*CreatePaymentResp, error) {
	client := pb.NewPaymentClient(m.cli.Conn())
	return client.CreatePayment(ctx, in, opts...)
}

// query transaction records based on sn
func (m *defaultPayment) GetPaymentBySn(ctx context.Context, in *GetPaymentBySnReq, opts ...grpc.CallOption) (*GetPaymentBySnResp, error) {
	client := pb.NewPaymentClient(m.cli.Conn())
	return client.GetPaymentBySn(ctx, in, opts...)
}

// Query transaction records based on order sn
func (m *defaultPayment) GetPaymentSuccessRefundByOrderSn(ctx context.Context, in *GetPaymentSuccessRefundByOrderSnReq, opts ...grpc.CallOption) (*GetPaymentSuccessRefundByOrderSnResp, error) {
	client := pb.NewPaymentClient(m.cli.Conn())
	return client.GetPaymentSuccessRefundByOrderSn(ctx, in, opts...)
}

// update trade state
func (m *defaultPayment) UpdateTradeState(ctx context.Context, in *UpdateTradeStateReq, opts ...grpc.CallOption) (*UpdateTradeStateResp, error) {
	client := pb.NewPaymentClient(m.cli.Conn())
	return client.UpdateTradeState(ctx, in, opts...)
}
// Code generated by goctl. DO NOT EDIT.
// Source: payment.proto

package server

import (
	"context"

	"microservices-go-zero/app/payment/cmd/rpc/internal/logic"
	"microservices-go-zero/app/payment/cmd/rpc/internal/svc"
	"microservices-go-zero/app/payment/cmd/rpc/pb"
)

type PaymentServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedPaymentServer
}

func NewPaymentServer(svcCtx *svc.ServiceContext) *PaymentServer {
	return &PaymentServer{
		svcCtx: svcCtx,
	}
}

// create a WeChat payment pre-processing order
func (s *PaymentServer) CreatePayment(ctx context.Context, in *pb.CreatePaymentReq) (*pb.CreatePaymentResp, error) {
	l := logic.NewCreatePaymentLogic(ctx, s.svcCtx)
	return l.CreatePayment(in)
}

// query transaction records based on sn
func (s *PaymentServer) GetPaymentBySn(ctx context.Context, in *pb.GetPaymentBySnReq) (*pb.GetPaymentBySnResp, error) {
	l := logic.NewGetPaymentBySnLogic(ctx, s.svcCtx)
	return l.GetPaymentBySn(in)
}

// Query transaction records based on order sn
func (s *PaymentServer) GetPaymentSuccessRefundByOrderSn(ctx context.Context, in *pb.GetPaymentSuccessRefundByOrderSnReq) (*pb.GetPaymentSuccessRefundByOrderSnResp, error) {
	l := logic.NewGetPaymentSuccessRefundByOrderSnLogic(ctx, s.svcCtx)
	return l.GetPaymentSuccessRefundByOrderSn(in)
}

// update trade state
func (s *PaymentServer) UpdateTradeState(ctx context.Context, in *pb.UpdateTradeStateReq) (*pb.UpdateTradeStateResp, error) {
	l := logic.NewUpdateTradeStateLogic(ctx, s.svcCtx)
	return l.UpdateTradeState(in)
}

// Code generated by goctl. DO NOT EDIT.
// Source: order.proto

package server

import (
	"context"

	"microservices-go-zero/app/order/cmd/rpc/internal/logic"
	"microservices-go-zero/app/order/cmd/rpc/internal/svc"
	"microservices-go-zero/app/order/cmd/rpc/pb"
)

type OrderServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedOrderServer
}

func NewOrderServer(svcCtx *svc.ServiceContext) *OrderServer {
	return &OrderServer{
		svcCtx: svcCtx,
	}
}

// create homestay order
func (s *OrderServer) CreateHomestayOrder(ctx context.Context, in *pb.CreateHomestayOrderReq) (*pb.CreateHomestayOrderResp, error) {
	l := logic.NewCreateHomestayOrderLogic(ctx, s.svcCtx)
	return l.CreateHomestayOrder(in)
}

// get homestay order detail
func (s *OrderServer) HomestayOrderDetail(ctx context.Context, in *pb.HomestayOrderDetailReq) (*pb.HomestayOrderDetailResp, error) {
	l := logic.NewHomestayOrderDetailLogic(ctx, s.svcCtx)
	return l.HomestayOrderDetail(in)
}

// update homestay order trade state
func (s *OrderServer) UpdateHomestayOrderTradeState(ctx context.Context, in *pb.UpdateHomestayOrderTradeStateReq) (*pb.UpdateHomestayOrderTradeStateResp, error) {
	l := logic.NewUpdateHomestayOrderTradeStateLogic(ctx, s.svcCtx)
	return l.UpdateHomestayOrderTradeState(in)
}

// get user homestay order list
func (s *OrderServer) UserHomestayOrderList(ctx context.Context, in *pb.UserHomestayOrderListReq) (*pb.UserHomestayOrderListResp, error) {
	l := logic.NewUserHomestayOrderListLogic(ctx, s.svcCtx)
	return l.UserHomestayOrderList(in)
}

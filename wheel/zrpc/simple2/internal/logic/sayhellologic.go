package logic

import (
	"context"

	"github.com/ct-zh/goLearn/wheel/zrpc/simple2/internal/svc"
	"github.com/ct-zh/goLearn/wheel/zrpc/simple2/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SayHelloLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSayHelloLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SayHelloLogic {
	return &SayHelloLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SayHelloLogic) SayHello(in *pb.HelloRequest) (*pb.HelloReply, error) {
	// todo: add your logic here and delete this line

	return &pb.HelloReply{}, nil
}

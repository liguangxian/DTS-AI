package logic

import (
	"context"

	"DTS-AI/dts/internal/svc"
	"DTS-AI/dts/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DtsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDtsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DtsLogic {
	return &DtsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DtsLogic) Dts(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return &types.Response{
		Message: "Hello go-zero",
	}, nil
}

package shop

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"looklook/app/shop/cmd/api/internal/svc"
)

type SyncPddGoodsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSyncPddGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SyncPddGoodsLogic {
	return &SyncPddGoodsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SyncPddGoodsLogic) SyncPddGoods() error {
	// todo: add your logic here and delete this line

	return nil
}

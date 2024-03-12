package userWonDynamicComment

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"looklook/app/comment/cmd/rpc/comment"
	"looklook/app/lottery/cmd/rpc/lottery"
	"looklook/app/usercenter/cmd/rpc/usercenter"
	"looklook/common/xerr"

	"looklook/app/usercenter/cmd/api/internal/svc"
	"looklook/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserWonDynamicCommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserWonDynamicCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserWonDynamicCommentListLogic {
	return &UserWonDynamicCommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserWonDynamicCommentListLogic) UserWonDynamicCommentList(req *types.UserWonDynamicCommentReq) (resp *types.UserWonDynamicCommentResp, err error) {
	// 获取用户动态列表
	dynamicList, err := l.svcCtx.UsercenterRpc.SearchUserDynamic(l.ctx, &usercenter.SearchUserDynamicReq{
		UserId: req.UserId,
	})
	if err != nil {
		//todo 要使用这种写法管理错误，否则Kibana无法收集到错误日志的详情
		return nil, errors.Wrapf(xerr.NewErrMsg("Failed to get SearchUserDynamic"), "Failed to get SearchUserDynamic err : %v ,req:%+v", err, req)
	}
	var userDynamicList []types.DynamicInfo
	if len(dynamicList.UserDynamic) > 0 {
		for _, item := range dynamicList.UserDynamic {
			var t types.DynamicInfo
			_ = copier.Copy(&t, item)

			userDynamicList = append(userDynamicList, t)
		}
	}

	// 获取累计奖品数量
	wonLotteryCount, err := l.svcCtx.LotteryRpc.GetWonListCount(l.ctx, &lottery.GetWonListCountReq{
		UserId: req.UserId,
	})
	if err != nil {
		//todo 要使用这种写法管理错误，否则Kibana无法收集到错误日志的详情
		return nil, errors.Wrapf(xerr.NewErrMsg("Failed to get GetWonList"), "Failed to get GetWonList err : %v ,req:%+v", err, req)
	}

	// 获取用户晒单列表
	userComment, err := l.svcCtx.CommentRpcConf.GetUserComment(l.ctx, &comment.GetUserCommentReq{
		UserId: req.UserId,
	})
	if err != nil {
		//todo 要使用这种写法管理错误，否则Kibana无法收集到错误日志的详情
		return nil, errors.Wrapf(xerr.NewErrMsg("Failed to get GetUserComment"), "Failed to get GetUserComment err : %v ,req:%+v", err, req)
	}
	// 4. 组装返回数据
	var userCommentList []types.UserCommentInfo
	// 转成map
	for _, v := range userComment.Comment {
		var t types.UserCommentInfo
		_ = copier.Copy(&t, v)
		userCommentList = append(userCommentList, t)
	}
	return &types.UserWonDynamicCommentResp{
		Count:           wonLotteryCount.Count,
		UserDynamicList: userDynamicList,
		UserCommentList: userCommentList,
	}, nil
}

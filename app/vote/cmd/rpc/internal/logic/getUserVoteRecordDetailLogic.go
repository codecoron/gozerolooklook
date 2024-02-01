package logic

import (
	"context"
	"looklook/app/vote/cmd/rpc/internal/svc"
	"looklook/app/vote/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserVoteRecordDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserVoteRecordDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserVoteRecordDetailLogic {
	return &GetUserVoteRecordDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserVoteRecordDetailLogic) GetUserVoteRecordDetail(in *pb.GetUserVoteRecordDetailReq) (*pb.GetUserVoteRecordDetailResp, error) {
	voteConfigId := in.LotteryId
	userId := in.UserId
	builder := l.svcCtx.VoteRecordModel.SelectBuilder().Where("lottery_id = ?", voteConfigId).Where("user_id = ?", userId)
	voteRecordData, err := l.svcCtx.VoteRecordModel.FindAll(l.ctx, builder, "")
	if err != nil {
		// 处理错误
		return nil, err
	}

	var userVoteRecordDetails []*pb.UserVoteRecordDetail

	// 遍历 voteRecordData，将每条记录转换为 UserVoteRecordDetail 结构
	for _, record := range voteRecordData {
		userVoteRecordDetail := &pb.UserVoteRecordDetail{
			LotteryId:      record.LotteryId,
			SelectedOption: record.SelectedOption,
		}

		userVoteRecordDetails = append(userVoteRecordDetails, userVoteRecordDetail)
	}

	//fmt.Println("----record----", userVoteRecordDetails)

	//resp := &pb.GetUserVoteRecordDetailResp{
	//	UserVoteRecordDetail: userVoteRecordDetails,
	//}

	resp := &pb.GetUserVoteRecordDetailResp{}
	resp.UserVoteRecordDetails = userVoteRecordDetails
	//if err := copier.Copy(resp.UserVoteRecordDetail, userVoteRecordDetails); err != nil {
	//	return nil, err
	//}

	//fmt.Println("----resp----", resp)

	return resp, nil
}

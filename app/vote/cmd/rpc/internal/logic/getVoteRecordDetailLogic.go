package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/copier"
	"looklook/app/vote/model"

	"looklook/app/vote/cmd/rpc/internal/svc"
	"looklook/app/vote/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVoteRecordDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVoteRecordDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVoteRecordDetailLogic {
	return &GetVoteRecordDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func MergeResults(voteConfigData *model.VoteConfig, voteRecordData []*model.VoteRecord) []*pb.VoteRecordDetail {
	var result []*pb.VoteRecordDetail

	fmt.Println("voteConfigData-------", voteConfigData)
	fmt.Println("voteConfigData.VoteConfig-------", voteConfigData.VoteConfig)

	// 检查 voteConfigData.VoteConfig 是否为空
	if !voteConfigData.VoteConfig.Valid {
		// 处理 voteConfigData.VoteConfig 为空的情况
		return nil
	}

	// 解析 voteConfigData 中的 vote_config 字段
	voteConfig := new(pb.VoteConfigJSONData)
	if err := json.Unmarshal([]byte(voteConfigData.VoteConfig.String), voteConfig); err != nil {
		// 处理错误
		return nil
	}

	fmt.Println("voteConfig.Options: ", voteConfig.Options)

	//遍历投票配置中的选项，合并投票记录
	//for _, option := range voteConfig.Options {
	//	optionText := option.Text
	//	//optionImage := option.Image
	//	fmt.Println("Text: ", optionText)
	//
	//}

	// 遍历投票配置中的选项，合并投票记录
	//for _, option := range voteConfig.Options {
	//	optionText := option.Text
	//	optionImage := option.Image
	//
	//	// 在 voteRecordData 中查找是否存在该选项
	//	// 如果存在，将投票数设置为 voteRecordData 中的值，否则设置为 0
	//	voteCount := 0
	//	for _, recordOption := range voteRecordData.Options {
	//		if recordOption.Text == optionText && recordOption.Image == optionImage {
	//			voteCount = recordOption.VoteCount
	//			break
	//		}
	//	}
	//
	//	// 构造结果并添加到 result 中
	//	result = append(result, &pb.VoteRecordDetail{
	//		OptionText:     optionText,
	//		SelectedOption: option.SelectedOption,
	//		OptionImage:    optionImage,
	//		VoteCount:      int64(voteCount),
	//	})
	//}

	return result
}

//func MergeResults1(voteConfigData *model.VoteConfig, voteRecordData *model.VoteRecord) []*pb.VoteRecordDetail {
//	var result []*pb.VoteRecordDetail
//
//	// 解析 voteConfigData 中的 vote_config 字段
//	voteConfig := &pb.VoteConfig{}
//	if err := json.Unmarshal([]byte(voteConfigData.VoteConfig), voteConfig); err != nil {
//		// 处理错误
//		return nil
//	}
//
//	// 遍历投票配置中的选项，合并投票记录
//	for _, option := range voteConfig.Options {
//		optionText := option.Text
//		optionImage := option.Image
//
//		// 在 voteRecordData 中查找是否存在该选项
//		// 如果存在，将投票数设置为 voteRecordData 中的值，否则设置为0
//		voteCount := 0
//		for _, recordOption := range voteRecordData.Options {
//			if recordOption.Text == optionText && recordOption.Image == optionImage {
//				voteCount = recordOption.VoteCount
//				break
//			}
//		}
//
//		// 构造结果并添加到 result 中
//		result = append(result, &pb.VoteRecordDetail{
//			OptionText:     optionText,
//			SelectedOption: option.SelectedOption,
//			OptionImage:    optionImage,
//			VoteCount:      int64(voteCount),
//		})
//	}
//
//	return result
//}

func (l *GetVoteRecordDetailLogic) GetVoteRecordDetail(in *pb.GetVoteRecordDetailReq) (*pb.GetVoteRecordDetailResp, error) {
	voteConfigId := in.LotteryId
	//1.查询配置数据x
	voteConfigData, err := l.svcCtx.VoteConfigModel.FindOne(l.ctx, voteConfigId)
	if err != nil {
		return nil, err
	}
	voteConfig := &pb.VoteConfig{}
	if err := copier.Copy(voteConfig, voteConfigData); err != nil {
		return nil, err
	}
	fmt.Println("----config----", voteConfig)

	//2.查询投票数据y
	//voteRecordData, err := l.svcCtx.VoteRecordModel.FindOne(l.ctx, voteConfigId)

	//builder := squirrel.Select("*").From("vote_record").Where(squirrel.Eq{"lottery_id": voteConfigId})
	builder := l.svcCtx.VoteRecordModel.SelectBuilder().Where("lottery_id = ?", voteConfigId)
	voteRecordData, err := l.svcCtx.VoteRecordModel.FindAll(l.ctx, builder, "")
	if err != nil {
		// 处理错误
		return nil, err
	}
	voteRecord := &pb.VoteRecord{}
	if err := copier.Copy(voteRecord, voteRecordData); err != nil {
		return nil, err
	}
	fmt.Println("----record----", voteRecord)

	// 3.在程序中将两个结果合并，将未出现的选项投票数设置为0
	result := MergeResults(voteConfigData, voteRecordData)

	fmt.Println(result)

	return &pb.GetVoteRecordDetailResp{}, nil
}

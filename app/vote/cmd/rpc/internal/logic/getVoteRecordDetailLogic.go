package logic

import (
	"context"
	"encoding/json"
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

func MergeResults(voteConfigData *model.VoteConfig, voteRecordData []*model.VoteRecord) *pb.GetVoteRecordDetailResp {
	resp := &pb.GetVoteRecordDetailResp{}

	// 检查 voteConfigData.VoteConfig 是否为空
	if !voteConfigData.VoteConfig.Valid {
		// 处理 voteConfigData.VoteConfig 为空的情况
		return resp
	}

	// 解析 voteConfigData 中的 vote_config 字段
	voteConfig := new(pb.VoteConfigJSONData)
	if err := json.Unmarshal([]byte(voteConfigData.VoteConfig.String), voteConfig); err != nil {
		// 处理错误
		return resp
	}

	// 构建投票记录的 map，以选项的索引作为键
	voteRecordMap := make(map[int]int)
	for _, recordOption := range voteRecordData {
		voteRecordMap[int(recordOption.SelectedOption)] += 1
	}

	// 遍历投票配置中的选项，合并投票记录
	var voteRecordDetails []*pb.VoteRecordDetail
	for index, option := range voteConfig.Options {
		optionText := option.Text
		optionImage := option.Image
		//fmt.Println("Text: ", optionText)
		//fmt.Println("Image: ", optionImage)
		//fmt.Println("index: ", index)

		// 从投票记录的 map 中查找对应索引的累计投票数
		voteCount := voteRecordMap[index]

		// 构造投票选项的详细信息并添加到切片中
		voteRecordDetails = append(voteRecordDetails, &pb.VoteRecordDetail{
			OptionText:     optionText,
			SelectedOption: int64(index),
			OptionImage:    optionImage,
			VoteCount:      int64(voteCount),
		})
	}

	// 将投票选项的详细信息切片赋值给 resp
	resp.VoteRecordDetails = voteRecordDetails

	return resp
}

func (l *GetVoteRecordDetailLogic) GetVoteRecordDetail(in *pb.GetVoteRecordDetailReq) (resp *pb.GetVoteRecordDetailResp, err error) {
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
	//fmt.Println("----config----", voteConfig)

	//2.查询投票数据y
	builder := l.svcCtx.VoteRecordModel.SelectBuilder().Where("lottery_id = ?", voteConfigId)
	voteRecordData, err := l.svcCtx.VoteRecordModel.FindAll(l.ctx, builder, "")
	if err != nil {
		// 处理错误
		return nil, err
	}
	voteRecord := new([]pb.VoteRecord)
	if err := copier.Copy(voteRecord, voteRecordData); err != nil {
		return nil, err
	}
	//fmt.Println("----record----", voteRecord)

	// 3.在程序中将两个结果合并，将未出现的选项投票数设置为0
	result := MergeResults(voteConfigData, voteRecordData)

	return result, nil
}

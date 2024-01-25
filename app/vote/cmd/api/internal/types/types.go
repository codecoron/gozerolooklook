// Code generated by goctl. DO NOT EDIT.
package types

type CreateVoteReq struct {
	UserId     int64              `json:"userId"`     //用户id
	LotteryId  int64              `json:"lotteryId"`  //抽奖id
	EnableVote int64              `json:"enableVote"` //是否启用投票功能 1是 0否
	VoteConfig VoteConfigJSONData `json:"voteConfig"` //投票配置字段说明
}

type CreateVoteResp struct {
	Id int64 `json:"id"`
}

type VoteConfigJSONData struct {
	Title           string       `json:"title"`           //投票标题
	Description     string       `json:"description"`     //投票描述【非必填】
	WinnerSelection int64        `json:"winnerSelection"` //中奖者设置：1从所有投票者中抽取 2从票数最多的一方中抽取
	Type            int64        `json:"type"`            //投票类型：1单选 2多选
	MinVotes        int64        `json:"minVotes"`        //最小投票范围
	MaxVotes        int64        `json:"maxVotes"`        //最大投票范围
	Options         []VoteOption `json:"options"`         //选项列表
}

type VoteOption struct {
	Text  string `json:"text"`  //选项名称
	Image string `json:"image"` //选项图片
}

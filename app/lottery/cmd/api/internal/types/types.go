// Code generated by goctl. DO NOT EDIT.
package types

type Lottery struct {
	Id            int64  `json:"id"`
	UserId        int64  `json:"userId"`        //发起抽奖用户ID
	Name          string `json:"name"`          //默认一等奖名称
	Thumb         string `json:"thumb"`         //默认一等奖配图
	PublishType   int64  `json:"publishType"`   //开奖设置：1按时间开奖 2按人数开奖 3即抽即中
	IsSelected    int64  `json:"isSelected"`    //是否精选 1是 0否
	PublishTime   int64  `json:"publish_time"`  //开奖时间
	JoinNumber    int64  `json:"join_number"`   //自动开奖人数标准
	Introduce     string `json:"introduce"`     //抽奖说明
	AwardDeadline int64  `json:"awardDeadline"` //领奖截止时间
}

type LotteryListReq struct {
	LastId     int64 `json:"lastId"`
	Page       int64 `json:"page"`
	PageSize   int64 `json:"pageSize"`
	IsSelected int64 `json:"isSelected"`
}

type LotteryListResp struct {
	List []Lottery `json:"list"`
}

type CreateLotteryResp struct {
	Id int64 `json:"id"`
}
type Prize struct {
	Id        int64 `json:"id"`
	LotteryId int64 `json:"lotteryId"` //抽奖ID
	CreatePrize
}

type CreatePrize struct {
	Type      int64  `json:"type"`      //奖品类型：1奖品 2优惠券 3兑换码 4商城 5微信红包封面 6红包
	Name      string `json:"name"`      //奖品名称
	Count     int64  `json:"count"`     //奖品份数
	Thumb     string `json:"thumb"`     //默认一等奖配图
	GrantType int64  `json:"grantType"` //奖品发放方式：1快递邮寄 2让中奖者联系我 3中奖者填写信息 4跳转到其他小程序
}

type CreateLotteryReq struct {
	Name          string         `json:"name"`          //默认一等奖名称
	Thumb         string         `json:"thumb"`         //默认一等奖配图
	PublishType   int64          `json:"publishType"`   //开奖设置：1按时间开奖 2按人数开奖 3即抽即中
	IsSelected    int64          `json:"isSelected"`    //是否精选 1是 0否
	PublishTime   int64          `json:"publish_time"`  //开奖时间
	JoinNumber    int64          `json:"join_number"`   //自动开奖人数标准
	Introduce     string         `json:"introduce"`     //抽奖说明
	AwardDeadline int64          `json:"awardDeadline"` //领奖截止时间
	Prizes        []*CreatePrize `json:"prizes"`        //奖品 支持多个
}

type SetLotteryIsSelectedReq struct {
	Id int64 `json:"id"`
}

type SetLotteryIsSelectedResp struct {
	IsSelected int64 `json:"isSelected"`
}

type LotteryDetail struct {
	Id            int64          `json:"id"`
	Name          string         `json:"name"`
	UserId        int64          `json:"userId"`        //发起抽奖用户ID
	Prizes        []*CreatePrize `json:"prizes"`        //奖品信息
	PublishType   int64          `json:"publishType"`   //开奖设置：1按时间开奖 2按人数开奖 3即抽即中
	IsSelected    int64          `json:"isSelected"`    //是否精选 1是 0否
	PublishTime   int64          `json:"publish_time"`  //开奖时间
	JoinNumber    int64          `json:"join_number"`   //自动开奖人数标准
	Introduce     string         `json:"introduce"`     //抽奖说明
	AwardDeadline int64          `json:"awardDeadline"` //领奖截止时间
}

type LotteryDetailReq struct {
	Id int64 `json:"id"`
}

type LotteryDetailResp struct {
	LotteryDetail LotteryDetail `json:"lotteryDetail"`
}

type Sponsor struct {
	Id       int64  `json:"id"`
	NickName string `json:"nickName"`
	Avatar   string `json:"avatar"`
	Info     string `json:"info"`
}

type LotterySponsorReq struct {
	Id int64 `json:"id"`
}

type LotterySponsorResp struct {
	Sponsor Sponsor `json:"sponsor"`
}

type UpdateLotteryReq struct {
	Id int64 `json:"id"`
}

type UpdateLotteryResp struct {
}

type AddLotteryParticipationReq struct {
	UserId    int64 `json:"userId"`
	LotteryId int64 `json:"lotteryId"`
}

type AddLotteryParticipationResp struct {
	Id int64 `json:"id"`
}

type SearchLotteryParticipationReq struct {
	LotteryId int64 `json:"lotteryId"`
	IsAll     bool  `json:"isAll"`
}

type SearchLotteryParticipationResp struct {
	Count int64    `json:"count"`
	List  []string `json:"list"`
}

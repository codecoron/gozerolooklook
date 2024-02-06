// Code generated by goctl. DO NOT EDIT.
package types

type Lottery struct {
	Id            int64  `json:"id"`
	UserId        int64  `json:"userId"`        //发起抽奖用户ID
	Name          string `json:"name"`          //默认一等奖名称
	Thumb         string `json:"thumb"`         //默认一等奖配图
	PublishTime   int64  `json:"publishTime"`   //发布抽奖时间
	JoinNumber    int64  `json:"joinNumber"`    //自动开奖人数标准
	Introduce     string `json:"introduce"`     //抽奖说明
	AwardDeadline int64  `json:"awardDeadline"` //领奖截止时间
	IsSelected    int64  `json:"isSelected"`    //是否精选 1是 0否
	AnnounceType  int64  `json:"announceType"`  //开奖设置：1按时间开奖 2按人数开奖 3即抽即中
	AnnounceTime  int64  `json:"announceTime"`  //开奖时间
	IsAnnounced   int64  `json:"isAnnounced"`   // 是否已经开奖：0未开奖 1已开奖
	SponsorId     int64  `json:"sponsorId"`     // 赞助商Id
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
	AnnounceType  int64          `json:"announceType"`  //开奖设置：1按时间开奖 2按人数开奖 3即抽即中
	AnnounceTime  int64          `json:"announceTime"`  //开奖时间
	JoinNumber    int64          `json:"joinNumber"`    //自动开奖人数标准
	Introduce     string         `json:"introduce"`     //抽奖说明
	AwardDeadline int64          `json:"awardDeadline"` //领奖截止时间
	SponsorId     int64          `json:"sponsorId"`     // 赞助商Id
	Prizes        []*CreatePrize `json:"prizes"`        //奖品 支持多个
}

type CreateLotteryResp struct {
	Id int64 `json:"id"`
}

type UpdateLotteryReq struct {
	Id int64 `json:"id"`
}

type UpdateLotteryResp struct {
}

type SetLotteryIsSelectedReq struct {
	Id int64 `json:"id"`
}

type SetLotteryIsSelectedResp struct {
	IsSelected int64 `json:"isSelected"`
}

type LotterySponsor struct {
	Id       int64  `json:"id"`       // 赞助商Id
	NickName string `json:"nickName"` // 赞助商昵称
	Avatar   string `json:"avatar"`   // 赞助商头像
	Info     string `json:"info"`     // 赞助商信息
}

type LotteryDetailReq struct {
	Id int64 `json:"id"`
}

type LotteryDetailResp struct {
	Id            int64           `json:"id"`
	UserId        int64           `json:"userId"`        //发起抽奖用户ID
	Name          string          `json:"name"`          //默认一等奖名称
	Thumb         string          `json:"thumb"`         //默认一等奖配图
	PublishTime   int64           `json:"publishTime"`   //发布抽奖时间
	JoinNumber    int64           `json:"joinNumber"`    //自动开奖人数标准
	Introduce     string          `json:"introduce"`     //抽奖说明
	AwardDeadline int64           `json:"awardDeadline"` //领奖截止时间
	IsSelected    int64           `json:"isSelected"`    //是否精选 1是 0否
	AnnounceType  int64           `json:"announceType"`  //开奖设置：1按时间开奖 2按人数开奖 3即抽即中
	AnnounceTime  int64           `json:"announceTime"`  //开奖时间
	IsAnnounced   int64           `json:"isAnnounced"`   // 是否已经开奖：0未开奖 1已开奖
	SponsorId     int64           `json:"sponsorId"`     // 赞助商Id
	Prizes        []*CreatePrize  `json:"prizes"`        //奖品信息
	Sponsor       *LotterySponsor `json:"sponsor"`       // 抽奖赞助商信息
}

type CheckIsParticipatedReq struct {
	LotteryId int64 `json:"lotteryId"` // 当前抽奖Id
}

type CheckIsParticipatedResp struct {
	IsParticipated int64 `json:"isParticipated"` // 当前用户是否参与了当前抽奖： 0未参与 1已参与
}

type TestReq struct {
	Age        int64  `json:"age" validate:"gte=1,lte=130"`
	Name       string `json:"name" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required"`
	RePassword string `json:"re_password" validate:"required"`
	Date       string `json:"date" validate:"required,datetime=2006-01-02,checkDate"`
}

type TestResp struct {
}

type AddLotteryParticipationReq struct {
	LotteryId int64 `json:"lotteryId"`
}

type AddLotteryParticipationResp struct {
}

type SearchLotteryParticipationReq struct {
	LotteryId int64 `json:"lotteryId"`
	PageIndex int64 `json:"pageIndex"`
	PageSize  int64 `json:"pageSize"`
}

type UserInfo struct {
	Mobile       string `json:"mobile"`
	Nickname     string `json:"nickname"`
	Sex          int64  `json:"sex"`
	Avatar       string `json:"avatar"`
	Info         string `json:"info"`
	Signature    string `json:"signature"`
	LocationName string `json:"locationName"`
}

type SearchLotteryParticipationResp struct {
	Count int64       `json:"count"`
	List  []*UserInfo `json:"list"`
}

// Code generated by goctl. DO NOT EDIT.
package types

type AddLotteryParticipationReq struct {
	LotteryId int64 `json:"lotteryId"`
}

type AddLotteryParticipationResp struct {
}

type CheckIsParticipatedReq struct {
	LotteryId int64 `json:"lotteryId"` // 当前抽奖Id
}

type CheckIsParticipatedResp struct {
	IsParticipated int64 `json:"isParticipated"` // 当前用户是否参与了当前抽奖： 0未参与 1已参与
}

type ClockTaskTypeListReq struct {
}

type ClockTaskTypeListResp struct {
	List []CockTaskType `json:"list"`
}

type CockTaskType struct {
	Type    int64  `json:"type"`
	Text    string `json:"text"`
	Seconds int64  `json:"seconds"`
}

type CreateClockTask struct {
	Type             int64  `json:"type"`                       // 任务类型 1: 体验小程序 2： 浏览指定公众号文章 3: 浏览图片（微信图片二维码等） 4： 浏览视频号视频
	Seconds          int64  `json:"seconds"`                    // 任务秒数
	AppletType       int64  `json:"appletType, optional"`       // type=1时该字段才有意义 小程序跳转类型，1小程序链接 2小程序路径
	PageLink         string `json:"pageLink, optional"`         // type=1 并且 applet_type=1时 该字段才有意义 配置要跳转小程序的页面链接 （如 #小程序://抽奖/oM....）
	AppId            string `json:"appId, optional"`            // type=1 并且 applet_type=2时 该字段才有意义 配置要跳转的小程序AppID
	PagePath         string `json:"pagePath, optional"`         // type=1 并且 applet_type=2时 该字段才有意义 配置要跳转的小程序路径（如：/pages/index）
	Image            string `json:"image, optional"`            // type=3时 该字段才有意义 添加要查看的图片
	VideoAccountId   string `json:"videoAccountId, optional"`   // type=4时 该字段才有意义 视频号ID
	VideoId          string `json:"videoId, optional"`          // type=4时 该字段才有意义 视频ID
	ArticleLink      string `json:"articleLink, optional"`      // type=2时 该字段才有意义 公众号文章链接
	Copywriting      string `json:"copywriting"`                // 引导参与者完成打卡任务的文案
	ChanceType       int64  `json:"chanceType"`                 // 概率类型 1: 随机 2: 指定
	IncreaseMultiple int64  `json:"increaseMultiple, optional"` // ChanceType=2时 该字段才有意义，概率增加倍数
}

type CreateLotteryReq struct {
	Name          string           `json:"name"`                //默认一等奖名称
	Thumb         string           `json:"thumb"`               //默认一等奖配图
	AnnounceType  int64            `json:"announceType"`        //开奖设置：1按时间开奖 2按人数开奖 3即抽即中
	AnnounceTime  int64            `json:"announceTime"`        //开奖时间
	JoinNumber    int64            `json:"joinNumber"`          //自动开奖人数标准
	Introduce     string           `json:"introduce"`           //抽奖说明
	AwardDeadline int64            `json:"awardDeadline"`       //领奖截止时间
	SponsorId     int64            `json:"sponsorId"`           // 赞助商Id
	Prizes        []*CreatePrize   `json:"prizes"`              //奖品 支持多个
	IsClocked     int64            `json:"isClocked"`           //是否开启打卡任务 0未开启；1已开启
	ClockTask     *CreateClockTask `json:"clockTask, optional"` //打卡任务 支持一个
}

type CreateLotteryResp struct {
	Id int64 `json:"id"`
}

type CreatePrize struct {
	Type      int64  `json:"type"`      //奖品类型：1奖品 2优惠券 3兑换码 4商城 5微信红包封面 6红包
	Name      string `json:"name"`      //奖品名称
	Count     int64  `json:"count"`     //奖品份数
	Thumb     string `json:"thumb"`     //默认一等奖配图
	GrantType int64  `json:"grantType"` //奖品发放方式：1快递邮寄 2让中奖者联系我 3中奖者填写信息 4跳转到其他小程序
}

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

type LotteryListReq struct {
	LastId     int64 `json:"lastId"`
	Page       int64 `json:"page"`
	PageSize   int64 `json:"pageSize"`
	IsSelected int64 `json:"isSelected"`
}

type LotteryListResp struct {
	List []Lottery `json:"list"`
}

type LotterySponsor struct {
	Id         int64  `json:"id"`         //id
	UserId     int64  `json:"userId"`     //userId
	Type       int64  `json:"type"`       //1微信号 2公众号 3小程序 4微信群 5视频号
	AppletType int64  `json:"appletType"` //type=3时该字段才有意义，1小程序链接 2路径跳转 3二维码跳转
	Name       string `json:"name"`       //名称
	Desc       string `json:"desc"`       //描述
	Avatar     string `json:"avatar"`     //avatar
	IsShow     int64  `json:"isShow"`     //1显示 2不显示
	QrCode     string `json:"qrCode"`     //二维码图片地址, type=1 2 3&applet_type=3 4的时候启用
	InputA     string `json:"inputA"`     //type=5 applet_type=2 or applet_type=1 输入框A
	InputB     string `json:"inputB"`     //type=5 applet_type=2输入框B
}

type Prize struct {
	Id        int64 `json:"id"`
	LotteryId int64 `json:"lotteryId"` //抽奖ID
	CreatePrize
}

type SearchLotteryParticipationReq struct {
	LotteryId int64 `json:"lotteryId"`
	PageIndex int64 `json:"pageIndex"`
	PageSize  int64 `json:"pageSize"`
}

type SearchLotteryParticipationResp struct {
	Count int64       `json:"count"`
	List  []*UserInfo `json:"list"`
}

type SetLotteryIsSelectedReq struct {
	Id int64 `json:"id"`
}

type SetLotteryIsSelectedResp struct {
	IsSelected int64 `json:"isSelected"`
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

type UpdateLotteryReq struct {
	Id int64 `json:"id"`
}

type UpdateLotteryResp struct {
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

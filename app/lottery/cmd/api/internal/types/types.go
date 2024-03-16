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
	Level     int64  `json:"level"`     //奖品等级 1一等奖 2二等奖 3三等奖，依次类推
	GrantType int64  `json:"grantType"` //奖品发放方式：1快递邮寄 2让中奖者联系我 3中奖者填写信息 4跳转到其他小程序
}

type CreateClockTask struct {
	Type             int64  `json:"type, optional"`             // 任务类型 1: 体验小程序 2： 浏览指定公众号文章 3: 浏览图片（微信图片二维码等） 4： 浏览视频号视频
	Seconds          int64  `json:"seconds, optional"`          // 任务秒数
	AppletType       int64  `json:"appletType, optional"`       // type=1时该字段才有意义 小程序跳转类型，1小程序链接 2小程序路径
	PageLink         string `json:"pageLink, optional"`         // type=1 并且 applet_type=1时 该字段才有意义 配置要跳转小程序的页面链接 （如 #小程序://抽奖/oM....）
	AppId            string `json:"appId, optional"`            // type=1 并且 applet_type=2时 该字段才有意义 配置要跳转的小程序AppID
	PagePath         string `json:"pagePath, optional"`         // type=1 并且 applet_type=2时 该字段才有意义 配置要跳转的小程序路径（如：/pages/index）
	Image            string `json:"image, optional"`            // type=3时 该字段才有意义 添加要查看的图片
	VideoAccountId   string `json:"videoAccountId, optional"`   // type=4时 该字段才有意义 视频号ID
	VideoId          string `json:"videoId, optional"`          // type=4时 该字段才有意义 视频ID
	ArticleLink      string `json:"articleLink, optional"`      // type=2时 该字段才有意义 公众号文章链接
	Copywriting      string `json:"copywriting,optional"`       // 引导参与者完成打卡任务的文案
	ChanceType       int64  `json:"chanceType,optional"`        // 概率类型 1: 随机 2: 指定
	IncreaseMultiple int64  `json:"increaseMultiple, optional"` // ChanceType=2时 该字段才有意义，概率增加倍数
}

type CreateLotteryReq struct {
	Name          string           `json:"name"`                              //默认一等奖名称
	Thumb         string           `json:"thumb"`                             //默认一等奖配图
	AnnounceType  int64            `json:"announceType" validate:"oneof=1 2"` //开奖设置：1按时间开奖 2按人数开奖 3即抽即中
	AnnounceTime  int64            `json:"announceTime"`                      //开奖时间
	JoinNumber    int64            `json:"joinNumber"`                        //自动开奖人数标准
	Introduce     string           `json:"introduce"`                         //抽奖说明
	AwardDeadline int64            `json:"awardDeadline"`                     //领奖截止时间
	SponsorId     int64            `json:"sponsorId"`                         // 赞助商Id
	Prizes        []*CreatePrize   `json:"prizes"`                            //奖品 支持多个
	IsClocked     int64            `json:"isClocked"`                         //是否开启打卡任务 0未开启；1已开启
	ClockTask     *CreateClockTask `json:"clockTask, optional"`               //打卡任务 支持一个
	PublishType   int64            `json:"publishType" validate:"oneof=1 2"`  //发布类型 1发布抽奖 2发布测试
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
	Id         int64  `json:"id"`              //id
	UserId     int64  `json:"userId"`          //userId
	Type       int64  `json:"type"`            //1微信号 2公众号 3小程序 4微信群 5视频号
	AppletType int64  `json:"appletType"`      //type=3时该字段才有意义，1小程序链接 2路径跳转 3二维码跳转
	Name       string `json:"name"`            //名称
	Desc       string `json:"desc"`            //描述
	Avatar     string `json:"avatar"`          //avatar
	IsShow     int64  `json:"isShow"`          //1显示 2不显示
	QrCode     string `json:"qrCode,optional"` //二维码图片地址, type=1 2 3&applet_type=3 4的时候启用
	InputA     string `json:"inputA,optional"` //type=5 applet_type=2 or applet_type=1 输入框A
	InputB     string `json:"inputB,optional"` //type=5 applet_type=2输入框B
}

type LotteryDetailReq struct {
	Id int64 `json:"id"`
}

type LotteryDetailResp struct {
	Id             int64           `json:"id"`
	UserId         int64           `json:"userId"`         //发起抽奖用户ID
	Name           string          `json:"name"`           //默认一等奖名称
	Thumb          string          `json:"thumb"`          //默认一等奖配图
	PublishTime    int64           `json:"publishTime"`    //发布抽奖时间
	JoinNumber     int64           `json:"joinNumber"`     //自动开奖人数标准
	Introduce      string          `json:"introduce"`      //抽奖说明
	AwardDeadline  int64           `json:"awardDeadline"`  //领奖截止时间
	IsSelected     int64           `json:"isSelected"`     //是否精选 1是 0否
	AnnounceType   int64           `json:"announceType"`   //开奖设置：1按时间开奖 2按人数开奖 3即抽即中
	AnnounceTime   int64           `json:"announceTime"`   //开奖时间
	IsAnnounced    int64           `json:"isAnnounced"`    // 是否已经开奖：0未开奖 1已开奖
	SponsorId      int64           `json:"sponsorId"`      // 赞助商Id
	Prizes         []*CreatePrize  `json:"prizes"`         //奖品信息
	Sponsor        *LotterySponsor `json:"sponsor"`        // 抽奖赞助商信息
	IsParticipated int64           `json:"isParticipated"` // 当前用户是否参与了当前抽奖： 0未参与 1已参与
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

type CockTaskType struct {
	Type    int64  `json:"type"`
	Text    string `json:"text"`
	Seconds int64  `json:"seconds"`
}

type ClockTaskTypeListReq struct {
}

type ClockTaskTypeListResp struct {
	List []CockTaskType `json:"list"`
}

type ChanceType struct {
	Type int64  `json:"type"`
	Text string `json:"text"`
}

type ChanceTypeListReq struct {
}

type ChanceTypeListResp struct {
	List []ChanceType `json:"list"`
}

type CreateClockTaskRecordReq struct {
	LotteryId   int64 `json:"lotteryId"`
	ClockTaskId int64 `json:"clockTaskId"`
}

type CreateClockTaskRecordResp struct {
	Id int64 `json:"id"`
}

type GetCreateLotteryListByUserIdReq struct {
	LastId   int64 `json:"lastId"`
	PageSize int64 `json:"pageSize"`
}

type GetCreateLotteryListByUserIdResp struct {
	List []Prize `json:"list"`
}

type LotteryParticipation struct {
	Id        int64 `json:"id"`         // 主键
	LotteryId int64 `json:"lottery_id"` // 参与的抽奖的id
	UserId    int64 `json:"user_id"`    // 用户id
	IsWon     int64 `json:"is_won"`     // 中奖了吗？
	PrizeId   int64 `json:"prize_id"`   // 中奖id
}

type Prizes struct {
	Id        int64  `json:"id"`
	LotteryId int64  `json:"lottery_id"` // 抽奖ID
	Type      int64  `json:"type"`       // 奖品类型：1奖品 2优惠券 3兑换码 4商城 5微信红包封面 6红包
	Name      string `json:"name"`       // 奖品名称
	Level     int64  `json:"level"`      // 几等奖 默认1
	Thumb     string `json:"thumb"`      // 奖品图
	Count     int64  `json:"count"`      // 奖品份数
	GrantType int64  `json:"grant_type"` // 奖品发放方式：1快递邮寄 2让中奖者联系我 3中奖者填写信息 4跳转到其他小程序
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

type WonList struct {
	Id        int64   `json:"id"`         // 主键
	LotteryId int64   `json:"lottery_id"` // 参与的抽奖的id
	UserId    int64   `json:"user_id"`    // 用户id
	IsWon     int64   `json:"is_won"`     // 中奖了吗？
	Prize     *Prizes `json:"prize"`      // 中奖奖品
}

type GetLotteryWinListReq struct {
	LastId int64 `json:"lastId"`
	Size   int64 `json:"size"`
}

type GetLotteryWinListResp struct {
	List []*WonList `json:"list"`
}

type CheckIsWinReq struct {
	LotteryId int64 `json:"lotteryId"`
}

type CheckIsWinResp struct {
	IsWon int64 `json:"isWon"`
}

type GetLotteryWinList2Req struct {
	LotteryId int64 `json:"lotteryId"`
}

type WonList2 struct {
	Prize       *Prizes     `json:"prize"`
	WinnerCount int64       `json:"winnerCount"`
	Users       []*UserInfo `json:"users"`
}

type GetLotteryWinList2Resp struct {
	List []*WonList2 `json:"list"`
}

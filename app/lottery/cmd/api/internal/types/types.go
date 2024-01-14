// Code generated by goctl. DO NOT EDIT.
package types

type Lottery struct {
	Id            int64  `json:"id"`
	UserId        int64  `json:"userId"`
	Name          string `json:"name"`          //发起抽奖用户ID
	Thumb         string `json:"thumb"`         //默认取一等奖名称
	PublishType   int8   `json:"publishType"`   //开奖设置：1按时间开奖 2按人数开奖 3即抽即中
	IsSelected    int8   `json:"isSelected"`    //是否精选 1是 0否
	PublishTime   string `json:"publish_time"`  //开奖时间
	JoinNumber    int64  `json:"join_number"`   //自动开奖人数标准
	Introduce     string `json:"introduce"`     //抽奖说明
	AwardDeadline string `json:"awardDeadline"` //领奖截止时间
}

type LotteryListReq struct {
	LastId     int64 `json:"lastId"`
	Page       int64 `json:"page"`
	PageSize   int64 `json:"pageSize"`
	IsSelected int8  `json:"isSelected"`
}

type LotteryListResp struct {
	List []Lottery `json:"list"`
}

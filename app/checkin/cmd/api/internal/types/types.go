// Code generated by goctl. DO NOT EDIT.
package types

type CheckinRecord struct {
	Id                    int64 `json:"id"`
	UserId                int64 `json:"userId"`                //用户ID
	ContinuousCheckinDays int64 `json:"continuousCheckinDays"` //用户连续签到的天数
	LastCheckinDate       int64 `json:"lastCheckinDate"`       //用户最后一次签到
	State                 int64 `json:"state"`                 //当天用户是否签到，0为未签，1为已签
}

type CheckinReq struct {
}

type CheckinResp struct {
	ContinuousCheckinDays int64 `json:"continuousCheckinDays"` //用户连续签到的天数
	State                 int64 `json:"state"`                 //当天用户是否签到，0为未签，1为已签
	Integral              int64 `json:"integral"`              //心愿值
}

type ClaimRewardReq struct {
	TaskId int64 `json:"taskId"`
}

type ClaimRewardResp struct {
}

type GetCheckinReq struct {
}

type GetCheckinResp struct {
	ContinuousCheckinDays int64 `json:"continuousCheckinDays"` //用户连续签到的天数
	State                 int64 `json:"state"`                 //当天用户是否签到，0为未签，1为已签
	Integral              int64 `json:"integral"`              //心愿值
}

type GetTasksReq struct {
}

type GetTasksResp struct {
	TasksList []*Tasks `json:"tasksList"`
	DayCount  int64    `json:"dayCount"`
	WeekCount int64    `json:"weekCount"`
}

type IntegralRecord struct {
	Id       int64  `json:"id"`
	UserId   int64  `json:"userId"`   //用户ID
	Integral int64  `json:"integral"` //增加或减少积分的数量
	Content  string `json:"content"`  //增加或减少积分的原因
}

type Intergral struct {
	Id       int64 `json:"id"`
	UserId   int64 `json:"userId"`   //用户ID
	Integral int64 `json:"integral"` //心愿值
}

type TaskProgress struct {
	Id                    int64 `json:"id"`
	UserId                int64 `json:"userId"`
	IsParticipatedLottery int64 `json:"isParticipatedLottery"`
	IsInitiatedLottery    int64 `json:"isInitiatedLottery"`
	IsSubCheckin          int64 `json:"isSubCheckin"`
}

type Tasks struct {
	Id         int64  `json:"id"`
	Type       int64  `json:"type"`
	Content    string `json:"content"`
	Integral   int64  `json:"integral"`
	IsFinished int64  `json:"isFinished"`
}

type UpdateSubReq struct {
	State int64 `json:"state"`
}

type UpdateSubResp struct {
}

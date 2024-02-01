package xerr

var message map[uint32]string

func init() {
	message = make(map[uint32]string)
	message[OK] = "SUCCESS"
	message[SERVER_COMMON_ERROR] = "服务器开小差啦,稍后再来试一试"
	message[REUQEST_PARAM_ERROR] = "参数错误"
	message[TOKEN_EXPIRE_ERROR] = "token失效，请重新登陆"
	message[TOKEN_GENERATE_ERROR] = "生成token失败"
	message[DB_ERROR] = "数据库繁忙,请稍后再试"
	message[DB_UPDATE_AFFECTED_ZERO_ERROR] = "更新数据影响行数为0"
	message[VOTE_VOTE_CONFIG_NOT_FOUND] = "投票配置有问题,请检查配置" //投票配置通用错误,内容可根据业务来定义
	message[VOTE_VOTE_RECORD_NOT_FOUND] = "投票记录有问题,请检查配置" //投票记录通用错误,内容可根据业务来定义
	//签到模块
	message[VOTE_VOTE_CONFIG_NOT_FOUND] = "投票信息不存在"
	message[CHECKIN_RECORD_NOT_FOUND] = "签到信息不存在"
	message[CHECKIN_REPEAT] = "今日已签到"
	message[CHECKIN_TASK_NOT_FOUND] = "任务不存在"
	message[CHECKIN_TASK_REWARD_COLLECTED] = "不可重复领取奖励"
	// 抽奖模块
	// 开奖
	message[DB_FIND_PARTICIPATOR_ERROR] = "获取当前用户参与抽奖信息失败"
	message[GET_PARTICIPATION_USERIDS_BYLOTTERYID_ERROR] = "获取当前抽奖所有参与者Id失败"
	message[AnnounceLottery_ERROR] = "开奖策略运行失败"
	message[UPDATE_WINNER_ERROR] = "更新中奖者信息失败"
	message[GETLOTTERY_BYLESSTHAN_CURRENTTIME_ERROR] = "根据小于当前时间获取抽奖信息失败"
	message[UPDATE_LOTTERY_STATUS_ERROR] = "根据lotteryId更新lottery状态为已开奖失败"
	message[GET_TYPEIS2_AND_ISNOT_ANNOUNCE_LOTTERYS_ERROR] = "根据开奖类型为2获取所有未开奖的抽奖失败"
	message[GET_PARTICIPATORS_COUNT_BYLOTTERYID_ERROR] = "根据抽奖id获取参与者总数失败"
	// lotterydetail
	message[DB_FIND_PRIZES_BYLOTTERYID_ERROR] = "根据抽奖Id获取奖品列表失败"
	message[DB_FIND_LOTTERY_BYLOTTERYID_ERROR] = "根据抽奖Id获取抽奖失败"
	message[DB_FIND_USERID_BYLOTTERYID_ERROR] = "根据抽奖Id获取赞助商Id失败"
	message[DB_USERID_NOTFOUND] = "当前抽奖没有赞助商Id"
	// 发起抽奖
	message[DB_INSERTLOTTERY_ERROR] = "发起抽奖插入抽奖信息失败"
	message[DB_INSERTPRIZE_ERROR] = "插入奖品信息失败"
	// 抽奖列表
	message[DB_GET_LOTTERY_LIST_ERROR] = "获取抽奖列表失败"
	// 检验当前用户是否发起过抽奖
	message[DB_GET_LOTTERYID_BYUSERID_ERROR] = "获取当前用户抽奖Id失败"
	message[DB_LOTTERYID_NOTFOUND] = "当前用户没有发起过抽奖"
}

func MapErrMsg(errcode uint32) string {
	if msg, ok := message[errcode]; ok {
		return msg
	} else {
		return "服务器开小差啦,稍后再来试一试"
	}
}

func IsCodeErr(errcode uint32) bool {
	if _, ok := message[errcode]; ok {
		return true
	} else {
		return false
	}
}

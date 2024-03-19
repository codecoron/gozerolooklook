package xerr

// 成功返回
const OK uint32 = 200

/**(前3位代表业务,后三位代表具体功能)**/

// 全局错误码
const (
	SERVER_COMMON_ERROR uint32 = 100001 + iota
	REUQEST_PARAM_ERROR
	TOKEN_EXPIRE_ERROR
	TOKEN_GENERATE_ERROR
	DB_ERROR
	DB_UPDATE_AFFECTED_ZERO_ERROR
	DB_ERROR_NOT_FOUND
)

// 用户模块
const (
	USERCENTER_NOTREGISTER uint32 = 400001 + iota
)

// 投票模块

const (
	VOTE_VOTE_CONFIG_NOT_FOUND uint32 = 120001 + iota
	VOTE_VOTE_RECORD_NOT_FOUND
)

// 抽奖模块
const (
	// 开奖
	AnnounceLottery_ERROR uint32 = 500001 + iota
	GET_PARTICIPATION_USERIDS_BYLOTTERYID_ERROR
	UPDATE_WINNER_ERROR
	GETLOTTERY_BYLESSTHAN_CURRENTTIME_ERROR
	UPDATE_LOTTERY_STATUS_ERROR
	GET_TYPEIS2_AND_ISNOT_ANNOUNCE_LOTTERYS_ERROR
	GET_PARTICIPATORS_COUNT_BYLOTTERYID_ERROR
	// 抽奖详情
	DB_FIND_PRIZES_BYLOTTERYID_ERROR
	DB_FIND_LOTTERY_BYLOTTERYID_ERROR
	DB_FIND_USERID_BYLOTTERYID_ERROR
	DB_USERID_NOTFOUND

	// 发起抽奖
	DB_INSERTLOTTERY_ERROR
	DB_INSERTPRIZE_ERROR

	// 抽奖列表
	DB_GET_LOTTERY_LIST_ERROR
	DB_GETLOTTERYLIST_AFTERLOGIN_ERROR
	DB_GETLASTID_ERROR
	FIND_ALLBYUSERID_ERROR
	DB_FIND_ALLBYUSERID_ERROR

	// 检验当前用户是否参与
	DB_FIND_PARTICIPATOR_ERROR
	// 检查当前用户是否有发起过抽奖
	DB_GET_LOTTERYID_BYUSERID_ERROR
	DB_LOTTERYID_NOTFOUND

	DB_GET_WEEKLOTTERYIDS_BYUSREID_ERROR
	DB_GET_TODAYLOTTERYIDSBYUSERID_ERROR
	DB_GET_LOTTERY_BYLOTTERYID_ERROR
	DB_UPDATE_LOTTERY_ERROR
	DB_NO_SET_LOTTERY_ISSELECT_PERMISSION_ERROR

	// 参与抽奖
	CHECK_ISWON_BYUSERID_ANDLOTTERYID_ERROR
	GET_WONLIST_BYUSERID_ERROR
	GET_WONLISTCOUNT_BYUSERID_ERROR
	CHECK_ISPARTICIPATED_BYUSERID_ANDLOTTERYID_ERROR
	GET_PARTICIPATED_LOTTERYIDS_BYUSERID_ERROR
	FIND_WONLIST_BYUSERID_ERROR

	DB_INSERTCLOCKTASKRECORD_ERROR
)

// 抽奖模块
const LOTTERY_HAS_BEEN_ANOUNCED uint32 = 200001

// 签到模块
const (
	CHECKIN_RECORD_NOT_FOUND uint32 = 700001 + iota
	CHECKIN_REPEAT
	CHECKIN_TASK_NOT_FOUND
	CHECKIN_TASK_REWARD_COLLECTED
	CHECKIN_TASK_NOT_FINISHED
)

// 晒单模块
const (
	ErrUserNotWon uint32 = 800001 + iota
	DB_INSERTCOMMENT_ERROR
	DB_INSERTPRAISE_ERROR
	DB_DELETECOMMENT_ERROR
	DB_FINDCOMMENT_ERROR
)

package xerr

// 成功返回
const OK uint32 = 200

/**(前3位代表业务,后三位代表具体功能)**/

// 全局错误码
const SERVER_COMMON_ERROR uint32 = 100001
const REUQEST_PARAM_ERROR uint32 = 100002
const TOKEN_EXPIRE_ERROR uint32 = 100003
const TOKEN_GENERATE_ERROR uint32 = 100004
const DB_ERROR uint32 = 100005
const DB_UPDATE_AFFECTED_ZERO_ERROR uint32 = 100006

// 用户模块
const USERCENTER_NOTREGISTER uint32 = 500001

// 投票模块
const VOTE_VOTE_CONFIG_NOT_FOUND uint32 = 120001
const VOTE_VOTE_RECORD_NOT_FOUND uint32 = 120002

// 抽奖模块
const LOTTERY_HAS_BEEN_ANOUNCED uint32 = 200001

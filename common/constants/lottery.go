package constants

const (
	_ = iota
	AnnounceTypeTimeLottery
	AnnounceTypePeopleLottery
)

const (
	IsSelectedLottery  = 1
	NotSelectedLottery = 0
)

// 参与抽奖表
const (
	LotteryTodayParticipantsCount    = 5
	LotteryThisWeekParticipantsCount = 10
	IsWon                            = 1
)

// ------ 抽奖 打卡任务 --------

// 任务类型
const (
	_                             = iota
	ExperienceMiniPrograms        // 体验小程序
	BrowseOfficialAccountArticles // 浏览公众号文章
	BrowseImage                   // 浏览图片
	BrowseVideo                   // 浏览视频号视频
)

// 任务秒数
const (
	ExperienceMiniProgramsSeconds        = 15
	BrowseOfficialAccountArticlesSeconds = 15
	BrowseImageSeconds                   = 6
	BrowseVideoSeconds                   = 15
)

// 任务类型文案
const (
	ExperienceMiniProgramsText        = "体验小程序 %d 秒"
	BrowseOfficialAccountArticlesText = "浏览指定公众号文章 %d 秒"
	BrowseImageText                   = "浏览图片 %d 秒"
	BrowseVideoText                   = "浏览视频号视频 %d 秒"
)

// 增加概率类型
const (
	_       = iota
	Random  // 随机增加1～10倍概率
	Appoint // 指定增加1/2/3/4/5/6/7/8/9/10倍概率
)

// 增加概率文案
const (
	RandomText  = "随机增加 1~10 倍"
	AppointText = "增加 %d 倍"
)
